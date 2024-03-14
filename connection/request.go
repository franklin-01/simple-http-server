package connection

import "strings"

type Header string

const (
	Host      Header = "Host"
	UserAgent Header = "User-Agent"
)

type Request struct {
	Method    string
	Path      string
	Protocol  string
	Host      string
	UserAgent string
	Body      []byte
}

func RequestParser(request []byte) Request {

	requestLines := strings.Split(string(request), "\r\n")
	startLine := strings.Split(requestLines[0], " ")

	result := Request{
		Method:   startLine[0],
		Path:     startLine[1],
		Protocol: startLine[2],
	}

	for _, line := range requestLines {

		if line == "" {
			break
		}

		slicedLine := strings.Split(line, ": ")

		header := Header(slicedLine[0])

		switch header {
		case Host:
			result.Host = slicedLine[1]
		case UserAgent:
			result.UserAgent = slicedLine[1]
		}
	}

	req := strings.Split(string(request), "\r\n\r\n")

	if len(req) > 1 {
		result.Body = []byte(req[1])
	}

	return result
}
