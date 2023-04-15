package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const numMessages = 10

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("Connected to server")

	successfulCount := 0
	reader := bufio.NewReader(conn)
	for {
		_, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from server:", err)
			break
		}

		successfulCount++
	}

	fmt.Printf("Received %d out of %d messages (%.2f%%)\n", successfulCount, numMessages, float64(successfulCount)/float64(numMessages)*100)
}
