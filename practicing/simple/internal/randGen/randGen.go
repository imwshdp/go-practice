package randGen

import (
	"math"
	"math/rand"
	"time"
)

func New(count int) chan int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	resultCh := make(chan int)

	go func() {
		for range count {
			randInt := r.Intn(math.MaxInt)
			resultCh <- randInt
		}
		close(resultCh)
	}()

	return resultCh
}
