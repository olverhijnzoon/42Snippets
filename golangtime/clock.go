package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func clock(stop chan bool) {
	secondTicker := time.NewTicker(time.Duration(time.Second))
	for {
		select {
		case <-secondTicker.C:
			fmt.Print("\033[H\033[2J") // Move cursor to the start of the line and clear it
			fmt.Printf("%v\r", time.Now().Format("Mon 15:04:05, Jan 2 CEST"))
		case <-stop:
			fmt.Println("\nStopping the clock.")
			return
		}
	}
}

func main() {
	stop := make(chan bool)
	go func() {
		reader := bufio.NewReader(os.Stdin)
		_, _ = reader.ReadString('\n')
		stop <- true
	}()
	clock(stop)
}
