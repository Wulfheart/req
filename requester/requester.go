package requester

import (
	"os"
	"strings"
)

func BuildRequestsFromFile(filepath string, c Config, session *map[string]string, showResult bool) []Request {
	// Prepare File
	content := openAndPrepareFile(filepath)
	// Split file in all requests
	lines := strings.Split(content, Newline)

	requests := make([]Request, 0)

	lastRequestBreakFoundInLine := 0
	var i int
	for _, line := range lines {
		if strings.Index(line, RequestSeparator) == 0 {
			requests = append(requests, Request{rawString: strings.Join(lines[lastRequestBreakFoundInLine:i], Newline), config: &c, sessionVariables: session, ShowResult: showResult})
			lastRequestBreakFoundInLine = i + 1
		}
		i++
	}

	if i != lastRequestBreakFoundInLine {
		requests = append(requests, Request{rawString: strings.Join(lines[lastRequestBreakFoundInLine:], Newline), config: &c, sessionVariables: session, ShowResult: showResult})

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
