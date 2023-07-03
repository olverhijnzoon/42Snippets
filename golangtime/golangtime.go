package main

import (
	"fmt"
	"time"
)

func tickerAndTimer() {
	secondTicker := time.NewTicker(time.Duration(time.Second))
	threeSecondsTimer := time.NewTimer(3 * time.Duration(time.Second))
	for {
		select {
		case <-secondTicker.C:
			fmt.Printf("--- %v\n", time.Now().Format("Mon 15:04:05, Jan 2"))
		case <-threeSecondsTimer.C:
			secondTicker.Stop()
			threeSecondsTimer.Stop()
			fmt.Println("---")
			fmt.Println("--- Out of time ")
			return
		}
	}
}

func main() {

	fmt.Println("# 42Snippets")
	fmt.Println("## Golang Time")

	for i := 0; i < 2; i++ {
		time.Sleep(time.Duration(1000 * 1000 * 1000))
		fmt.Println("---")
		time.Sleep(time.Duration(time.Second))
		fmt.Printf("--- %v\n", time.Now())
	}
	fmt.Println("---")
	fmt.Println("### Clock --")
	fmt.Println("---")
	tickerAndTimer()
	fmt.Println("---")
}
