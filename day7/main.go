package main

import (
	"fmt"
	"time"
)

// worker simulates a worker that processes tasks
func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, job)
		time.Sleep(time.Second) // simulate work
		fmt.Printf("Worker %d finished job %d\n", id, job)
		results <- job * 2 // return result
	}
}

func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// start 3 workers
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// send jobs to the workers
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// collect results
	for a := 1; a <= numJobs; a++ {
		result := <-results
		fmt.Println("Result:", result)
	}
}
