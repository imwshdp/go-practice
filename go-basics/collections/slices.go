package collections

import "fmt"

func Slices() {
	var defaultSlice []int
	fmt.Printf("defaultSlice (%T): %#v\n", defaultSlice, defaultSlice)

	// len and cap
	fmt.Printf("len = %d, cap = %d\n", len(defaultSlice), cap(defaultSlice))

	strLiteralSlice := []string{"a", "b", "c"}
	fmt.Printf("strLiteralSlice (%T): %#v\n", strLiteralSlice, strLiteralSlice)
	fmt.Printf("len = %d, cap = %d\n", len(strLiteralSlice), cap(strLiteralSlice))

	sliceByMake := make([]int, 0, 5)
	fmt.Printf("sliceByMake (%T): %#v\n", sliceByMake, sliceByMake)
	fmt.Printf("len = %d, cap = %d\n", len(sliceByMake), cap(sliceByMake))

	sliceByMake = append(sliceByMake, 1, 2, 3, 4, 5)
	fmt.Printf("sliceByMake (%T): %#v\n", sliceByMake, sliceByMake)
	fmt.Printf("len = %d, cap = %d\n", len(sliceByMake), cap(sliceByMake))

	sliceByMake = append(sliceByMake, 6)
	fmt.Printf("sliceByMake (%T): %#v\n", sliceByMake, sliceByMake)
	fmt.Printf("len = %d, cap = %d\n", len(sliceByMake), cap(sliceByMake))

	for inx, value := range sliceByMake {
		fmt.Printf("[%d] -> %d  ", inx, value)
	}
}
