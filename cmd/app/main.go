package main

import (
	"fmt"
	"net"
	"os"
	handlers "github.com/YacineMK/doku/internal/handlers"
)

func main() {
	fmt.Println("Hello, Doku!")

	listener, err := net.Listen("tcp", ":4000")
	if err != nil {
		fmt.Println("Failed to bind to port")
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Println("Listening on port 4000...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handlers.HandleConnection(conn)
	}
}
