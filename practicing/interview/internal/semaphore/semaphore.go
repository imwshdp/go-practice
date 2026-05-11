package semaphore

import (
	"sync"
	"sync/atomic"
)

type semaphore struct {
	ch chan struct{}
}

func (s *semaphore) Release() {
	<-s.ch
}

func (s *semaphore) Catch() {
	s.ch <- struct{}{}
}

func Run(jobs []int, limit int) {
	sem := semaphore{
		ch: make(chan struct{}, limit),
	}

	for range jobs {
		sem.Catch()

		go func() {
			defer sem.Release()
			// some job
		}()
	}
}

func RunWithCounter(jobs []int, limit int) int {
	count := atomic.Int64{}
	wg := sync.WaitGroup{}

	sem := semaphore{
		ch: make(chan struct{}, limit),
	}

	for range jobs {
		wg.Add(1)
		sem.Catch()

		go func() {
			defer func() {
				sem.Release()
				wg.Done()
			}()
			count.Add(1)
		}()
	}

	wg.Wait()
	return int(count.Load())
}
