package main

import "fmt"

// Define a struct type
type Point struct {
	X, Y int
}

// Function that takes Point by value (copy)
func moveValue(p Point) {
	p.X += 10
	p.Y += 10
}

// Function that takes Point by pointer (reference)
func movePointer(p *Point) {
	p.X += 20
	p.Y += 20
}

// Method with value receiver
func (p Point) MoveBy(dx, dy int) {
	p.X += dx
	p.Y += dy
}

// Method with pointer receiver
func (p *Point) MovePointerBy(dx, dy int) {
	p.X += dx
	p.Y += dy
}

func main() {
	pt := Point{X: 1, Y: 2}

	// Call function by value
	moveValue(pt)
	fmt.Println("After moveValue:", pt) // {1 2} - unchanged

	// Call function by pointer
	movePointer(&pt)
	fmt.Println("After movePointer:", pt) // {21 22} - changed

	// Call method with value receiver
	pt.MoveBy(5, 5)
	fmt.Println("After MoveBy (value receiver):", pt) // {21 22} - unchanged

	// Call method with pointer receiver
	pt.MovePointerBy(5, 5)
	fmt.Println("After MovePointerBy (pointer receiver):", pt) // {26 27} - changed
}
