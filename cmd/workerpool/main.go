package main

import (
	"fmt"
	"sync"
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("worker %d processing job %d\n", id, job)
	}
}

func main() {
	jobs := make(chan int, 10)

	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}

	for j := 1; j <= 10; j++ {
		jobs <- j
	}

	close(jobs)

	wg.Wait()

	fmt.Println("all jobs processed")
}