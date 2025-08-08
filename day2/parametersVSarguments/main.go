package main

import "fmt"

// name and age are PARAMETERS
func greet(name string, age int) {
	fmt.Printf("Hello %s! You are %d years old.\n", name, age)
}

func main() {
	// "Mujammal" and 25 are ARGUMENTS
	greet("Mujammal", 25)
}
