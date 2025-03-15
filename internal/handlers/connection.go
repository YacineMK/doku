package handlers

import (
	"net"
)

func HandleConnection(conn net.Conn) {
	defer conn.Close()
	
		RequestHandler(conn)
}