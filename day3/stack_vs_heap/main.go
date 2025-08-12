package main

import "fmt"

// run `go mod init any_name` for creating module

// Demonstrates simple stack vs heap by returning a pointer vs a value.
func heapEscape() *int {
    x := 42
    // returning address forces x to be heap allocated
    return &x
}

func noEscape() int {
    y := 100
    // returning value can be kept on stack
    return y
}

func main() {
    p := heapEscape()
    fmt.Printf("heapEscape -> pointer: %p value: %d\n", p, *p)

    v := noEscape()
    fmt.Printf("noEscape -> value: %d\n", v)

    fmt.Println()
    fmt.Println("Tip: run `go build -gcflags='-m'` to see compiler escape analysis messages.")
}
