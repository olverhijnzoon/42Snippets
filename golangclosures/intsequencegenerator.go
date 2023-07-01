package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	fmt.Println("# 42Snippets")
	fmt.Println("## Golang Closures")
	fmt.Println("### IntSequenceGenerator Example")

	/*
		Capturing variables from the outer function scope: A closure is not just a function, but also a scope that captures variables from the outer function. This means that a closure can "remember" the values of these variables even after the outer function has finished execution.
	*/

	nextInt := intSeq()
	fmt.Println(nextInt()) // prints 1
	fmt.Println(nextInt()) // prints 2
	fmt.Println(nextInt()) // prints 3
}
