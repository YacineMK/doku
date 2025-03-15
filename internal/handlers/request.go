package handlers

import (
	"fmt"
	"net"
	pkg "github.com/YacineMK/doku/internal/pkg"
)

func RequestHandler(conn net.Conn) error {
	request := make([]byte, 1024)
	_, err := conn.Read(request)
	if err != nil {
		fmt.Println("Read error:", err)
		return err
	}

	parsedRequest, err := pkg.ParseRequest(string(request))
	if err != nil {
		fmt.Println("Parsing error:", err)
		conn.Write([]byte("HTTP/1.1 400 Bad Request\r\n\r\nInvalid Request"))
		return err
	}

	fmt.Println("Parsed Request:", parsedRequest)

	var response string

	switch parsedRequest.Method {
	case "GET":
		response = "HTTP/1.1 200 OK\r\n\r\nGET request received!"
	case "POST":
		response = "HTTP/1.1 200 OK\r\n\r\nPOST request received!"
	case "PUT":
		response = "HTTP/1.1 200 OK\r\n\r\nPUT request received!"
	case "DELETE":
		response = "HTTP/1.1 200 OK\r\n\r\nDELETE request received!"
	case "PATCH":
		response = "HTTP/1.1 200 OK\r\n\r\nPATCH request received!"
	default:
		response = "HTTP/1.1 405 Method Not Allowed\r\n\r\nMethod Not Allowed"
	}

	conn.Write([]byte(response))
	return nil
}
