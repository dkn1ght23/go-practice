package main

import "fmt"

func modify(val *int) {
	*val = 99 // Modify original variable via pointer
}

func main() {
	x := 10
	p := &x // p points to x

	fmt.Println(x) // Output: 10
	modify(p)      // Pass pointer to function
	fmt.Println(x) // Output: 99 (value changed via pointer)
}
