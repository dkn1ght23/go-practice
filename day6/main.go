package main

import (
	"fmt"
)

// INTERFACE: Says "Anything that can Play() is a Toy!"
type Toy interface {
	Play()
}

// STRUCTS: Car and Doll - each one will have its own Play() method.
type Car struct{}

func (c Car) Play() {
	fmt.Println("Vroom! The car zooms!")
}

type Doll struct{}

func (d Doll) Play() {
	fmt.Println("Hello! The doll waves!")
}

// GENERICS: Print any slice of any type.
func Print[T any](items []T) {
	fmt.Println("Printing items:")
	for _, item := range items {
		fmt.Println(item)
	}
}

// TYPE SWITCH: Check what kind of thing is inside the mysteryBox!
func whatsInside(mysteryBox interface{}) {
	switch v := mysteryBox.(type) {
	case Car:
		fmt.Println("It's a Car!")
		v.Play()
	case Doll:
		fmt.Println("It's a Doll!")
		v.Play()
	case string:
		fmt.Println("It's a String:", v)
	default:
		fmt.Println("Something else in the box!")
	}
}

// TYPE ASSERTION: Try to turn the mysteryBox into a Toy.
func tryPlay(mysteryBox interface{}) {
	toy, ok := mysteryBox.(Toy) // Check if it can Play()
	if ok {
		fmt.Println("We found a Toy, let's play!")
		toy.Play()
	} else {
		fmt.Println("This isn't a Toy, can't play.")
	}
}

// CONVERSIONS: Example of converting int to float64
func convertNum() {
	i := 42         // int
	f := float64(i) // Now it's a float64
	fmt.Printf("Converted %d (int) to %.2f (float64)\n", i, f)
}

func main() {
	// INTERFACE: Both Car and Doll are Toys.
	var t Toy

	t = Car{} // Car as a Toy
	t.Play()  // Vroom! The car zooms!

	t = Doll{} // Doll as a Toy
	t.Play()   // Hello! The doll waves!

	// GENERICS: Use Print for any type.
	nums := []int{1, 2, 3}
	words := []string{"red", "blue", "green"}
	Print(nums)
	Print(words)

	// CAR IN A BOX: Put any type in an interface{}
	var mysteryBox interface{}

	mysteryBox = Car{}      // Put a Car inside
	whatsInside(mysteryBox) // Uses type switch

	mysteryBox = Doll{} // Put a Doll inside
	whatsInside(mysteryBox)

	mysteryBox = "Toy Story" // Put a String inside
	whatsInside(mysteryBox)

	// TYPE ASSERTION: Try to play with what's inside.
	mysteryBox = Car{}
	tryPlay(mysteryBox)

	mysteryBox = 100 // Not a Toy!
	tryPlay(mysteryBox)

	// CONVERSION: Change int to float64
	convertNum()
}
