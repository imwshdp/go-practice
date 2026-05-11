package race

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func initial() {
	counter := 0
	wg := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			counter++
		}()
	}

	wg.Wait()

	fmt.Println(counter)
}

type count struct {
	C  int
	mu sync.Mutex
}

func SolutionMutex() {
	counter := count{}
	wg := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func() {
			defer func() {
				counter.mu.Unlock()
				wg.Done()
			}()

			counter.mu.Lock()
			counter.C++
		}()
	}

	wg.Wait()

	fmt.Println(counter.C)
}

func SolutionAtomic() {
	counter := atomic.Int64{}
	wg := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			counter.Add(1)
		}()
	}

	wg.Wait()

	fmt.Println(counter.Load())
}

func SolutionChan() {
	counter := 0

	syncCh := make(chan struct{}, 100)
	wg := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			syncCh <- struct{}{}
		}()
	}

	go func() {
		defer close(syncCh)
		wg.Wait()
	}()

	for range syncCh {
		counter++
	}

	fmt.Println(counter)
}
