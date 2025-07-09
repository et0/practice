package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	files := []int{1, 2, 3, 4, 5, 6, 7, 7, 8, 9, 10, 11, 12}
	sem := make(chan struct{}, 3)
	wg := sync.WaitGroup{}

	for _, v := range files {
		wg.Add(1)
		go func() {
			sem <- struct{}{} // занимаем очередь

			defer wg.Done()

			// освобождаем очередь по завершению
			defer func() {
				<-sem
			}()

			fmt.Printf("Download %d start\n", v)
			time.Sleep(time.Second)
			fmt.Printf("Download %d end\n", v)
		}()
	}

	wg.Wait()
}
