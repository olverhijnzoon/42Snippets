package main

import (
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func TestQuicksort(t *testing.T) {
	testCases := []struct {
		input    []int
		expected []int
	}{
		{[]int{5, 2, 6, 1, 3, -5}, []int{-5, 1, 2, 3, 5, 6}},
		{[]int{5, 2, 6, 1, 3}, []int{1, 2, 3, 5, 6}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{3, 3, 3}, []int{3, 3, 3}},
		{[]int{}, []int{}},
		{[]int{math.MaxInt32, math.MaxInt32, math.MaxInt32}, []int{math.MaxInt32, math.MaxInt32, math.MaxInt32}},
		{[]int{math.MinInt32, math.MinInt32, math.MinInt32}, []int{math.MinInt32, math.MinInt32, math.MinInt32}},
		{[]int{math.MaxInt32, math.MinInt32, 100, -100}, []int{math.MinInt32, -100, 100, math.MaxInt32}},
		{[]int{10, 20, 30, 40, 50, 60}, []int{10, 20, 30, 40, 50, 60}},
		{[]int{-10, -20, -30, -40, -50, -60}, []int{-60, -50, -40, -30, -20, -10}},
		{[]int{0, 10, -10, 20, -20, 30, -30}, []int{-30, -20, -10, 0, 10, 20, 30}},
	}

	for _, tc := range testCases {
		result := Quicksortalgo(tc.input)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("quicksort(%v) = %v, expected %v", tc.input, result, tc.expected)
		}
	}
}

func TestQuicksortPerformance(t *testing.T) {

	/*
		This test function sorts arrays with exponentially growing sizes, and measures the elapsed time for each array. It then prints the elapsed time for each array to the console, and checks whether it is above a certain threshold. If it is above the threshold, the test will fail. By testing the performance of quick sort on arrays with exponentially growing sizes, you can get a sense of how the time complexity of the algorithm scales with the size of the input. If the elapsed time grows roughly in proportion to n*log(n), as expected for quick sort, then this would be strong evidence that the algorithm has a time complexity of O(n*log(n)).
	*/

	// Test sorting arrays with exponential sizes
	sizes := []int{100, 1000, 10000, 100000, 1000000, 10000000}
	thresholds := []time.Duration{time.Millisecond * 10, time.Millisecond * 100, time.Millisecond * 1000, time.Second * 1, time.Second * 10, time.Second * 100}

	for i, size := range sizes {
		// Generate an array of random integers with the current size
		arr := make([]int, size)
		for j := range arr {
			arr[j] = rand.Int()
		}

		// Measure the time it takes to sort the array
		start := time.Now()
		Quicksortalgo(arr)
		elapsed := time.Since(start)

		// Print the elapsed time to the console
		fmt.Printf("Size %d: elapsed time = %s\n", size, elapsed)

		// Check whether the elapsed time is above the threshold
		if elapsed > thresholds[i] {
			t.Errorf("quicksort took %s, which is above the threshold of %s", elapsed, thresholds[i])
		}
	}
}
