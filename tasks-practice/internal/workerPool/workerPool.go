package workerPool

import "sync"

type Job struct {
	Value int
	Func  func(param int) int
}

func New(jobsCh <-chan *Job, count int) <-chan int {
	resultCh := make(chan int)

	wg := sync.WaitGroup{}
	wg.Add(count)

	for range count {
		go func() {
			defer wg.Done()
			for {
				job, ok := <-jobsCh
				if !ok {
					return
				}
				resultCh <- job.Func(job.Value)
			}
		}()
	}

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	return resultCh
}
