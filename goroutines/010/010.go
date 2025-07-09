package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	sem := make(chan struct{}, 5)
	wg := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			sem <- struct{}{}

			defer func() {
				<-sem
			}()

			defer wg.Done()

			fmt.Println("Request API", i)
			time.Sleep(time.Millisecond * 1000)
		}()
	}

	wg.Wait()
}
