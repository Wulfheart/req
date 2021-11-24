package requester

import (
	"bufio"
	"bytes"
	"github.com/joho/godotenv"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

type Request struct {
	OriginalRequest       *http.Request
	internalRequestToFire *http.Request
	Response              *http.Response
	ResponseBody          string
	rawString             string
	workingString         string
	config                *Config
	sessionVariables      *SessionVariables
	environmentVariables  map[string]string
	collectionVariables   map[string]string
}

func (r *Request) Prepare() error {
	r.workingString = strings.TrimSpace(r.rawString)
	r.loadVariablesFromFile()
	r.injectVariables()
	r.addContentLengthIfNeeded()
	r.addHttpToFirstLineOfRequestIfNeeded()

	err := r.parse()
	if err != nil {
		return err
	}
	return nil
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
	regex := regexp.MustCompile(TemplateVariableDeclarationLeft + ".*" + TemplateVariableDeclarationRight)

	rawVariables := regex.FindAllString(r.workingString, -1)

	var variables []string
	for _, rawVariable := range rawVariables {
		s := strings.TrimSpace(strings.TrimRight(strings.TrimLeft(rawVariable, TemplateVariableDeclarationLeft), TemplateVariableDeclarationRight))
		if len(s) > 0 {
			variables = append(variables, s)
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
	regex := regexp.MustCompile(TemplateVariableDeclarationLeft + ".*" + varName + ".*" + TemplateVariableDeclarationRight)

	r.workingString = regex.ReplaceAllString(r.workingString, strings.TrimSpace(value))
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

	req, err := http.ReadRequest(bufio.NewReader(strings.NewReader(r.workingString)))

	if err != nil {
		return err
	}

	req.RequestURI = ""

	r.OriginalRequest = req

	req, _ = http.ReadRequest(bufio.NewReader(bytes.NewReader([]byte(r.workingString))))

	req.RequestURI = ""
	r.internalRequestToFire = req

	return nil
}

func (r *Request) DoRequest() error {
	res, err := http.DefaultClient.Do(r.internalRequestToFire)

	if err != nil {
		return err
	}

	r.Response = res

	all, err := io.ReadAll(res.Body)

	if err != nil {
		return err
	}

	r.ResponseBody = string(all)

	return nil
}

func (r *Request) DoStuffAfterTheRequest() {

}
