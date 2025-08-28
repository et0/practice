package main

import (
	"fmt"
	"sync"
)

func mergeChs(chs ...<-chan int) chan int {
	// канал, в который будут складываться значения из всех входных каналов
	out := make(chan int)

	// создаётся счётчик ожидания, который будет отслеживать завершение всех горутин
	var wg sync.WaitGroup
	wg.Add(len(chs))

	for _, ch := range chs {
		go func() {
			// уменьшаем счётчик
			defer wg.Done()

			// читаем все значения
			for v := range ch {
				// отправляем в результирующий
				out <- v
			}
		}()
	}

	// обязательно в отдельной горутине, для немедленного возврата результирующего канала
	// это помогает избежать дэдлока
	go func() {
		wg.Wait()  // ждем завершения
		close(out) // и после закрывается канал
	}()

	return out
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	// Записываем данные в каналы
	go func() {
		ch1 <- 1
		ch1 <- 2
		close(ch1)
	}()

	go func() {
		ch2 <- 3
		ch2 <- 4
		close(ch2)
	}()

	merged := mergeChs(ch1, ch2)

	for v := range merged {
		fmt.Println(v)
	}
}
