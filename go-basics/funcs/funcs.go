package funcs

import "fmt"

func greet() {
	fmt.Println("Hello from Greet func")
}

func personGreet(name string) {
	fmt.Printf("Hello, %s\n", name)
}

func fullPersonGreet(name, surname string) {
	fmt.Printf("Hello, %s %s\n", name, surname)
}

func sum(first, second int) int {
	sum := first + second
	return sum
}

func sumAndMultiply(first, second int) (int, int) {
	sum := first + second
	multiply := first * second
	return sum, multiply
}

func namedSumAndMultiply(first, second int) (sum int64, multiply int64) {
	sum = int64(first + second)
	multiply = int64(first) * int64(second)

	return // means return sum, multiply
}

func Funcs() {
	greet()
	personGreet("User")
	fullPersonGreet("User", "Smith")

	firstNum, secondNum := 2, 3

	sum := sum(firstNum, secondNum)
	fmt.Println("Sum =", sum)

	sumResult, multiplyResult := sumAndMultiply(firstNum, secondNum)
	fmt.Printf("Sum = %v, Multiply = %v\n", sumResult, multiplyResult)

	_, namedMultiplyResult := namedSumAndMultiply(firstNum, secondNum)
	fmt.Printf("NamedMultiply = %v\n", namedMultiplyResult)
}
