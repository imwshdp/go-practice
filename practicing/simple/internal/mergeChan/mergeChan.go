package mergeChan

import "sync"

func Merge[T comparable](chans ...<-chan T) <-chan T {
	resultCh := make(chan T, len(chans))

	wg := sync.WaitGroup{}
	wg.Add(len(chans))

	for inx := range len(chans) {
		ch := chans[inx]

		go func() {
			defer wg.Done()
			for val := range ch {
				resultCh <- val
			}
		}()
	}

	go func() {
		wg.Wait()
		close(resultCh)
	}()
	return resultCh
}
