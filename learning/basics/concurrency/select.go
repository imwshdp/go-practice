package concurrency

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func baseSelect() {
	// when buffer equals 1 = write into channel will block goroutine
	// so that why select is picking non-blocking operation of read

	// when buffer equals 2 = both of read and write will not block goroutine
	// so that why select is picking one of these operations

	// bufChan := make(chan string, 1)
	bufChan := make(chan string, 2)

	bufChan <- "first"

	select {
	case str := <-bufChan:
		fmt.Println("read:", str)
	case bufChan <- "second":
		fmt.Println("write and read:", <-bufChan, <-bufChan)
	}
}

func selectWithDefault() {
	bufChan := make(chan string, 2)
	unbufChan := make(chan int)

	bufChan <- "first"

	go func() {
		time.Sleep(time.Second)
		unbufChan <- 1
	}()

	select {
	case bufChan <- "second":
		fmt.Println("Unblocking writing")
	case val := <-unbufChan:
		fmt.Println("Blocking reading", val)
	case <-time.After(time.Millisecond * 1500):
		fmt.Println("Time's up")
	default:
		fmt.Println("Default case")
	}
}

func useChannelsWithTimer() {
	resultChan := make(chan int)

	// it is important to create timer outside of select for avoid reassignment
	timer := time.After(time.Second / 4)

	go func() {
		defer close(resultChan)

		for iter := 0; iter < 1000; iter++ {
			select {
			case <-timer:
				fmt.Println("Time's up")
				return
			default:
				time.Sleep(time.Nanosecond)
				resultChan <- iter
			}
		}
	}()

	for value := range resultChan {
		fmt.Print(value, " ")
	}
}

func gracefulShutdown() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	timer := time.After(10 * time.Second)

	for {
		select {
		case <-timer:
			fmt.Println("Time's up")
			return

		case sig := <-signalChan:
			// can be stopped with ctrl + C from console
			fmt.Println("Stopped by signal:", sig)
			return
		}
	}
}

func Select() {
	baseSelect()
	selectWithDefault()
	useChannelsWithTimer()
	gracefulShutdown()
}
