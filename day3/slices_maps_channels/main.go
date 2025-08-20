package main

import "fmt"

func main() {
    fmt.Println("Slice backing array growth and addresses:")

    s := make([]int, 1, 2)
    s[0] = 1
    fmt.Printf("initial: len=%d cap=%d &s[0]=%p\n", len(s), cap(s), &s[0])

    s = append(s, 2)
    fmt.Printf("after append 1: len=%d cap=%d &s[0]=%p\n", len(s), cap(s), &s[0])

    s = append(s, 3)
    fmt.Printf("after append 2: len=%d cap=%d &s[0]=%p\n", len(s), cap(s), &s[0])

    fmt.Println()
    fmt.Println("Maps are heap objects; channels also allocate on heap:")

    m := make(map[string]int)
    m["a"] = 1
    fmt.Println("map len:", len(m))

    ch := make(chan int, 2)
    ch <- 1
    fmt.Println("chan len, cap:", len(ch), cap(ch))
    <-ch
    close(ch)

    fmt.Println("\nTip: backing arrays may move to heap depending on escape analysis and usage.")
}
