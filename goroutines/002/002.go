package main

import "fmt"

func generate(ch chan<- int) { // Только для записи
	for i := 1; i <= 10; i++ {
		ch <- i // Пишим в канал
	}
	close(ch) // Закрытие канала, после отправки
}

func main() {
	ch := make(chan int)

	go generate(ch)

	// Читаем из канала, пока не закроется
	for v := range ch {
		fmt.Println(v)
	}

}
