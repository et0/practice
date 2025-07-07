package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var (
		numJobs    = 10
		numWorkers = 3
		wg         sync.WaitGroup
	)
	jobs := make(chan int, numJobs)

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for j := range jobs {
				fmt.Println("Worker", i, "do job ", j)
				time.Sleep(time.Second)
			}
		}()
	}

	for i := 0; i < numJobs; i++ {
		jobs <- i
	}
	close(jobs)

	wg.Wait()
}
