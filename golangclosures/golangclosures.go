package main

import "fmt"

type callback func(x int) int

func externalFunction(x int, cb callback) int {
	return cb(x)
}

func main() {

	fmt.Println("# 42Snippets")
	fmt.Println("## Golang Closures")

	/*
		In this snippet, externalFunction takes an integer x and a callback function cb. The behavior of externalFunction is modified by passing different closures as the cb argument. In this case, square closure will return the square of the input while the cube closure will return the cube of the input. Without modifying the external function, we can change the behavior of it by passing different closures.

		There are further examples for clarification.
	*/

	square := func(x int) int {
		return x * x
	}

	cube := func(x int) int {
		return x * x * x
	}

	fmt.Println(externalFunction(2, square)) // prints 4
	fmt.Println(externalFunction(2, cube))   // prints 8
}
