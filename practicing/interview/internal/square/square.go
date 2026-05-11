package square

import "sync"

func Slice(data []int) []int {
	numCh := make(chan int, len(data))
	result := make([]int, 0, len(data))

	wg := sync.WaitGroup{}
	wg.Add(len(data))

	for _, num := range data {
		go func() {
			defer wg.Done()
			numCh <- num * num
		}()
	}

	go func() {
		defer close(numCh)
		wg.Wait()
	}()

	for num := range numCh {
		result = append(result, num)
	}

	return result
}
