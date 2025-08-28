package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

const maxThread int = 3

// Написать функцию, которая запрашивает URL из списка и в случае положительного кода 200 выводит
// в stdout в отдельной строке url: , code:
// В случае ошибки выводит в отдельной строке url: , code:
// Функция должна завершаться при отмене контекста.
// Доп. задание: реализовать ограничение количества одновременно запущенных горутин.
func fetchParallel(ctx context.Context, urls []string) {
	semathor := make(chan struct{}, maxThread)
	wg := sync.WaitGroup{}

	for _, u := range urls {
		// проверка контекста перед запуском и
		select {
		case <-ctx.Done():
			return
		case semathor <- struct{}{}:
			// занимаем слот
		}

		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			defer func() {
				<-semathor
			}()

			// полезная нагрузка
			time.Sleep(5 * time.Second)

			select {
			case <-ctx.Done():
				return
			default:
				fmt.Printf("url: %s done\n", url)
			}
		}(u)

	}

	wg.Wait()
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	urls := []string{
		"google.com",
		"yandex.ru",
		"ya.ru",
		"dzen.ru",
	}

	fetchParallel(ctx, urls)
}
