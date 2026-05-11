package pool

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(
	ctx context.Context,
	wg *sync.WaitGroup,
	workerID int,
	jobsCh chan int,
	resCh chan int,
) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		case val, ok := <-jobsCh:
			if !ok {
				return
			}

			fmt.Printf("worker %d <- %d\n", workerID, val)

			select {
			case <-ctx.Done():
				return
			case resCh <- val:
			}
		}
	}
}

func RunWorkerPool(
	ctx context.Context,
	workerNumber int,
	jobsCh chan int,
) {
	resCh := make(chan int, workerNumber)
	wg := sync.WaitGroup{}

	for inx := range workerNumber {
		wg.Add(1)
		go worker(ctx, &wg, inx, jobsCh, resCh)
	}

	go func() {
		defer close(resCh)
		wg.Wait()
	}()

	for jobRes := range resCh {
		fmt.Printf("resCh <- %d\n", jobRes)
	}
}

func Demo() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	num := 5

	jobsCh := make(chan int, num)

	go func() {
		defer close(jobsCh)

		ticker := time.NewTicker(50000 * time.Microsecond)
		defer ticker.Stop()

		jobID := 0

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				select {
				case <-ctx.Done():
					return
				case jobsCh <- jobID:
					jobID++
				}
			}
		}
	}()

	RunWorkerPool(ctx, num, jobsCh)
}
