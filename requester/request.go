package requester

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Wulfheart/req/str"
	"github.com/joho/godotenv"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	v8 "rogchap.com/v8go"
	"strconv"
	"strings"
	"time"
)

type Request struct {
	OriginalRequest       *http.Request
	internalRequestToFire *http.Request
	Response              *http.Response
	ResponseBody          string
	rawString             string
	workingString         string
	config                *Config
	sessionVariables      *map[string]string
	environmentVariables  map[string]string
	collectionVariables   map[string]string
	clientJS              string
	ShowResult            bool
	TimeNeeded            time.Duration
}

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
	return nil
}

func (r *Request) getJS() {
	regex := regexp.MustCompile(`>\s*{%[\w\d\s]*%}`)
	js := regex.FindAllString(r.workingString, 1)
	if len(js) == 0 {
		return
	}

	regex.ReplaceAllString(r.workingString, "")

	regex = regexp.MustCompile(`>\s*{%`)
	regex.ReplaceAllString(js[0], "")
	r.clientJS = strings.TrimSpace(js[0])
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

	r.OriginalRequest = req

	req, _ = http.ReadRequest(bufio.NewReader(bytes.NewReader([]byte(r.workingString))))

	req.RequestURI = ""
	r.internalRequestToFire = req

	return nil
}

func (r *Request) DoRequest() error {
	start := time.Now()
	res, err := http.DefaultClient.Do(r.internalRequestToFire)
	r.TimeNeeded = time.Since(start)

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

func (r *Request) DoStuffAfterTheRequest() error {
	if r.clientJS == "" {
		return nil
	}

	session, err := r.jsonify(*r.sessionVariables)
	if err != nil {
		return err
	}
	environment, err := r.jsonify(r.environmentVariables)
	if err != nil {
		return err
	}
	collection, err := r.jsonify(r.collectionVariables)
	if err != nil {
		return err
	}
	body, err := r.jsonify(r.ResponseBody)
	if err != nil {
		return err
	}
	status, err := r.jsonify(r.Response.StatusCode)
	if err != nil {
		return err
	}
	headers, err := r.jsonify(r.Response.Header)
	if err != nil {
		return err
	}

	f, _ := os.ReadFile("./requester/js/main.js")
	o := str.StrOf(string(f)).Replace("{{session}}", session).
		Replace("{{environment}}", environment).
		Replace("{{collection}}", collection).
		Replace("{{body}}", body).
		Replace("{{status}}", status).
		Replace("{{headers}}", headers).
		Replace("// {{Custom Code}}", r.clientJS)

	x := o.ToString()

	ctx, _ := v8.NewContext() // creates a new V8 context with a new Isolate aka VM
	_, err = ctx.RunScript(x, "main.js")
	if err != nil {
		return err
	}

	val, err := ctx.RunScript("JSON.stringify(client.session)", "main.js")
	if err != nil {
		return err
	}
	fmt.Println(val)
	return nil
}

func (r *Request) jsonify(v interface{}) (string, error) {
	res, err := json.Marshal(v)

	if err != nil {
		return "", err
	}
	return string(res), nil
}
