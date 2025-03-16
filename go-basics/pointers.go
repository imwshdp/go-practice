package main

import "fmt"

func square(num int) {
	num *= num
}

func squarePointer(num *int) {
	*num *= *num
}

func hasWallet(money *int) bool {
	return money != nil
}

func Pointers() {
	var intPointer *int
	fmt.Printf("var intPointer (%T): %#v\n", intPointer, intPointer)

	var a int64 = 7
	fmt.Printf("var a (%T): %#v\n", a, a)

	var pointerA *int64 = &a
	fmt.Printf("var pointerA (%T): %#v => %#v\n", pointerA, pointerA, *pointerA)

	var newPointer = new(float64)
	fmt.Printf("var newPointer (%T): %#v => %#v\n", newPointer, newPointer, *newPointer)

	*newPointer = 3.14
	fmt.Printf("var newPointer (%T): %#v => %#v\n", newPointer, newPointer, *newPointer)

	num := 3
	square(num)
	fmt.Println("num without pointer pass =", num)

	var numPointer *int = &num
	squarePointer(numPointer)
	fmt.Println("num with pointer pass =", *numPointer)

	var wallet1 *int
	wallet2 := 0
	wallet3 := 100

	fmt.Printf("wallet1 has %#v in the balance\n", hasWallet(wallet1))
	fmt.Printf("wallet2 has %#v in the balance\n", hasWallet(&wallet2))
	fmt.Printf("wallet3 has %#v in the balance\n", hasWallet(&wallet3))
}
