package main

import (
	"fmt"
	"time"
)

// Calculation struct represents a task with a number to calculate its result.
type Calculation struct {
	number, result int
}

func main() {
	const jobsCount = 40
	jobs := make(chan Calculation, jobsCount)
	results := make(chan Calculation, jobsCount)
	start := time.Now()

	// Start worker goroutines
	for w := 1; w <= 4; w++ {
		go worker(w, jobs, results)
	}

	// Send jobs to the jobs channel
	for j := 1; j <= jobsCount; j++ {
		calc := Calculation{j, 0}
		jobs <- calc
	}
	close(jobs)

	// Collect and print results from the results channel
	for a := 1; a <= jobsCount; a++ {
		result := <-results
		fmt.Println("Fib(", result.number, ")=", result.result)
	}
	close(results)

	// Calculate and print the total elapsed time
	duration := time.Since(start)
	fmt.Println("Execution completed! Elapsed time:", duration.Milliseconds(), "milliseconds")
}

// Worker function performs the Fibonacci calculation for each job received from the jobs channel.
func worker(id int, jobs <-chan Calculation, results chan<- Calculation) {
	for j := range jobs {
		fmt.Println("Worker", id, "started task", j.number)
		j.result = calculateFibonacci(j.number)
		results <- j
		fmt.Println("Worker", id, "completed task", j.number)
	}
}

// CalculateFibonacci calculates the Fibonacci number for a given input 'n'.
func calculateFibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return calculateFibonacci(n-1) + calculateFibonacci(n-2)
}
