package main

import (
	"fmt"
	"sync"
)

func main() {
	var (
		wg      sync.WaitGroup
		mu      sync.Mutex
		counter int
	)

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			mu.Lock() // блокируем доступ к counter
			counter++
			fmt.Printf("Worker %d increase counter to %d\n", i, counter)
			mu.Unlock() // снимаем блокировку
		}()
	}

	wg.Wait()

	fmt.Println(counter)
}
