package fanin

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func runLoopForCh(
	ctx context.Context,
	initValue int,
	ch chan int,
	chNum int,
) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	inx := initValue

	for {
		select {
		case <-ctx.Done():
			close(ch)
			return
		case <-ticker.C:
			fmt.Printf("ch%d <- %d\n", chNum, inx)
			ch <- inx
			inx += 1
		}
	}
}

func RunLoop() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	ch1 := make(chan int, 5)
	ch2 := make(chan int, 5)

	go runLoopForCh(ctx, 0, ch1, 1)
	go runLoopForCh(ctx, 100, ch2, 2)

	MergeCh(ctx, ch1, ch2)
}

func MergeCh(
	ctx context.Context,
	ch1 chan int,
	ch2 chan int,
) {
	resultCh := make(chan int, 5)

	go func() {
		for {
			select {
			case <-ctx.Done():
				close(resultCh)
				return
			case val, ok := <-ch1:
				if !ok {
					ch1 = nil
					continue
				}
				resultCh <- val
			case val, ok := <-ch2:
				if !ok {
					ch2 = nil
					continue
				}
				resultCh <- val
			}
		}
	}()

	for val := range resultCh {
		fmt.Printf("resCh <- %d\n", val)
	}
}

func RunSliceLoop() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	ch1 := make(chan int, 5)
	ch2 := make(chan int, 5)
	ch3 := make(chan int, 5)

	go runLoopForCh(ctx, 0, ch1, 1)
	go runLoopForCh(ctx, 100, ch2, 2)
	go runLoopForCh(ctx, 1000, ch3, 3)

	MergeChSlice(ctx, []chan int{ch1, ch2, ch3})
}

func MergeChSlice(ctx context.Context, chs []chan int) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	resCh := make(chan int, 5)
	wg := sync.WaitGroup{}

	for _, ch := range chs {
		wg.Add(1)
		go func(ctx context.Context, ch, resCh chan int) {
			defer wg.Done()

			for {
				select {
				case <-ctx.Done():
					return

				case val, ok := <-ch:
					if !ok {
						return
					}

					select {
					case <-ctx.Done():
						return
					case resCh <- val:
					}
				}
			}
		}(ctx, ch, resCh)
	}

	go func() {
		defer close(resCh)
		wg.Wait()
	}()

	for val := range resCh {
		fmt.Printf("resCh <- %d\n", val)
	}
}
