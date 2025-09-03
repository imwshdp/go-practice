package main

import "fmt"

func customTypesBehavior() {
	type Celsius float64
	type Fahrenheit float64

	var cel Celsius = 1
	var fah Fahrenheit = 1

	fmt.Printf("is celsius equals fahrenheit = %v\n", cel == Celsius(fah))
}

func arrayComparison() {
	var array1 = [3]int{1, 2, 3}
	var array2 = [3]int{1, 2, 3}

	fmt.Printf("are int arrays equal = %v\n", array1 == array2)
}

type TestStruct struct {
	str string
	num int
}

func structsArrayComparison() {

	struct1, struct2, struct3 := TestStruct{str: "1", num: 1},
		TestStruct{str: "2", num: 2},
		TestStruct{str: "3", num: 3}

	var array1 = [3]TestStruct{
		struct1,
		struct2,
		struct3,
	}

	// comment for "true" result
	struct3 = TestStruct{str: "4", num: 4}

	var array2 = [3]TestStruct{
		struct1,
		struct2,
		struct3,
	}

	fmt.Printf("\narray1\n")
	for i, v := range array1 {
		fmt.Printf("array1[%d] = %v\n", i, v)
	}

	fmt.Printf("\narray2\n")
	for i, v := range array2 {
		fmt.Printf("array2[%d] = %v\n", i, v)
	}

	fmt.Printf("are struct arrays equal = %v\n", array1 == array2)
}

func pointersArrayComparison() {
	struct1, struct2, struct3 := &TestStruct{str: "1", num: 1},
		&TestStruct{str: "2", num: 2},
		&TestStruct{str: "3", num: 3}

	var array1 = [3]*TestStruct{
		struct1,
		struct2,
		struct3,
	}

	struct3 = &TestStruct{str: "4", num: 4}

	// even if uncomment - there is no "true" result after changing struct3
	// array1[2] = &TestStruct{str: "4", num: 4}

	var array2 = [3]*TestStruct{
		struct1,
		struct2,
		struct3,
	}

	fmt.Printf("\narray1\n")
	for i, v := range array1 {
		fmt.Printf("array1[%d] = %v (%p)\n", i, *v, v)
	}

	fmt.Printf("\narray2\n")
	for i, v := range array2 {
		fmt.Printf("array2[%d] = %v (%p)\n", i, *v, v)
	}

	fmt.Printf("are arrays equal = %v", array1 == array2)
}

func main() {
	// 1. Custom types comparison
	// It is impossible to compare two custom Go types even if they have the same foundation. We need to cast one type to other
	customTypesBehavior()

	// 2. Arrays comparison
	// In Go arrays are compared by the equivalence of each element
	arrayComparison()

	// 3. Struct arrays comparison
	// Same logic, arrays are compared by the equivalence of each element
	structsArrayComparison()

	// 3.1 Addresses arrays comparison
	pointersArrayComparison()
}
