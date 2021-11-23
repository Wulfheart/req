package requester

import (
	"github.com/joho/godotenv"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type Request struct {
	OriginalRequest      http.Request
	Response             http.Response
	rawString            string
	workingString        string
	config               *Config
	sessionVariables     *SessionVariables
	environmentVariables map[string]string
	collectionVariables  map[string]string
}

func (r *Request) Prepare() {
	r.workingString = strings.TrimSpace(r.rawString)
	r.injectVariables()
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
