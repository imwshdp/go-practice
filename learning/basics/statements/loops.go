package statements

import "fmt"

func Loops() {
	for i := 0; i < 10; i++ {
		fmt.Print(i, " ")
	}

	fmt.Println()

	var i int
	for i = 0; i < 10; i++ {
		fmt.Print(i, " ")
	}

	fmt.Println()

	sum := 1
	for sum < 20 {
		sum += sum
		fmt.Print(sum, " ")
	}

	fmt.Println()

	whileCondition := sum
	for whileCondition < 100 {
		whileCondition++
		fmt.Print(whileCondition, " ")
	}

	fmt.Println()

	// for {
	// 	fmt.Println("Infinite loop")
	// }

	for i := 1; i <= 20; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Print(i, " ")
	}

	fmt.Println()

label:
	for i := 1; i <= 6; i++ {
		for j := 1; j <= 2; j++ {
			if i > 3 {
				continue label
			}
			fmt.Println("I: ", i, ", J: ", j)
		}
	}

	for i := 0; i < 10; i++ {
		if i >= 10 {
			break
		}

		fmt.Print(i, " ")
	}

	fmt.Println()

label2:
	for i := 1; i <= 6; i++ {
		for j := 1; j <= 2; j++ {
			if i > 3 {
				break label2
			}
			fmt.Println("I: ", i, ", J: ", j)
		}
	}
}
