package main 

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	fmt.Println("Hello, Doku!")

	l,err := net.Listen("tcp", ":4000")
	if err != nil {
		fmt.Println("Failed to bind to port");
		os.Exit(1)
	}
 
	conn,err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}
	req := make([]byte , 1024)
	conn.Read(req)
	if !strings.HasPrefix(string(req), "GET / HTTP/1.1") {
		conn.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
		return
	}

	conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
}