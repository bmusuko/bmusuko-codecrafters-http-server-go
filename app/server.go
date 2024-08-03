package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage

	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	for {
		// Block until we receive an incoming connection
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		// Handle client connection
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	// Read data
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Printf("failed to read data\n")
		return
	}

	rawStr := string(buf[:n])
	fmt.Printf("raw str %s\n", strconv.Quote(rawStr))

	req := newRequest(rawStr)
	if req.path == "/" {
		conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
	} else {
		conn.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
	}
}
