package requester

import (
	"os"
	"strings"
)

func BuildRequestsFromFile(filepath string, c Config, session *SessionVariables, showResult bool) []Request {
	// Prepare File
	content := openAndPrepareFile(filepath)
	// Split file in all requests
	lines := strings.Split(content, "\n")

	requests := make([]Request, 0)

	lastRequestBreakFoundInLine := 0
	for i, line := range lines {
		if strings.Index(line, RequestSeparator) == 0 {
			requests = append(requests, Request{rawString: strings.Join(lines[lastRequestBreakFoundInLine:i], Newline), config: &c, sessionVariables: session, ShowResult: showResult})
		}
	}

	if len(requests) == 0 {
		requests = append(requests, Request{rawString: strings.Join(lines[:], Newline), config: &c, sessionVariables: session, ShowResult: showResult})
	}

	return requests
}

func openAndPrepareFile(filepath string) (s string) {
	contentBytes, err := os.ReadFile(filepath)
	if err != nil {
		panic(err)
	}

	// For Windows
	s = strings.ReplaceAll(string(contentBytes), UnixNewline, Newline)

	s = strings.Trim(s, RequestSeparator)

	return s
}
