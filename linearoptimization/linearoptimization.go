package main

import (
	"fmt"
	"math"

	"gonum.org/v1/gonum/optimize"
)

func main() {

	fmt.Println("# Golang Snippets")
	fmt.Println("## Linear Optimization")

	/*
		This code demonstrates how to use the optimize package from the Gonum library to perform optimization on two simple functions: a quadratic function and the cosine function.

		The optimize package provides algorithms for solving optimization problems, which are problems that involve finding the minimum or maximum of a function. In this case, the code is using the Minimize function from the optimize package to find the minimum of the two functions.
	*/

	// Definition of a simple quadratic function (spoiler: with minimum at (1, 1)).
	fq := func(x []float64) float64 {
		return (x[0]-1)*(x[0]-1) + (x[1]-1)*(x[1]-1)
	}

	// Set up the optimization problem.
	problemq := optimize.Problem{
		Func: fq,
	}

	// Set the initial point at (42, 69).
	xq := []float64{42, 69}

	// Run the optimization.
	resultq, err := optimize.Minimize(problemq, xq, nil, nil)
	if err != nil {
		panic(err)
	}

	// Print the result.
	fmt.Printf("Minimum found at %v\n", resultq.X)

	// Doing the same with another function: cos(x)
	f := func(x []float64) float64 {
		return math.Cos(x[0])
	}

	// Set up the optimization problem.
	problem := optimize.Problem{
		Func: f,
	}

	// Set the initial point at -3.
	x := []float64{-3}

	// Run the optimization.
	result, err := optimize.Minimize(problem, x, nil, nil)
	if err != nil {
		panic(err)
	}

	// Print the result.
	fmt.Printf("Minimum found at %v\n", result.X)
}
