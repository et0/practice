package main

import "fmt"

func generate(readCh chan<- int) {
	for i := 10; i <= 15; i++ {
		readCh <- i
	}
	close(readCh)
}

func double(readCh <-chan int, writeCh chan<- int) {
	for v := range readCh {
		writeCh <- v * 2
	}
	close(writeCh)
}

func main() {
	generateCh := make(chan int)
	doubleCh := make(chan int)

	go generate(generateCh)
	go double(generateCh, doubleCh)

	for v := range doubleCh {
		fmt.Println(v)
	}

}
