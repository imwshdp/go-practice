package concurrency

import "fmt"

func Panic() {
	defer func() {
		panicValue := recover()
		fmt.Println("panicValue =", panicValue)
	}()

	panic("Panic!")

	fmt.Println("Unreachable code")
}
