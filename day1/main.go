package main

import (
	"fmt"
)

// Package-level variable (available throughout this file)
var globalCounter int = 5

// init function runs automatically before main
func init() {
	fmt.Println(">> init function: Setting up initial state...")
	globalCounter = 10 // You can initialize or modify variables here
}

// Standard function: adds two integers and returns the sum
func add(a int, b int) int {
	return a + b
}

func main() {
	// Local variable, only accessible inside main
	localVar := 15
	fmt.Println("Local variable:", localVar)

	// Using a conditional statement
	if globalCounter > localVar {
		fmt.Println("Global counter is greater than local variable")
	} else {
		fmt.Println("Global counter is not greater than local variable")
	}

	// Standard function usage
	result := add(globalCounter, localVar)
	fmt.Println("Sum from add() function:", result)

	// Immediately Invoked Function Expression (IIFE) in Go
	func(message string) {
		fmt.Println("Inside IIFE:", message)
	}("This function runs right away!")

	// Demonstrate variable scope
	{
		blockVar := "Accessible only inside this block"
		fmt.Println(blockVar)
	}
	// Uncommenting the next line would cause an error, because blockVar is out of scope
	// fmt.Println(blockVar)
}
