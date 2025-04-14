package concurrency

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func baseKnowledge() {
	ctx := context.Background()
	fmt.Println(ctx)

	todo := context.TODO()
	fmt.Print(todo, "\n\n")

	wValue := context.WithValue(ctx, "name", "Vasya")
	fmt.Print("wValue context value", wValue.Value("name"), "\n\n")

	wCancel, cancel := context.WithCancel(ctx)
	fmt.Println("Err():", wCancel.Err())

	cancel()
	fmt.Print("Err() after cancel:", wCancel.Err(), "\n\n")

	wDeadline, cancel := context.WithDeadline(ctx, time.Now().Add(time.Second))
	defer cancel()

	deadline, ok := wDeadline.Deadline()
	fmt.Println("wDeadline.Deadline():", deadline, ", ok:", ok)
	fmt.Println("wDeadline.Err():", wDeadline.Err())
	fmt.Print("<-wDeadline.Done():", <-wDeadline.Done(), "\n\n")

	wTimeout, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	fmt.Println("<-wTimeout.Done():", <-wTimeout.Done())
}

func worker(ctx context.Context, toProcess <-chan int, processed chan<- int) {
	for {
		select {
		case <-ctx.Done():
			return
		case value, ok := <-toProcess:
			// it is important to check if channel is open
			if !ok {
				return
			}

			time.Sleep(time.Millisecond)
			processed <- value * value
		}
	}
}

func workerPool() {
	baseCtx := context.Background()

	// ctx, ctxCancel := context.WithCancel(baseCtx)
	ctx, ctxCancel := context.WithTimeout(baseCtx, time.Millisecond*20)

	defer ctxCancel()

	wg := &sync.WaitGroup{}

	numberToProcess, processedNumbers := make(chan int, 5), make(chan int, 5)

	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(ctx, numberToProcess, processedNumbers)
		}()
	}

	go func() {
		for i := 0; i < 1000; i++ {
			// if i == 500 {
			// 	ctxCancel()
			// }

			numberToProcess <- i
		}
		close(numberToProcess)
	}()

	go func() {
		wg.Wait()
		close(processedNumbers)
	}()

	var counter int
	for result := range processedNumbers {
		counter++
		fmt.Print(result, " ")
	}

	fmt.Println("\n\ncounter =", counter)
}

func Context() {
	baseKnowledge()
	workerPool()
}
