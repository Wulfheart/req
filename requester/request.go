package requester

import (
	"net/http"
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
