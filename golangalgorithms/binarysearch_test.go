package main

import (
	"testing"
)

func TestLow(t *testing.T) {
	// Test the low function with an array of increasing integers and a target value that is smaller than the first element in the array
	array := []int{1, 2, 3, 4, 5}
	target := 0
	expected := 0
	result := low(array, target)
	if result != expected {
		t.Errorf("TestLow: expected %d, got %d", expected, result)
	}

	// Test the low function with an array of increasing integers and a target value that is larger than the last element in the array
	array = []int{1, 2, 3, 4, 5}
	target = 6
	expected = 5
	result = low(array, target)
	if result != expected {
		t.Errorf("TestLow: expected %d, got %d", expected, result)
	}

	// Test the low function with an array of increasing integers and a target value that is equal to an element in the array
	array = []int{1, 2, 3, 4, 5}
	target = 3
	expected = 2
	result = low(array, target)
	if result != expected {
		t.Errorf("TestLow: expected %d, got %d", expected, result)
	}

	// Test the low function with an array of decreasing integers and a target value that is smaller than the first element in the array
	array = []int{5, 4, 3, 2, 1}
	target = 0
	expected = 0
	result = low(array, target)
	if result != expected {
		t.Errorf("TestLow: expected %d, got %d", expected, result)
	}

	// Test the low function with an array of decreasing integers and a target value that is larger than the last element in the array
	array = []int{5, 4, 3, 2, 1}
	target = 6
	expected = 5
	result = low(array, target)
	if result != expected {
		t.Errorf("TestLow: expected %d, got %d", expected, result)
	}
}

func TestHigh(t *testing.T) {
	// Test the high function with an array of increasing integers and a target value that is smaller than the first element in the array
	array := []int{1, 2, 3, 4, 5}
	target := 0
	expected := -1
	result := high(array, target)
	if result != expected {
		t.Errorf("TestHigh: expected %d, got %d", expected, result)
	}

	// Test the high function with an array of increasing integers and a target value that is larger than the last element in the array
	array = []int{1, 2, 3, 4, 5}
	target = 6
	expected = 4
	result = high(array, target)
	if result != expected {
		t.Errorf("TestHigh: expected %d, got %d", expected, result)
	}

	// Test the high function with an array of increasing integers and a target value that is equal to an element in the array
	array = []int{1, 2, 3, 4, 5}
	target = 3
	expected = 2
	result = high(array, target)
	if result != expected {
		t.Errorf("TestHigh: expected %d, got %d", expected, result)
	}

	// Test the high function with an array of decreasing integers and a target value that is smaller than the first element in the array
	array = []int{5, 4, 3, 2, 1}
	target = 0
	expected = -1
	result = high(array, target)
	if result != expected {
		t.Errorf("TestHigh: expected %d, got %d", expected, result)
	}

	// Test the high function with an array of decreasing integers and a target value that is larger than the last element in the array
	array = []int{5, 4, 3, 2, 1}
	target = 6
	expected = 4
	result = high(array, target)
	if result != expected {
		t.Errorf("TestHigh: expected %d, got %d", expected, result)
	}
}

func TestCount(t *testing.T) {
	// Test the count function with an array that contains the target value multiple times
	array := []int{1, 2, 3, 3, 3, 4, 5}
	target := 3
	expected := 3
	result := count(array, target)
	if result != expected {
		t.Errorf("TestCount: expected %d, got %d", expected, result)
	}

	// Test the count function with an array that does not contain the target value
	array = []int{1, 2, 4, 5}
	target = 3
	expected = 0
	result = count(array, target)
	if result != expected {
		t.Errorf("TestCount: expected %d, got %d", expected, result)
	}
}

func TestFirstLastCount(t *testing.T) {
	// Test the firstLastCount function with an array that contains the target value multiple times
	array := []int{1, 2, 3, 3, 3, 4, 5}
	target := 3
	lowExpected := 2
	highExpected := 4
	countExpected := 3
	lowResult, highResult, countResult := firstLastCount(array, target)
	if lowResult != lowExpected {
		t.Errorf("TestFirstLastCount: expected low %d, got %d", lowExpected, lowResult)
	}
	if highResult != highExpected {
		t.Errorf("TestFirstLastCount: expected high %d, got %d", highExpected, highResult)
	}
	if countResult != countExpected {
		t.Errorf("TestFirstLastCount: expected count %d, got %d", countExpected, countResult)
	}
}
