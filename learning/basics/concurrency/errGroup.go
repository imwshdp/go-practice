package concurrency

import (
	"context"
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

func makeRequest(number int) <-chan string {
	responseChan := make(chan string)

	go func() {
		time.Sleep(time.Second)
		responseChan <- fmt.Sprintf("Response number %d", number)
	}()

	return responseChan
}

func chanAsPromise() {
	firstResponseChan := makeRequest(1)
	secondResponseChan := makeRequest(2)

	// do something
	fmt.Println("do something")

	fmt.Println(<-firstResponseChan, <-secondResponseChan)
}

func chanAsMutex() {
	var counter int
	mutexChan := make(chan struct{}, 1)
	wg := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			mutexChan <- struct{}{}
			counter++
			<-mutexChan
		}()
	}

	wg.Wait()
	fmt.Printf("counter: %d\n", counter)
}

func withoutErrGroup() {
	var err error

	ctx, ctxCancel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}

	wg.Add(3)

	go func() {
		defer wg.Done()
		time.Sleep(time.Second)

		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println("First started")
			time.Sleep(time.Second)
		}
	}()

	go func() {
		defer wg.Done()

		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println("Second started")
			err = fmt.Errorf("any error")
			ctxCancel()
		}
	}()

	go func() {
		defer wg.Done()

		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println("Third started")
			time.Sleep(time.Second)
		}
	}()

	wg.Wait()
	fmt.Println("error:", err)
}

func withErrGroup() {
	g, ctx := errgroup.WithContext(context.Background())

	g.Go(func() error {
		time.Sleep(time.Second)

		select {
		case <-ctx.Done():
			return nil
		default:
			fmt.Println("First started")
			time.Sleep(time.Second)
			return nil
		}
	})

	g.Go(func() error {
		fmt.Println("Second started")
		return fmt.Errorf("any error")
	})

	g.Go(func() error {
		select {
		case <-ctx.Done():
		default:
			fmt.Println("Third started")
			time.Sleep(time.Second)
		}

		return nil
	})

	if err := g.Wait(); err != nil {
		fmt.Println("error:", err)
	}
}

func ErrGroup() {
	chanAsPromise()
	chanAsMutex()

	withoutErrGroup()
	withErrGroup()
}
