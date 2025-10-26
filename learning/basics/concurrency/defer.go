package concurrency

import (
	"fmt"
)

func sum(a, b int) (sum int) {
	defer func() {
		sum *= 2
	}()

	sum = a + b
	return
}

func deferValues() {
	fmt.Println("deferValues func")

	for inx := 0; inx < 10; inx++ {
		defer fmt.Println("first", inx)
	}

	for inx := 0; inx < 10; inx++ {
		defer func() {
			fmt.Println("second", inx)
		}()
	}

	for inx := 0; inx < 10; inx++ {
		current_inx := inx
		defer func() {
			fmt.Println("third", current_inx)
		}()
	}

	for inx := 0; inx < 10; inx++ {
		defer func(current_inx int) {
			fmt.Println("fourth", current_inx)
		}(inx)
	}
}

func Defer() {
	defer fmt.Println(1)
	defer fmt.Println(2)

	fmt.Println(sum(2, 3))

	deferValues()

	fmt.Println("exit")
}
