package collections

import "fmt"

type person struct {
	Name string
	Age  int
}

func changeArray(arr [3]int) [3]int {
	arr[2] = 100
	return arr
}

func Arrays() {
	var intArray [3]int
	fmt.Printf("intArray (%T): %#v\n", intArray, intArray)

	intArray[0] = 5
	intArray[1] = 6

	fmt.Printf("intArray (%T): %#v\n", intArray, intArray)

	people := [2]person{
		{Name: "Jane", Age: 30},
		{Name: "Jack", Age: 40},
	}

	fmt.Printf("people (%T): %v\n", people, people)

	// creation without len
	stringsArr := [...]string{"a", "b", "c"}
	fmt.Printf("stringsArr (%T): %v\n", stringsArr, stringsArr)

	// len and cap
	fmt.Printf("len = %d, cap = %d\n", len(stringsArr), cap(stringsArr))

	// iterate through a loop
	for i := 0; i < len(stringsArr); i++ {
		fmt.Printf("[%d] -> %s  ", i, stringsArr[i])
	}
	fmt.Println()

	for inx, value := range stringsArr {
		fmt.Printf("[%d] -> %s  ", inx, value)
	}
	fmt.Println()

	for _, value := range stringsArr {
		fmt.Printf("%s ", value)
	}
	fmt.Println()

	newIntArray := changeArray(intArray)
	fmt.Printf("newIntArray (%T): %#v\n", newIntArray, newIntArray)
	fmt.Printf("intArray (%T): %#v\n", intArray, intArray)
}
