package main

import (
	"fmt"
	"sync"
)

func read(number int, buf <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for v := range buf {
		fmt.Println(number, v)
	}
}

func main() {
	buf := make(chan int)

	wg := sync.WaitGroup{}

	wg.Add(2)

	go read(1, buf, &wg)
	go read(2, buf, &wg)

	for i := range 10 {
		buf <- i
	}
	close(buf)

	wg.Wait()
}
