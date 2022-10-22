package main

import (
	"io"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":1111")
	if err != nil {
		log.Fatalln("Unable to bind on port")
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}
		go handleConn(conn)
	}
}

func handleConn(src net.Conn) {
	dst, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Fatalln("Unable to connect to target server")
	}

	defer dst.Close()
	go func() {
		if _, err := io.Copy(dst, src); err != nil {
			log.Fatalln(err)
		}
	}()
	if _, err := io.Copy(src, dst); err != nil {
		log.Fatalln(err)
	}
}
