package connection

import (
	"fmt"
)

var Status map[int]string = map[int]string{
	200: "HTTP/1.1 200 OK",
	404: "HTTP/1.1 404 Not Found",
}

func GenerateResponse(statusCode int, message string) []byte {
	return []byte(fmt.Sprintf(
		fmt.Sprintf("%s%s%s %d\r\n\r\n%s",
			fmt.Sprintf("%s\r\n", Status[statusCode]),
			"Content-Type: text/plain\r\n",
			"Content-Length:",
			len(message),
			message,
		),
	))
}
