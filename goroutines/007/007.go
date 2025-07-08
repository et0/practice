package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			select {
			case <-ctx.Done():
				fmt.Println("Exit by context")
				goto Exit
			default:
				fmt.Println("Working...")
			}
		}
	Exit:
	}()
	wg.Wait()
}
