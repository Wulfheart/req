package requester

import (
	_ "embed"
	"encoding/json"
	"github.com/Wulfheart/req/str"
	"github.com/joho/godotenv"
	"os"
	"path/filepath"
	v8 "rogchap.com/v8go"
)

//go:embed js/src/main.js
var file string

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

	status, err := r.jsonify(r.Response.StatusCode)
	if err != nil {
		return err
	}
	headers, err := r.jsonify(r.Response.Header)
	if err != nil {
		return err
	}

	o := str.StrOf(string(file)).Replace("{{session}}", session).
		Replace("{{environment}}", environment).
		Replace("{{collection}}", collection).
		Replace("{{body}}", r.ResponseBody).
		Replace("{{status}}", status).
		Replace("{{headers}}", headers).
		Replace("// {{Custom Code}}", r.clientJS)

	x := o.ToString()

	ctx, _ := v8.NewContext() // creates a new V8 context with a new Isolate aka VM
	_, err = ctx.RunScript(x, "main.js")
	if err != nil {

		return err
	}

	// Session
	val, err := ctx.RunScript("JSON.stringify(client.session.variableBag)", "main.js")
	if err != nil {
		return err
	}
	m := make(map[string]string)
	err = json.Unmarshal([]byte(val.DetailString()), &m)
	if err != nil {
		return err
	}

	for key, value := range m {
		(*r.sessionVariables)[key] = value
	}

	// Environment
	val, err = ctx.RunScript("JSON.stringify(client.environment.variableBag)", "main.js")
	if err != nil {
		return err
	}
	m = make(map[string]string)
	err = json.Unmarshal([]byte(val.DetailString()), &m)
	if err != nil {
		return err
	}
	r.environmentVariables = m
	res, err := godotenv.Marshal(m)
	if err != nil {
		return err
	}
	err = os.WriteFile(filepath.Join(r.config.ConfigPath, EnvironmentVariablesFileName), []byte(res), os.ModePerm)
	if err != nil {
		return err
	}

	// Collection
	val, err = ctx.RunScript("JSON.stringify(client.collection.variableBag)", "main.js")
	if err != nil {
		return err
	}
	m = make(map[string]string)
	err = json.Unmarshal([]byte(val.DetailString()), &m)
	if err != nil {
		return err
	}
	r.environmentVariables = m
	res, err = godotenv.Marshal(m)
	if err != nil {
		return err
	}
	err = os.WriteFile(filepath.Join(r.config.ConfigPath, CollectionVariablesFileName), []byte(res), os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func (r *Request) jsonify(v interface{}) (string, error) {
	res, err := json.Marshal(v)

	if err != nil {
		return "", err
	}
	return string(res), nil
}
