package main

import (
	"fmt"
	"time"
)

// measureLookupTime measures the time it takes to look up an element in a map
// with the given number of elements, repeating the lookup operation numLookups times.
func measureLookupTime(numElements, numLookups int) time.Duration {
	m := make(map[int]int)
	for i := 0; i < numElements; i++ {
		m[i] = i * i
	}

	totalTime := int64(0)
	for i := 0; i < numLookups; i++ {
		start := time.Now()
		_ = m[numElements-1]
		elapsed := time.Since(start)
		totalTime += elapsed.Nanoseconds()
	}

	return time.Duration(totalTime/int64(numLookups)) * time.Nanosecond
}

func main() {
	fmt.Println("# 42Snippets")
	fmt.Println("## Maps")

	// The program measures the average time it takes to perform a lookup in maps
	// of different sizes. As maps in Go are implemented using hash tables,
	// the lookup time complexity is constant, meaning the average time should
	// be roughly the same regardless of the map size.

	numLookups := 4206912
	avgTime := measureLookupTime(1000, numLookups)
	fmt.Printf("\nAverage time to look up element in map with 1000 elements: %v\n", avgTime)

	avgTime = measureLookupTime(100000, numLookups)
	fmt.Printf("\nAverage time to look up element in map with 100,000 elements: %v\n", avgTime)

	avgTime = measureLookupTime(10000000, numLookups)
	fmt.Printf("\nAverage time to look up element in map with 10,000,000 elements: %v\n", avgTime)
}
