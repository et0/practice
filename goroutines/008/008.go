package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	numWorkers := 3
	numJobs := 100

	jobs := make(chan int, numJobs)
	out := make(chan int, numJobs)

	go func() {
		for i := 1; i <= numJobs; i++ {
			jobs <- i
		}
		close(jobs)
	}()

	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for v := range jobs {
				fmt.Printf("Worker %d get %d and save in out\n", i, v)
				out <- v * 2
				time.Sleep(time.Millisecond * 100)
			}
		}()
	}

	wg.Wait()
	close(out)

	for v := range out {
		fmt.Println(v)
	}

}
