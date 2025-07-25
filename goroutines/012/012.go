package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var (
		mu       sync.Mutex
		cond     = sync.NewCond(&mu)
		tasks    []int
		minTasks = 3
	)

	for i := 0; i < 3; i++ {
		go func() {
			for {
				mu.Lock()
				for len(tasks) < minTasks {
					fmt.Printf("Worker %d wait task\n", i)

					cond.Wait()
				}
				task := tasks[0]
				tasks = tasks[1:]
				mu.Unlock()

				fmt.Printf("Worker %d start task %d\n", i, task)
				time.Sleep(time.Second)
			}
		}()
	}

	for i := 0; i < 10; i++ {
		mu.Lock()
		tasks = append(tasks, i)
		fmt.Printf("Add task %d (all %d)\n", i, len(tasks))
		if len(tasks) >= minTasks {
			cond.Broadcast()
		}
		mu.Unlock()
	}

	time.Sleep(10 * time.Second)
}
