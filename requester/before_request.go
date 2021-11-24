package requester

import (
	"bufio"
	"bytes"
	"github.com/Wulfheart/req/str"
	"github.com/joho/godotenv"
	"net/http"
	"net/textproto"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

func (r *Request) Prepare() error {
	r.workingString = strings.TrimSpace(r.rawString)
	r.getJS()
	r.loadVariablesFromFile()
	r.injectVariables()
	r.addContentLengthIfNeeded()
	r.addHttpToFirstLineOfRequestIfNeeded()

	err := r.parse()
	if err != nil {
		return err
	}

	r.addGlobalHeaders()

	return nil
}

func (r *Request) getJS() {
	// r.workingString = str.StrOf(r.workingString).Replace(Newline, UnixNewline).ToString()
	regex := regexp.MustCompile(`(?s)(>\s*{%)(.+?)(%})`)
	js := regex.FindAllString(r.workingString, 1)
	if len(js) == 0 {
		return
	}

	r.workingString = regex.ReplaceAllString(r.workingString, "")

	regex = regexp.MustCompile(`>\s*{%`)
	js[0] = regex.ReplaceAllString(js[0], "")
	r.clientJS = str.StrOf(strings.TrimSpace(js[0])).Replace("%}", "").ToString()
}

func (r *Request) loadVariablesFromFile() {
	environmentVariablesFilePath := filepath.Join(r.config.ConfigPath, EnvironmentVariablesFileName)
	b, err := os.ReadFile(environmentVariablesFilePath)
	if err != nil {
		panic(err)
	}
	r.environmentVariables, err = godotenv.Unmarshal(string(b))
	if err != nil {
		panic(err)
	}

	collectionVariablesFilePath := filepath.Join(r.config.ConfigPath, CollectionVariablesFileName)
	b, err = os.ReadFile(collectionVariablesFilePath)
	if err != nil {
		panic(err)
	}
	r.collectionVariables, err = godotenv.Unmarshal(string(b))
	if err != nil {
		panic(err)
	}
}

func (r *Request) injectVariables() {

	variables := r.getAllVariables()

	for _, variable := range variables {
		r.replacePlaceholderWithActualValue(variable, r.getVariableValue(variable))
	}

}

func (r Request) getAllVariables() []string {
	return r.getAllVariablesFromString(r.workingString)
}

func (r Request) getAllVariablesFromString(s string) []string {
	regex := regexp.MustCompile(TemplateVariableDeclarationLeft + ".*" + TemplateVariableDeclarationRight)

	rawVariables := regex.FindAllString(s, -1)

	var variables []string
	for _, rawVariable := range rawVariables {
		w := strings.TrimSpace(strings.TrimRight(strings.TrimLeft(rawVariable, TemplateVariableDeclarationLeft), TemplateVariableDeclarationRight))
		if len(s) > 0 {
			variables = append(variables, w)
		}
	}

	return variables
}

func (r Request) getVariableValue(varName string) string {
	// Precedence Session -> Environment -> Collection

	if v, found := (*r.sessionVariables)[varName]; found {
		return v
	}
	if v, found := r.environmentVariables[varName]; found {
		return v
	}
	return r.collectionVariables[varName]
}

func (r *Request) replacePlaceholderWithActualValue(varName string, value string) {
	r.workingString = r.replacePlaceholderWithActualValueOnString(r.workingString, varName, value)
}

func (r *Request) replacePlaceholderWithActualValueOnString(s string, varName string, value string) string {
	regex := regexp.MustCompile(TemplateVariableDeclarationLeft + ".*" + varName + ".*" + TemplateVariableDeclarationRight)

	return regex.ReplaceAllString(s, strings.TrimSpace(value))
}

func (r *Request) addContentLengthIfNeeded() {
	regex := regexp.MustCompile("\\n\\s*\\n")
	splitted := regex.Split(r.workingString, 2)

	regex = regexp.MustCompile("Content-Length:.*")
	if regex.MatchString(splitted[0]) {
		return
	}

	var s string = ""
	x := splitted[0]
	if x[len(x)-1:] != Newline {
		s = Newline
	}
	if len(splitted) < 2 {
		// splitted[0] = x + s + "Content-Length: 0\n"
		//
		r.workingString = splitted[0] + Newline + Newline
		return
	}

	splitted[0] = x + s + "Content-Length: " + strconv.Itoa(len(splitted[1])) + Newline

	r.workingString = splitted[0] + Newline + strings.Join(splitted[1:], "")
}

func (r *Request) addHttpToFirstLineOfRequestIfNeeded() {
	splitted := strings.SplitN(r.workingString, Newline, 2)

	regex := regexp.MustCompile("HTTP(\\d)")
	if regex.MatchString(splitted[0]) {
		return
	}
	s := strings.TrimSpace(splitted[0]) + " " + r.config.HttpVersion + Newline
	r.workingString = s + strings.Join(splitted[1:], "")

}

func (r *Request) parse() error {

	x := str.StrOf(r.workingString).Replace("\r\r", "\r").ToString()
	req, err := http.ReadRequest(bufio.NewReader(strings.NewReader(x)))

	if err != nil {
		return err
	}

	r.OriginalRequest = req

	req, _ = http.ReadRequest(bufio.NewReader(bytes.NewReader([]byte(x))))

	req.RequestURI = ""
	r.internalRequestToFire = req

	return nil
}

func (r *Request) addGlobalHeaders() {
	// Load file
	b, err := os.ReadFile(filepath.Join(r.config.ConfigPath, GlobalHeaderFileName))
	if err != nil {
		panic(err)
	}

	globals := string(b)
	globals = str.StrOf(globals).Replace(UnixNewline, Newline).ToString()
	variables := r.getAllVariablesFromString(globals)

	for _, variable := range variables {
		globals = r.replacePlaceholderWithActualValueOnString(globals, variable, r.getVariableValue(variable))
	}

	reader := bufio.NewReader(strings.NewReader(globals + Newline + Newline))
	tp := textproto.NewReader(reader)

	mimeHeader, err := tp.ReadMIMEHeader()
	if err != nil {
		panic(err)
	}

	httpHeader := http.Header(mimeHeader)

	// Only set if does not exist
	for key, values := range httpHeader {
		if r.OriginalRequest.Header.Values(key) == nil {
			for _, value := range values {
				value = strings.ReplaceAll(value, "\r\r", "\r")
				r.OriginalRequest.Header.Set(key, value)
				// Keep them in sync
				r.internalRequestToFire.Header.Set(key, value)
			}
		}
	}

}
