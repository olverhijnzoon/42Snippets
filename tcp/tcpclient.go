package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const numMessages = 3000

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		os.Exit(1)
	}
	defer conn.Close()

	// TCP retransmission timeouts are OS-managed and cannot be configured in Go. However, read/write deadlines are set here independent of TCP retransmission mechanisms.
	// timeout := 3 * time.Second
	// err = conn.SetDeadline(time.Now().Add(timeout))
	// if err != nil {
	//   fmt.Println("Error setting deadline:", err)
	//	 os.Exit(1)
	// }

	fmt.Println("Connected to server")

	successfulCount := 0
	reader := bufio.NewReader(conn)
	for successfulCount < numMessages {
		_, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from server:", err)
			break
		}
		successfulCount++
		fmt.Printf("Received %d out of %d messages (%.2f%%)\n", successfulCount, numMessages, float64(successfulCount)/float64(numMessages)*100)
	}
}
