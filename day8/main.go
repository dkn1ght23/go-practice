package main

import (
	"fmt"
	"runtime"
	"time"
)

func worker(id int) {
	for i := 1; i <= 3; i++ {
		fmt.Printf("Worker %d doing iteration %d\n", id, i)
		time.Sleep(time.Millisecond * 500) // Simulate work
	}
}

func heavyWorker(id int) {
	start := time.Now()
	sum := 0
	for i := 0; i < 1e8; i++ {
		sum += i
	}
	fmt.Printf("Heavy Worker %d finished in %v\n", id, time.Since(start))
}

func main() {
	// Use all available CPU cores for parallelism
	runtime.GOMAXPROCS(runtime.NumCPU())

	// --- Concurrency Example ---
	fmt.Println("=== Concurrency Example ===")
	go worker(1)
	go worker(2)

	// --- Parallelism Example ---
	fmt.Println("=== Parallelism Example ===")
	go heavyWorker(1)
	go heavyWorker(2)

	// Wait for all goroutines
	time.Sleep(5 * time.Second)
	fmt.Println("Done")
}
