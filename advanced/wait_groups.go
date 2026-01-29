package main

import (
	"fmt"
	"sync"
	"time"
)

type Worker struct {
	ID   int
	Task string
}

func (w *Worker) PerformTask(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("WorkerID %d started %s\n", w.ID, w.Task)
	time.Sleep(time.Second)
	fmt.Printf("WorkerID %d finished %s\n", w.ID, w.Task)
}

func main() {
	var wg sync.WaitGroup

	tasks := []string{"digging", "laying bricks", "painting"}

	for i, task := range tasks {
		worker := Worker{ID: i + 1, Task: task}
		wg.Add(1)
		go worker.PerformTask(&wg)
	}

	wg.Wait()

	fmt.Println("Construction finished")
}

// ----------------------------------------------------

// func worker(id int, tasks <-chan int, result chan<- int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Printf("WorkerID %d starting\n", id)
// 	time.Sleep(time.Second)
// 	for task := range tasks {
// 		result <- task * 2
// 	}

// 	fmt.Printf("WorkerID %d finished\n", id)
// }

// func main() {
// 	var wg sync.WaitGroup
// 	numWorkers := 3
// 	numJobs := 5
// 	results := make(chan int, numJobs)
// 	tasks := make(chan int, numJobs)

// 	wg.Add(numWorkers)

// 	for i := range numWorkers {
// 		go worker(i+1, tasks, results, &wg)
// 	}

// 	for i := range numJobs {
// 		tasks <- i + 1
// 	}
// 	close(tasks)

// 	go func() {
// 		wg.Wait()
// 		close(results)
// 	}()

// 	for result := range results {
// 		time.Sleep(time.Second)
// 		fmt.Println("Result:", result)
// 	}
// }

// ----------------------------------------------------

// func worker(id int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	// wg.Add(1)
// 	fmt.Printf("Worker %d starting\n", id)
// 	time.Sleep(time.Second)
// 	fmt.Printf("Worker %d finished\n", id)
// }

// func main() {
// 	var wg sync.WaitGroup
// 	numWorkers := 3

// 	wg.Add(numWorkers)

// 	for i := range numWorkers {
// 		go worker(i, &wg)
// 	}

// 	wg.Wait()
// 	fmt.Println("All workers finished")

// }
