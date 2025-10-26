package concurrency

import (
	"fmt"
	"sync"
	"time"
)

func withoutWait() {
	for inx := 0; inx < 10; inx++ {
		go fmt.Println(inx + 1)
	}

	fmt.Println("withoutWait done!\n")
}

func withWait() {
	var wg sync.WaitGroup
	wg.Add(10)

	for inx := 0; inx < 10; inx++ {
		go func(index int) {
			fmt.Println(index + 1)
			wg.Done()
		}(inx)
	}

	wg.Wait()
	fmt.Println("withWait done!\n")
}

func wrongAdd() {
	var wg sync.WaitGroup

	for inx := 0; inx < 16; inx++ {
		go func(index int) {
			// it is not good to add goroutine to waitGroup inside new goroutine
			// some of goroutines can be lost
			wg.Add(1)
			defer wg.Done()

			fmt.Println(index + 1)
		}(inx)
	}

	wg.Wait()
	fmt.Println("wrongAdd done!\n")
}

func writeWithoutConcurrent() {
	fmt.Println("writeWithoutConcurrent results:")
	start := time.Now()

	var counter int

	for inx := 0; inx < 1000; inx++ {
		time.Sleep(time.Nanosecond)
		counter++
	}

	fmt.Println(counter)
	fmt.Println(time.Now().Sub(start).Seconds())
}

func writeDataRace() {
	fmt.Println("writeDataRace results:")
	start := time.Now()

	var counter int
	var wg sync.WaitGroup

	wg.Add(1000)

	for inx := 0; inx < 1000; inx++ {
		go func() {
			defer wg.Done()
			time.Sleep(time.Nanosecond)
			counter++
		}()
	}

	wg.Wait()
	fmt.Println(counter)
	fmt.Println(time.Now().Sub(start).Seconds())
}

func writeWithMutex() {
	fmt.Println("writeWithMutex results:")
	start := time.Now()

	var counter int
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(1000)

	for inx := 0; inx < 1000; inx++ {
		go func() {
			defer wg.Done()
			time.Sleep(time.Nanosecond)

			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println(counter)
	fmt.Println(time.Now().Sub(start).Seconds())
}

func readWithMutex() {
	fmt.Println("readWithMutex results:")
	start := time.Now()

	var (
		counter int
		wg      sync.WaitGroup
		mu      sync.Mutex
	)

	wg.Add(100)

	for inx := 0; inx < 50; inx++ {
		go func() {
			defer wg.Done()

			mu.Lock()
			time.Sleep(time.Nanosecond)
			_ = counter
			mu.Unlock()
		}()

		go func() {
			defer wg.Done()

			mu.Lock()
			time.Sleep(time.Nanosecond)
			counter++
			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println(counter)
	fmt.Println(time.Now().Sub(start).Seconds())
}

func readWithRWMutex() {
	fmt.Println("readWithRWMutex results:")
	start := time.Now()

	var (
		counter int
		wg      sync.WaitGroup
		mu      sync.RWMutex
	)

	wg.Add(100)

	for inx := 0; inx < 50; inx++ {
		go func() {
			defer wg.Done()

			mu.RLock()
			time.Sleep(time.Nanosecond)
			_ = counter
			mu.RUnlock()
		}()

		go func() {
			defer wg.Done()

			mu.Lock()
			time.Sleep(time.Nanosecond)
			counter++
			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println(counter)
	fmt.Println(time.Now().Sub(start).Seconds())
}

func Sync() {
	// not valid example - no sync throw goroutines
	withoutWait()

	// example with WaitGroup (sync throw goroutines)
	withWait()

	// not valid example - adding goroutine to waitGroup after creation
	wrongAdd()

	// example without goroutines sync
	writeWithoutConcurrent()
	fmt.Println()

	// example with waitGroup, but with data race
	writeDataRace()
	fmt.Println()

	// example with waitGroup and mutex
	writeWithMutex()
	fmt.Println()

	// read and write operations example with waitGroup and mutex
	readWithMutex()
	fmt.Println()

	// read and write operations example with waitGroup and rwmutex
	readWithRWMutex()
}
