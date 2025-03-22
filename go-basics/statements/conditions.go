package statements

import (
	"fmt"
	"math/rand"
	"time"
)

func isChildren(age int) bool {
	return age < 18
}

func getFour() int {
	return 4
}

func Conditions() {
	age := 15

	if age < 18 {
		fmt.Println("Too young for this")
	}

	if isChild := isChildren(age); isChild == true {
		fmt.Println("You are too young for this condition")
	}

	if age < 5 {
		fmt.Println("You are a baby (if)")
	} else {
		fmt.Println("You are a child (else)")
	}

	if age >= 10 && age <= 18 {
		fmt.Println("You are a teenager")
	}

	if age == 14 || age == 20 || age == 40 {
		fmt.Println("Go get a new passport")
	}

	if !isChildren(age) {
		fmt.Println("You are an adult")
	}

	const (
		min = 1
		max = 5
	)

	rand.Seed(time.Now().UnixNano())
	randValue := rand.Intn(max-min) + min

	fmt.Println("Random value:", randValue)

	if randValue == 1 {
		fmt.Println("You got a 1")
	} else if randValue == 2 || randValue == 3 {
		fmt.Println("You got a 2 or 3")
	} else if randValue == getFour() {
		fmt.Println("You got a 4")
	} else {
		fmt.Println("You got something else")
	}

	switch randValue {
	case 1:
		fmt.Println("You got a 1")
	case 2, 3:
		fmt.Println("You got a 2 or 3")
	case getFour():
		fmt.Println("You got a 4")
	default:
		fmt.Println("You got something else")
	}

	switch {
	case randValue < 2:
		fmt.Println(randValue, "is less than 2")

	case randValue > 2:
		fmt.Println(randValue, "is greater than 2")
	}

	switch num := rand.Intn(max-min) + min; num {
	case 1:
		fmt.Println("You got a 1")
	case 2, 3:
		fmt.Println("You got a 2 or 3")
	case getFour():
		fmt.Println("You got a 4")
		fallthrough
	case 10:
		fmt.Println("10 is impossible to get, it is fallthrough")
	default:
		fmt.Println("You got something else")
	}
}
