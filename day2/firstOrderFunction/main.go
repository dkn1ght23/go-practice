package main

import "fmt"

/*
A first-order function is just a regular function that:
	i. Takes only data types (int, string, struct, etc.) as parameters
	ii. Returns only data types
	iii. Does not take other functions as parameters
	iv. Does not return other functions
*/

func add(a int, b int) int {
	return a + b
}

func main() {
	sum := add(3, 4)
	fmt.Println(sum)
}
