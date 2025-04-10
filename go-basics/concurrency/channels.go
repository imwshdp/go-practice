package concurrency

import (
	"fmt"
	"time"
)

func unbufferedChannelWork() {
	// there are reader (main) and writer (custom) goroutines
	unbufferedChannel := make(chan int)

	go func(writingChan chan<- int) {
		time.Sleep(time.Second)
		unbufferedChannel <- 1
	}(unbufferedChannel)

	value := <-unbufferedChannel
	fmt.Println(value)

	go func(readingChan <-chan int) {
		time.Sleep(time.Second)
		value := <-readingChan
		fmt.Println(value)
	}(unbufferedChannel)

	unbufferedChannel <- 2
}

func closedChannel() {
	unbufferedChannel := make(chan int)

	go func() {
		time.Sleep(time.Second / 2)
		close(unbufferedChannel)
	}()

	go func() {
		time.Sleep(time.Second)
		fmt.Println(<-unbufferedChannel)
	}()

	// panic: sending to closed channel
	unbufferedChannel <- 3
}

func doubleClosing() {
	unbufferedChannel := make(chan int)
	close(unbufferedChannel)
	close(unbufferedChannel)
}

func channelDirection() {
	readableOnlyChannel := make(<-chan int)
	fmt.Printf("readableOnlyChannel (%T): %#v\n", readableOnlyChannel, readableOnlyChannel)

	writableOnlyChannel := make(chan<- int)
	fmt.Printf("writableOnlyChannel (%T): %#v\n", writableOnlyChannel, writableOnlyChannel)
}

func iterateByChannelData() {
	bufferedChan := make(chan int, 3)
	numbers := []int{5, 6, 7, 8}

	go func() {
		for _, num := range numbers {
			bufferedChan <- num
		}
		close(bufferedChan)
	}()

	// not valid: values can be read from closed channel - it is default values
	// for {
	// 	value := <-bufferedChan
	// 	fmt.Println(value)
	// }

	for {
		value, ok := <-bufferedChan
		if !ok {
			break
		}
		fmt.Printf("%v ", value)
	}
}

func iterateByChannelDataRange() {
	bufferedChan := make(chan int, 3)
	numbers := []int{5, 6, 7, 8}

	go func() {
		for _, num := range numbers {
			bufferedChan <- num
		}
		close(bufferedChan)
	}()

	for value := range bufferedChan {
		fmt.Printf("%v ", value)
	}
}

func iterateByUnbufferedChannelDataRange() {
	unbufferedChan := make(chan int)
	numbers := []int{5, 6, 7, 8}

	go func() {
		for _, num := range numbers {
			unbufferedChan <- num
		}
		close(unbufferedChan)
	}()

	for value := range unbufferedChan {
		fmt.Printf("%v ", value)
	}
}

func Channels() {
	// nil channel
	var nilChannel chan int
	fmt.Printf("nilChannel (%T): %#v\n", nilChannel, nilChannel)
	fmt.Printf("len = %d, cap = %d\n\n", len(nilChannel), cap(nilChannel))

	// deadlock: sending to nil channel
	// nilChannel <- 1

	// deadlock: reading to nil channel
	// <-nilChannel

	// panic: closing nil channel
	// close(nilChannel)

	// unbuffered channel
	unbufferedChannel := make(chan int)
	fmt.Printf("unbufferedChannel (%T): %#v\n", unbufferedChannel, unbufferedChannel)
	fmt.Printf("len = %d, cap = %d\n\n", len(unbufferedChannel), cap(unbufferedChannel))

	// deadlock: no readers in channel - main goroutine is asleep
	// unbufferedChannel <- 1

	// deadlock: no writers in channel - main goroutine is asleep
	// <-unbufferedChannel

	unbufferedChannelWork()
	fmt.Println()

	// bad example: sending to channel after closing
	// closedChannel()

	// panic: close of closed channel
	// doubleClosing()

	channelDirection()
	fmt.Println()

	// buffered channel
	bufferedChannel := make(chan int, 2)
	fmt.Printf("bufferedChannel (%T): %#v\n", bufferedChannel, bufferedChannel)
	fmt.Printf("len = %d, cap = %d\n", len(bufferedChannel), cap(bufferedChannel))

	bufferedChannel <- 1
	bufferedChannel <- 2

	fmt.Printf("len = %d, cap = %d\n", len(bufferedChannel), cap(bufferedChannel))

	// blocks to write - buffer is full
	// bufferedChannel <- 3

	fmt.Println(<-bufferedChannel)
	fmt.Println(<-bufferedChannel)

	fmt.Printf("len = %d, cap = %d\n\n", len(bufferedChannel), cap(bufferedChannel))

	// blocks to read - buffer is empty, no goroutines writers
	// fmt.Println(<-bufferedChannel)

	// iteration examples by channel buffered values
	iterateByChannelData()
	fmt.Println()

	iterateByChannelDataRange()
	fmt.Println()

	iterateByUnbufferedChannelDataRange()
}
