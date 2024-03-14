package connection

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func HandleConnection(conn net.Conn, directory string) {

	defer conn.Close()

	buffer := make([]byte, 1024)

	rawRequest, error := conn.Read(buffer)

	if error != nil {
		log.Fatal(error)
	}

	request := RequestParser(buffer[:rawRequest])

	if strings.HasPrefix(request.Path, "/echo") {
		path := strings.TrimPrefix(request.Path, "/echo/")
		conn.Write(GenerateResponse(200, path))
		return
	}

	if strings.HasPrefix(request.Path, "/files") && request.Method == "GET" {
		filepath := strings.TrimPrefix(request.Path, "/files/")
		file, error := os.Open(directory + "/" + filepath)
		if error != nil {
			conn.Write([]byte(fmt.Sprintf("%s\r\n\r\n", Status[404])))
		}

		fileBody, error := io.ReadAll(file)
		if error != nil {
			conn.Write([]byte(fmt.Sprintf("%s\r\n\r\n", Status[404])))
		}

		res := fmt.Sprintf(
			"%s\r\nContent-Type: application/octet-stream\r\nContent-Length: %d\r\n\r\n%s",
			Status[200], len(fileBody), fileBody,
		)

		conn.Write([]byte(res))
		return
	}

	if strings.HasPrefix(request.Path, "/files") && request.Method == "POST" {
		filepath := strings.TrimPrefix(request.Path, "/files/")
		file, error := os.Create(directory + "/" + filepath)
		if error != nil {
			conn.Write([]byte(fmt.Sprintf("HTTP/1.1 500 Internal Server Error\r\n\r\n")))
		}
		_, error = file.Write(request.Body)
		if error != nil {
			conn.Write([]byte(fmt.Sprintf("HTTP/1.1 500 Internal Server Error\r\n\r\n")))
		}

		res := fmt.Sprintf(
			"%s\r\nContent-Type: application/octet-stream\r\nContent-Length: %d\r\n\r\n%s",
			"HTTP/1.1 201 Created", len("File created"), "File Created",
		)
		conn.Write([]byte(res))
		return
	}

	switch request.Path {
	case "/":
		conn.Write([]byte(fmt.Sprintf("%s\r\n\r\n", Status[200])))
	case "/user-agent":
		conn.Write(GenerateResponse(200, request.UserAgent))
	default:
		conn.Write([]byte(fmt.Sprintf("%s\r\n\r\n", Status[404])))
	}

}
