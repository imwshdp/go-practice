package main

import "fmt"

func customTypesBehavior() {
	type Celsius float64
	type Fahrenheit float64

	var cel Celsius = 1
	var fah Fahrenheit = 1

	fmt.Printf("1) is celsius equal to fahrenheit? Answer is: %v\n\n", cel == Celsius(fah))
}

func arrayComparison() {
	array1 := [3]int{1, 2, 3}
	array2 := [3]int{1, 2, 3}

	fmt.Printf("2) is array1 (%v) equal to array2 (%v)? Answer is: %v\n\n", array1, array2, array1 == array2)
}

type TestStruct struct {
	str string
	num int
}

func structsArrayComparison() {
	struct1, struct2, struct3 := TestStruct{str: "1", num: 1},
		TestStruct{str: "2", num: 2},
		TestStruct{str: "3", num: 3}

	array1 := [3]TestStruct{
		struct1,
		struct2,
		struct3,
	}

	// Even if struct3 is changed, the result will still be false because arrays are compared by value
	// struct3 = TestStruct{str: "4", num: 4}

	array2 := [3]TestStruct{
		struct1,
		struct2,
		struct3,
	}

	fmt.Printf("3.1) is array1 (%v) equal to array2 (%v)? Answer is: %v\n\n", array1, array2, array1 == array2)

	// 3.2 Slice of pointers comparison
	pointersArrayComparison()
}

func pointersArrayComparison() {
	struct1, struct2, struct3 := &TestStruct{str: "1", num: 1},
		&TestStruct{str: "2", num: 2},
		&TestStruct{str: "3", num: 3}

	array1 := [3]*TestStruct{
		struct1,
		struct2,
		struct3,
	}

	struct3 = &TestStruct{str: "4", num: 4}

	// Even if uncommented, the result will still be false after assigning a new struct to struct3
	// array1[2] = &TestStruct{str: "4", num: 4}

	// But assigning the reference to the same struct in array1 gives true
	// array1[2] = struct3

	array2 := [3]*TestStruct{
		struct1,
		struct2,
		struct3,
	}

	fmt.Printf("3.2) is array1 (%v) equal to array2 (%v)? Answer is: %v\n\n", array1, array2, array1 == array2)
}

func forRangeLoop() {
	// The for-range loop creates a new variable that is a copy of each slice element
	// Modifying this variable does not change the original slice
	fmt.Println("4.1) Changing slice value using loop variable in for-range")
	mySlice := []int{1, 2, 3, 4, 5}
	for _, val := range mySlice {
		val *= 2
		fmt.Println(val)
	}
	fmt.Printf("result: %v\n", mySlice)

	// Same logic applies to slices of structs (even though they are not pointers)
	fmt.Println("\n4.2) Changing slice struct values using loop variable in for-range")
	myStructSlice := []TestStruct{
		{"1", 1},
		{"2", 2},
		{"3", 3},
	}
	for _, val := range myStructSlice {
		val.num += 1
		fmt.Println(val)
	}
	fmt.Printf("result: %v\n", myStructSlice)

	// The for-range loop evaluates the right side once at the beginning
	// That's why this loop is not endless
	fmt.Println("\n4.3) Adding new elements to collection with for-range")
	for _, val := range mySlice {
		mySlice = append(mySlice, val*10)
	}
	fmt.Printf("result: %v\n", mySlice)

	// Since Go 1.22, the value variable is recreating in each iteration
	// That's why we can use it for pointer values
	fmt.Println("\n4.4) When for-range value variable is recreating")
	reduced := make(map[int]*TestStruct, len(myStructSlice))
	for inx, structVal := range myStructSlice {
		fmt.Printf("%v (address=%p)\n", structVal, &structVal)
		reduced[inx] = &structVal
	}
	fmt.Println("result:")
	for inx, val := range reduced {
		fmt.Printf("%d - %v (address=%p)\n", inx, val, &val)
	}

	// The map stores copies of values (not pointers) from the for-range loop variable
	// That's why changes to the original slice do not affect the map
	fmt.Println("\n4.5) Changing values in the original slice")
	myStructSlice[0].num = 777
	myStructSlice[0].str = "777"
	for inx, val := range reduced {
		fmt.Printf("%d - %v (address=%p)\n", inx, val, &val)
	}
}

func forLoop() {
	// In each iteration, the index variable is a new variable with a new address
	fmt.Println("\n5.1) Address of index variable in each iteration")
	for inx := 0; inx < 5; inx++ {
		fmt.Printf("address of inx = %p\n", &inx)
	}
}

func main() {
	// 1. Custom types comparison
	// In Go, two custom types cannot be compared directly even if they have the same underlying type; one must be explicitly converted
	customTypesBehavior()

	// 2. Arrays comparison
	// In Go, arrays are compared by checking the equivalence of each element
	arrayComparison()

	// 3. Struct arrays comparison
	// Same logic: arrays are compared by checking the equivalence of each element
	structsArrayComparison()

	// 4. For-range loop nuances
	forRangeLoop()

	// 5. For loop nuances
	forLoop()
}
