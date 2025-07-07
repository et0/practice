package main

import (
	"fmt"
)

func main() {
	genCh := make(chan int)
	go func() {
		for i := 1; i <= 100; i++ {
			genCh <- i
		}
		close(genCh)
	}()

	filterCh := make(chan int)
	go func() {
		for v := range genCh {
			if v%2 != 0 {
				continue
			}

			filterCh <- v
		}
		close(filterCh)
	}()

	resultCh := make(chan int)
	go func() {
		for v := range filterCh {
			resultCh <- v * 2
		}
		close(resultCh)
	}()

	for v := range resultCh {
		fmt.Println(v)
	}
}
