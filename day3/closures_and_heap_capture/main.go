package main

import "fmt"

// Shows closures capturing state and how that state persists.
func makeCounter(step int) func() int {
    count := 0
    return func() int {
        count += step
        return count
    }
}

func main() {
    c1 := makeCounter(1)
    c2 := makeCounter(10)

    fmt.Println("c1:", c1()) // 1
    fmt.Println("c1:", c1()) // 2
    fmt.Println("c2:", c2()) // 10
    fmt.Println("c2:", c2()) // 20

    fmt.Println("\nNote: the captured variables (like count) typically escape to heap.")
    fmt.Println("Use `go build -gcflags='-m'` to verify which variables moved to heap.")
}
