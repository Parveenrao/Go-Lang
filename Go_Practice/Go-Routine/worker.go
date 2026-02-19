package main
/*
import (
	"fmt"
	"time"
)

// worker RECEIVES jobs and SENDS results
func worker(id int, jobs <-chan int, results chan<- int) {

	for job := range jobs { // receive job

		fmt.Println("Worker", id, "started job", job)

		time.Sleep(time.Second) // pretend work

		result := job * 2

		fmt.Println("Worker", id, "finished job", job)

		results <- result // send result back
	}
}

func main() {

	jobs := make(chan int)    // send jobs here
	results := make(chan int) // receive results here

	// start 3 workers
	go worker(1, jobs, results)
	go worker(2, jobs, results)
	go worker(3, jobs, results)

	// SEND 5 jobs
	for i := 1; i <= 5; i++ {
		jobs <- i
		fmt.Println("Sent job", i)
	}

	close(jobs) // no more jobs

	// RECEIVE 5 results
	for i := 1; i <= 5; i++ {
		res := <-results
		fmt.Println("Received result", res)
	}
}

*/