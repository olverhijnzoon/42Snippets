package main

import (
	"fmt"
	"net"
	"os"
	"sync"
	"time"
)

const numMessages = 3000

func main() {

	fmt.Println("# 42Snippets")
	fmt.Println("## TCP Server")

	/*
		A simple way to test TCP reliability is by creating a server-client setup where the server sends messages to the client, and the client acknowledges receipt. This example does not cover all possible scenarios, but it serves as a starting point for understanding TCP reliability.
	*/

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error listening:", err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Println("Server listening on :8080...")

	// Add a WaitGroup to handle the number of connections
	var wg sync.WaitGroup

	// Allow only one connection to be handled
	wg.Add(1)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go func() {
			handleConnection(conn)
			wg.Done() // Mark the connection as handled
		}()

		break
	}

	// Wait for all connections to be handled
	wg.Wait()

	fmt.Println("Server stopping")
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
