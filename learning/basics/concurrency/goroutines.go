package concurrency

import (
	"fmt"
	"runtime"
)

func showNumbers(num int) {
	fmt.Print("showNumbers: ")
	for inx := 0; inx < num; inx++ {
		fmt.Print(inx, " ")
	}
	fmt.Println()
}

func Goroutines() {
	// number of logical cores
	fmt.Println("Number of CPUs:", runtime.NumCPU())

	// block showNumbers goroutine because main is the only one available goroutine
	// runtime.GOMAXPROCS(1)

	go showNumbers(10)

	// goroutine swap methods even after runtime.GOMAXPROCS(1)
	// runtime.Gosched() // manually swap goroutine
	// time.Sleep(time.Second) // swap goroutine because of idle

	fmt.Println("goroutines func finished")
}
