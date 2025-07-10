package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func worker(ctx context.Context, workerId int, taskCh <-chan int) error {
	for task := range taskCh {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			if task == 111 {
				return errors.New("Bad task")
			}
			fmt.Println("Worker", workerId, "do task #", task)
			time.Sleep(time.Second)
		}
	}

	return nil
}

func main() {
	var (
		numJobs   = 120
		numWorker = 10
		numBuffer = 10
	)

	g, ctx := errgroup.WithContext(context.Background())
	taskCh := make(chan int, numBuffer)

	for i := 0; i < numWorker; i++ {
		// workerId := 1
		g.Go(func() error {
			return worker(ctx, i, taskCh)
		})
	}

	g.Go(func() error {
		defer close(taskCh)
		for i := 0; i < numJobs; i++ {
			select {
			default:
				taskCh <- i
			case <-ctx.Done():
				return ctx.Err()
			}
		}

		return nil
	})

	if err := g.Wait(); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Done.")
	}
}
