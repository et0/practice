package main

import "fmt"

func generate(readCh chan<- int) { // канал только для записи
	for i := 10; i <= 15; i++ {
		readCh <- i // записываем в канал
	}
	close(readCh) // закрываем канал, т.к. записей в него больше не буден
}

func double(readCh <-chan int, writeCh chan<- int) { // первый канал только чтение, второй канал для записи
	for v := range readCh {
		writeCh <- v * 2 // увеличиваем на два и записываем в канал
	}
	close(writeCh) // закрываем канал, т.к. записей больше не будет
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
