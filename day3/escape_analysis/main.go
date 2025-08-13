package main

import "fmt"

// This file contains small examples that commonly trigger escape analysis.
// Build with: go build -gcflags='-m' to inspect what moves to heap.

// returning pointer -> escapes
func returnPtr() *int {
    a := 7
    return &a
}

// taking address of composite literal -> usually heap
func compositeLiteral() *struct{ n int } {
    s := &struct{ n int }{n: 9}
    return s
}

// closure capturing local variable -> escapes
func closureEscape() func() int {
    x := 1
    return func() int { // the closure captures x
        x++
        return x
    }
}

func main() {
    p := returnPtr()
    fmt.Println("returnPtr value ->", *p)

    c := compositeLiteral()
    fmt.Println("compositeLiteral ->", c.n)

    ctr := closureEscape()
    fmt.Println("closureEscape -> first:", ctr(), "second:", ctr())

    fmt.Println("\nCheck escape analysis with: go build -gcflags='-m' .")
}
