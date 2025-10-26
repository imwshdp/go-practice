package collections

import "fmt"

func showAllElements(values ...int) {
	fmt.Printf("showAllElements: ")
	for _, value := range values {
		fmt.Printf("%d ", value)
	}
	fmt.Println()
}

func variadicFunc() {
	// pass slice's values as args
	showAllElements(1, 2, 3)
	showAllElements(1, 2, 3, 4, 5, 6, 7, 8, 9)

	firstSlice := []int{5, 6, 7, 8}
	secondSlice := []int{9, 3, 2, 1}

	showAllElements(firstSlice...)

	// concat two slices
	newSlice := append(firstSlice, secondSlice...)
	fmt.Printf("newSlice (%T): %#v\n", newSlice, newSlice)
}

func convertToArrayPointer() {
	initialSlice := []int{1, 2, 3}
	fmt.Printf("initialSlice (%T): %#v\n", initialSlice, initialSlice)
	fmt.Printf("len = %d, cap = %d\n\n", len(initialSlice), cap(initialSlice))

	// get pointer to slice
	intArray := (*[2]int)(initialSlice)
	fmt.Printf("intArray (%T): %#v\n", intArray, intArray)
	fmt.Printf("len = %d, cap = %d\n", len(intArray), cap(intArray))
}

func changeValue(slice []int) {
	slice[1] = 500
}

func appendValue(slice []int) []int {
	// here new slice created because of cap overflow
	slice = append(slice, 1000, 993, 986)

	fmt.Printf("slice (%T): %#v\n", slice, slice)
	fmt.Printf("len = %d, cap = %d\n", len(slice), cap(slice))

	// so we need to return new slice
	return slice
}

func passToFunction() {
	mySlice := []int{1, 2}
	changeValue(mySlice)

	fmt.Printf("mySlice (%T): %#v\n", mySlice, mySlice)
	fmt.Printf("len = %d, cap = %d\n\n", len(mySlice), cap(mySlice))

	// mySlice is new slice, so assign it to mySlice var for avoid loss of slice

	mySlice = append(mySlice, 3)
	fmt.Printf("mySlice (%T): %#v\n", mySlice, mySlice)
	fmt.Printf("len = %d, cap = %d\n\n", len(mySlice), cap(mySlice))

	// example with overflow capacity and initial slice loss inside function

	mySlice2 := appendValue(mySlice)

	fmt.Printf("mySlice (%T): %#v\n", mySlice, mySlice)
	fmt.Printf("len = %d, cap = %d\n\n", len(mySlice), cap(mySlice))

	fmt.Printf("mySlice2 (%T): %#v\n", mySlice2, mySlice2)
	fmt.Printf("len = %d, cap = %d\n", len(mySlice2), cap(mySlice2))
}

func sliceWithNew() {
	newSlicePointer := new([]int)

	fmt.Printf("newSlicePointer (%T): %#v\n", newSlicePointer, *newSlicePointer)
	fmt.Printf("len = %d, cap = %d\n\n", len(*newSlicePointer), cap(*newSlicePointer))

	// convert pointer to slice by append util
	sliceFromPointer := append(*newSlicePointer, 1, 2)

	fmt.Printf("sliceFromPointer (%T): %#v\n", sliceFromPointer, sliceFromPointer)
	fmt.Printf("len = %d, cap = %d\n", len(sliceFromPointer), cap(sliceFromPointer))
}

func getSlice() {
	intArr := [...]int{1, 2, 3, 4, 5}
	fmt.Printf("intArr (%T): %#v\n\n", intArr, intArr)

	// create slice with array parts as elements
	sliceFromArray := intArr[1:3]
	fmt.Printf("sliceFromArray (%T): %#v\n", sliceFromArray, sliceFromArray)
	fmt.Printf("len = %d, cap = %d\n\n", len(sliceFromArray), cap(sliceFromArray))

	// create slice with full array as elements
	fullSliceFromArray := intArr[:] // intArr[0:len(intArr)]
	fmt.Printf("fullSliceFromArray (%T): %#v\n", fullSliceFromArray, fullSliceFromArray)
	fmt.Printf("len = %d, cap = %d\n\n", len(fullSliceFromArray), cap(fullSliceFromArray))

	// create slice with slice parts as elements
	sliceFromSlice := fullSliceFromArray[:3]
	fmt.Printf("fullSliceFromArray (%T): %#v\n", sliceFromSlice, sliceFromSlice)
	fmt.Printf("len = %d, cap = %d\n\n", len(sliceFromSlice), cap(sliceFromSlice))

	// initial array element changing will affect slice, created with this array element
	intArr[2] = 200
	fmt.Printf("intArr (%T): %#v\n", intArr, intArr)
	fmt.Printf("sliceFromArray (%T): %#v\n", sliceFromArray, sliceFromArray)
}

func copySlice() {
	destination := make([]string, 0, 2)
	source := []string{"Vasya", "Petya", "Ivan"}

	// not valid copy because len of destination is 0
	fmt.Println("Copied", copy(destination, source), "elements")
	fmt.Printf("destination (%T): %#v\n", destination, destination)
	fmt.Printf("len = %d, cap = %d\n\n", len(destination), cap(destination))

	// copy part of source slice
	destination = make([]string, 2, 3)
	fmt.Println("Copied", copy(destination, source), "elements")
	fmt.Printf("destination (%T): %#v\n", destination, destination)
	fmt.Printf("len = %d, cap = %d\n\n", len(destination), cap(destination))

	// copy full source slice with len using
	destination = make([]string, len(source))
	fmt.Println("Copied", copy(destination, source), "(len(source))", "elements")
	fmt.Printf("destination (%T): %#v\n", destination, destination)
	fmt.Printf("len = %d, cap = %d\n\n", len(destination), cap(destination))

	// not valid copy to empty slice (len is 0)
	var emptySlice []string
	fmt.Printf("emptySlice (%T): %#v\n", emptySlice, emptySlice)
	fmt.Printf("len = %d, cap = %d\n\n", len(emptySlice), cap(emptySlice))

	fmt.Println("Copied", copy(emptySlice, source), "elements")
	fmt.Printf("emptySlice (%T): %#v\n", emptySlice, emptySlice)
	fmt.Printf("len = %d, cap = %d\n\n", len(emptySlice), cap(emptySlice))

	// valid copy to empty slice (append + make)
	validCopy := append(make([]string, 0, len(source)), source...)
	fmt.Printf("validCopy (%T): %#v\n", validCopy, validCopy)
	fmt.Printf("len = %d, cap = %d\n", len(validCopy), cap(validCopy))
}

func deleteSliceElements() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Printf("slice (%T): %#v\n", slice, slice)
	fmt.Printf("len = %d, cap = %d\n\n", len(slice), cap(slice))

	// index of element, which need to delete
	inx := 2

	// break initial slice
	withAppend := append(slice[:inx], slice[inx+1:]...)

	fmt.Printf("withAppend (%T): %#v\n", withAppend, withAppend)
	fmt.Printf("len = %d, cap = %d\n", len(withAppend), cap(withAppend))
	fmt.Printf("initial slice = %v\n\n", slice)

	slice = []int{1, 2, 3, 4, 5}
	withCopy := slice[:inx+copy(slice[inx:], slice[inx+1:])]

	fmt.Printf("withCopy (%T): %#v\n", withCopy, withCopy)
	fmt.Printf("len = %d, cap = %d\n", len(withCopy), cap(withCopy))
	fmt.Printf("initial slice = %v\n", slice)
}

func Slices() {
	var defaultSlice []int
	fmt.Printf("defaultSlice (%T): %#v\n", defaultSlice, defaultSlice)

	// len and cap
	fmt.Printf("len = %d, cap = %d\n\n", len(defaultSlice), cap(defaultSlice))

	// initial state of slice
	strLiteralSlice := []string{"a", "b", "c"}
	fmt.Printf("strLiteralSlice (%T): %#v\n", strLiteralSlice, strLiteralSlice)
	fmt.Printf("len = %d, cap = %d\n\n", len(strLiteralSlice), cap(strLiteralSlice))

	// create slice with make
	sliceByMake := make([]int, 0, 5)
	fmt.Printf("sliceByMake (%T): %#v\n", sliceByMake, sliceByMake)
	fmt.Printf("len = %d, cap = %d\n\n", len(sliceByMake), cap(sliceByMake))

	// append to slice
	sliceByMake = append(sliceByMake, 1, 2, 3, 4, 5)
	fmt.Printf("sliceByMake (%T): %#v\n", sliceByMake, sliceByMake)
	fmt.Printf("len = %d, cap = %d\n\n", len(sliceByMake), cap(sliceByMake))

	// capacity change
	sliceByMake = append(sliceByMake, 6)
	fmt.Printf("sliceByMake (%T): %#v\n", sliceByMake, sliceByMake)
	fmt.Printf("len = %d, cap = %d\n\n", len(sliceByMake), cap(sliceByMake))

	// iterate with range
	for inx, value := range sliceByMake {
		fmt.Printf("[%d] -> %d  ", inx, value)
	}
	fmt.Println()

	fmt.Println()
	variadicFunc()

	fmt.Println()
	convertToArrayPointer()

	fmt.Println()
	passToFunction()

	fmt.Println()
	sliceWithNew()

	fmt.Println()
	getSlice()

	fmt.Println()
	copySlice()

	fmt.Println()
	deleteSliceElements()
}
