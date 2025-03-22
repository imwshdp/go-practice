package funcs

import "fmt"

func sumFunc(x, y int) int {
	return x + y
}

func subtractFunc(x, y int) int {
	return x - y
}

func calculate(x, y int, action func(x, y int) int) int {
	return action(x, y)
}

func createDivider(divider int) func(x int) int {
	return func(x int) int {
		return x / divider
	}
}

func FuncsAdv() {
	firstNum, secondNum := 6, 3

	var multiplier func(x, y int) int
	multiplier = func(x, y int) int {
		return x * y
	}
	fmt.Println(firstNum, "*", secondNum, "=", multiplier(firstNum, secondNum))

	divider := func(x, y int) int {
		return x / y
	}
	fmt.Println(firstNum, "/", secondNum, "=", divider(firstNum, secondNum))

	fmt.Println(firstNum, "+", secondNum, "=", calculate(firstNum, secondNum, sumFunc))
	fmt.Println(firstNum, "-", secondNum, "=", calculate(firstNum, secondNum, subtractFunc))

	divideByTwo := createDivider(2)
	fmt.Println(firstNum, "/ 2", "=", divideByTwo(firstNum))

	divideByTen := createDivider(10)
	fmt.Println(firstNum, "/ 10", "=", divideByTen(firstNum))

	dollar := 10
	getDollarValue := func() int {
		return dollar
	}

	fmt.Println("Now dollar from closure is about", getDollarValue())

	dollar = 70
	fmt.Println("But now dollar is about", getDollarValue())
}
