package concurrency

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func syncWithMutex() {
	fmt.Print("syncWithMutex results: ")
	start := time.Now()

	var (
		counter int
		wg      sync.WaitGroup
		mu      sync.Mutex
	)

	wg.Add(1000)

	for inx := 0; inx < 1000; inx++ {
		go func() {
			defer wg.Done()

			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Printf("%v, ", counter)
	fmt.Printf("%v\n", time.Now().Sub(start).Seconds())
}

func syncWithAtomic() {
	fmt.Print("syncWithAtomic results: ")
	start := time.Now()

	var (
		counter int64
		wg      sync.WaitGroup
	)

	wg.Add(1000)

	for inx := 0; inx < 1000; inx++ {
		go func() {
			defer wg.Done()
			atomic.AddInt64(&counter, 1)
		}()
	}

	wg.Wait()
	fmt.Printf("%v, ", counter)
	fmt.Printf("%v\n", time.Now().Sub(start).Seconds())
}

func storeLoadSwap() {
	var counter int64

	fmt.Println(atomic.LoadInt64(&counter))

	atomic.StoreInt64(&counter, 100)
	fmt.Println(atomic.LoadInt64(&counter))

	fmt.Println(atomic.SwapInt64(&counter, 10))
	fmt.Println(atomic.LoadInt64(&counter))
}

func compareAndSwap() {
	var (
		counter int64
		wg      sync.WaitGroup
	)

	wg.Add(100)

	for inx := 0; inx < 100; inx++ {
		go func(inx int) {
			defer wg.Done()

			if !atomic.CompareAndSwapInt64(&counter, 0, 1) {
				return
			}

			fmt.Println("Swapped goroutine number is", inx)
		}(inx)
	}

	wg.Wait()
	fmt.Println("counter:", counter)
}

func atomicValue() {
	var value atomic.Value

	value.Store(1)
	fmt.Println(value.Load())

	fmt.Println(value.Swap(2))
	fmt.Println(value.Load())

	fmt.Println(value.CompareAndSwap(2, 100))
	fmt.Println(value.Load())
}

func Atomic() {
	syncWithMutex()
	syncWithAtomic()
	fmt.Println()

	storeLoadSwap()
	fmt.Println()

	compareAndSwap()
	fmt.Println()

	atomicValue()
}
