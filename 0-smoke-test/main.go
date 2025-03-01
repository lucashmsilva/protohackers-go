package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", "0.0.0.0:6767")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 4096)

	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err != io.EOF {
				fmt.Println("reader.Read failed:", err)
			}
			break
		}

		if n > 0 {
			fmt.Printf("Received: %s\n", buf)

			_, err := conn.Write(buf[:n])
			if err != nil {
				fmt.Println("conn.Write failed:", err)
				break
			}
		}
	}
}
