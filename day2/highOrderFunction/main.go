package main

import "fmt"

/*
A higher-order function is a function that either:
	i. Takes another function as a parameter
	ii. Returns a function as its result
	iii. Or does both
*/

// higherOrder takes a function as a parameter and returns a new function
func higherOrder(operation func(int, int) int) func(int) int {
	return func(x int) int {
		// Call the passed-in function with x and a constant value
		return operation(x, 10)
	}
}

// Example operations
func add(a, b int) int {
	return a + b
}

func multiply(a, b int) int {
	return a * b
}

func main() {
	// Pass the 'add' function and get back a function that adds 10
	add10 := higherOrder(add)
	// Pass the 'multiply' function and get back a function that multiplies by 10
	mul10 := higherOrder(multiply)

	fmt.Println(add10(5))
	fmt.Println(mul10(5))
}
