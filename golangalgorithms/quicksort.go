package main

import (
	"math/rand"
)

func Quicksortalgo(sliceOfIntegers []int) []int {

	/*
		fmt.Println("# 42Snippets")
		fmt.Println("## Quicksort")

		Quicksort is a divide and conquer algorithm that works by selecting a "pivot" element from the array and partitioning the other elements into two sub-arrays according to whether they are less than or greater than the pivot. The sub-arrays are then recursively sorted using the same process.

		Here is how the algorithm works in more detail:

			1. If the length of the input array is less than 2, it is already sorted and the function returns it.
			2. Otherwise, the pivot element is selected as the last element in the array and the other elements are partitioned into two sub-arrays: one containing the elements that are less than the pivot, and the other containing the elements that are greater than the pivot.
			3. The pivot element is then placed between the two sub-arrays and the function is called recursively on each sub-array to sort them.
			4. The function returns the sorted array.

		The time complexity of quicksort is O(n*log(n)), which makes it more efficient than many other sorting algorithms on average. However, it has a worst-case time complexity of O(n^2) if the pivot element is always the smallest or largest element in the array, so it is important to choose the pivot element wisely in order to avoid this case.
	*/

	if len(sliceOfIntegers) < 2 {
		return sliceOfIntegers
	}

	left, right := 0, len(sliceOfIntegers)-1

	pivotIndex := rand.Int() % len(sliceOfIntegers)
	sliceOfIntegers[pivotIndex], sliceOfIntegers[right] = sliceOfIntegers[right], sliceOfIntegers[pivotIndex]

	for i := range sliceOfIntegers {
		if sliceOfIntegers[i] < sliceOfIntegers[right] {
			sliceOfIntegers[left], sliceOfIntegers[i] = sliceOfIntegers[i], sliceOfIntegers[left]
			left++
		}
	}

	sliceOfIntegers[left], sliceOfIntegers[right] = sliceOfIntegers[right], sliceOfIntegers[left]

	Quicksortalgo(sliceOfIntegers[:left])
	Quicksortalgo(sliceOfIntegers[left+1:])

	return sliceOfIntegers
}
