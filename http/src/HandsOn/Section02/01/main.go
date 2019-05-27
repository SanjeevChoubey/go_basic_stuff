package main

import (
	"io"
	"log"
	"net"
)

func main() {
	listner, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer listner.Close()
for{
	conn, err := listner.Accept()
	if err != nil {
		log.Fatalln(err)
	}

	io.WriteString(conn, "I see you connected.")
	conn.Close()
}
}
