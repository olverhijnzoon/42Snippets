package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {

	fmt.Println("# Golang Snippets")
	fmt.Println("## Wait Groups")

	/*
		In this first example, we launch five goroutines using a loop and pass a unique number to each goroutine. Each goroutine prints a message and sleeps for one second before signaling that it has finished by calling Done on the WaitGroup. The main goroutine then calls Wait to block until all five goroutines have signaled that they are done. Finally, it prints a message indicating that all goroutines have finished.
	*/

	var wg sync.WaitGroup

	// Launch five goroutines
	wg.Add(5)
	for i := 1; i <= 5; i++ {
		go func(num int) {
			defer wg.Done()
			fmt.Printf("Hello from goroutine %d\n", num)
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			fmt.Printf("Byebye from goroutine %d\n", num)
		}(i)
	}

	// Wait for the goroutines to finish
	wg.Wait()
	fmt.Println("All goroutines have finished")

	fmt.Println("Second example")

	/*
		In this second example, the main goroutine waits for the user to press Enter, then closes the done channel and waits for the goroutines to stop by calling Wait on the WaitGroup. Finally, it prints a message indicating that all goroutines have stopped.
	*/

	// Create a new WaitGroup
	wg = sync.WaitGroup{}

	// Create a channel to signal when to stop the goroutines
	done := make(chan struct{})

	// Create some random offset to simulate work time
	work := time.Duration(rand.Intn(5000))

	// Launch three goroutines
	wg.Add(3)
	for i := 1; i <= 3; i++ {
		go func(num int) {
			defer wg.Done()
			for {
				select {
				case <-done:
					fmt.Printf("Goroutine %d stopped\n", num)
					return
				default:
					work = time.Duration(rand.Intn(5000))
					time.Sleep((work) * time.Millisecond)
					fmt.Printf("Work %d finished goroutine %d - press any key to stop the goroutines\n", work, num)
				}
			}
		}(i)
	}

	// Wait for a key press
	fmt.Println("Press Enter to stop the goroutines")
	fmt.Scanln()

	// Send a signal to stop the goroutines
	fmt.Println("Stopping the goroutines")
	close(done)

	// Wait for the goroutines to stop
	wg.Wait()
	fmt.Println("All goroutines have stopped")
}
