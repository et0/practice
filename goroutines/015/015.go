package main

import (
	"context"
	"fmt"
	"time"
)

func slowFunc() bool {
	time.Sleep(5 * time.Second)
	return true
}

func ctxFunc(ctx *context.Context) (bool, error) {
	ch := make(chan bool, 1)

	go func() {
		ch <- slowFunc()
	}()

	// ожидает одного из двух вариантов
	select {
	case v := <-ch:
		return v, nil
	case <-(*ctx).Done():
		return false, (*ctx).Err()
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(2*time.Second))
	defer cancel()

	result, err := ctxFunc(&ctx)
	fmt.Print(result, err)
}
