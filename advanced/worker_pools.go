package main

import (
	"fmt"
	"time"
)

type ticketRequest struct {
	personID   int
	numTickets int
	cost       int
}

func ticketProcessor(requests <-chan ticketRequest, results chan<- int) {
	for req := range requests {
		fmt.Printf("Processing %d ticket of personID %d with total cost %d\n", req.numTickets, req.personID, req.cost)
		time.Sleep(time.Second)
		results <- req.personID
	}

}

func main() {
	numRequests := 5
	price := 5
	ticketRequests := make(chan ticketRequest, numRequests)
	ticketResults := make(chan int)

	for range 3 {
		go ticketProcessor(ticketRequests, ticketResults)
	}

	for i := range numRequests {
		ticketRequests <- ticketRequest{personID: i + 1, numTickets: (i + 1) * 2, cost: (i + 1) * price}
	}
	close(ticketRequests)

	for range numRequests {
		fmt.Printf("Ticket for personID %d processed successfully\n", <-ticketResults)
	}

}

// func worker(id int, tasks <-chan int, result chan<- int) {
// 	for task := range tasks {
// 		fmt.Printf("Worker %d processing task %d\n", id, task)
// 		time.Sleep(time.Second)
// 		result <- task * 2
// 	}
// }

// func main() {

// 	numWorkers := 3
// 	numJobs := 10
// 	tasks := make(chan int, numJobs)
// 	results := make(chan int, numJobs)

// 	for i := range numWorkers {
// 		go worker(i, tasks, results)
// 	}

// 	for i := range numJobs {
// 		tasks <- i
// 	}
// 	close(tasks)

// 	for range numJobs {
// 		result := <-results
// 		fmt.Println("Result:", result)
// 	}

// }
