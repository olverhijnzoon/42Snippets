package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

const numMessages = 1000

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error listening:", err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Println("Server listening on :8080...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("Client connected:", conn.RemoteAddr())

	for i := 1; i <= numMessages; i++ {
		message := fmt.Sprintf("Message %d\n", i)
		conn.Write([]byte(message))
		time.Sleep(50 * time.Millisecond)
	}
}
