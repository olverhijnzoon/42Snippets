package main

import "fmt"

func outer(x int) func(int) int {
	return func(y int) int {
		return x + y // x from outer function, y from current function
	}
}

func main() {

	fmt.Println("# 42Snippets")
	fmt.Println("## Golang Closures")
	fmt.Println("### SumGenerator Example")

	/*
		Closures have access to their own and the outer function's variables: A closure can use its own variables as well as those from the outer function. However, it cannot access variables from other functions. Even if two closures are created by the same outer function, they are unique and have their own copies of the outer function's variables.
	*/

	addFive := outer(5)
	fmt.Println(addFive(37)) // Prints 42
	addTwo := outer(2)
	fmt.Println(addTwo(7)) // Prints 9
}
