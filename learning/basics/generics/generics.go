package generics

import "fmt"

func sum[T int64 | float64](numbers []T) T {
	var sum T
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func contains[T comparable](elements []T, searchElement T) bool {
	for _, element := range elements {
		if element == searchElement {
			return true
		}
	}
	return false
}

func show[T any](entities ...T) {
	fmt.Println(entities)
}

func showSum() {
	floats := []float64{1.1, 2.2, 3.3}
	ints := []int64{1, 2, 3}

	fmt.Printf("sum of %v is %v\n", floats, sum(floats))
	fmt.Printf("sum of %v is %v\n", ints, sum[int64](ints))

	// wrongGenSlice := []interface{}{
	// 	1.0,
	// 	"str",
	// 	true,
	// }
	// fmt.Println(sum(wrongGenSlice))
}

func showContains() {
	type Person struct {
		name     string
		age      int64
		jobTitle string
	}

	ints := []int64{1, 2, 3, 4, 5}
	fmt.Printf("is %v contains %v: %v\n", ints, 3, contains(ints, 3))

	strings := []string{"a", "b", "c"}
	fmt.Printf("is %v contains %v: %v\n", strings, "a", contains(strings, "a"))
	fmt.Printf("is %v contains %v: %v\n", strings, "d", contains(strings, "d"))

	bobTheManager := Person{name: "Bob", age: 40, jobTitle: "Manager"}
	bobTheFakeManager := Person{name: "Bob", age: 40, jobTitle: "FakeManager"}

	people := []Person{
		{name: "John", age: 30, jobTitle: "Developer"},
		{name: "Jane", age: 25, jobTitle: "Designer"},
		bobTheManager,
	}

	fmt.Printf("is %v contains %v: %v\n", people, bobTheManager, contains(people, bobTheManager))
	fmt.Printf("is %v contains %v: %v\n", people, bobTheFakeManager, contains(people, bobTheFakeManager))
}

func showAny() {
	show(1, 2, 3)
	show("a", "b", "c")
	show(true, false, true)
	show([]int{1, 2, 3}, []int{4, 5, 6})
	show(
		map[string]int64{
			"a": 1,
			"b": 2,
			"c": 3,
		},
	)
}

type number interface {
	~int64 | float64
}

type numbers[T number] []T

// generics can implement only unions or interfaces without methods
func sumUnionInterface[T number](numbers []T) T {
	var sum T
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func unionInterfaceAndType() {
	var ints numbers[int64]
	ints = append(ints, []int64{1, 2, 3}...)

	floats := numbers[float64]{1.1, 2.2, 3.3}

	fmt.Println("sumUnionInterface(ints):", sumUnionInterface(ints))
	fmt.Println("sumUnionInterface(floats):", sumUnionInterface(floats))
}

type customInt int64

func (ci customInt) isPositive() bool {
	return ci > 0
}

func typeApproximation() {
	customInts := []customInt{1, 2, 3}
	castedInts := make([]int64, len(customInts))

	for idx, val := range customInts {
		castedInts[idx] = int64(val)
	}

	// customInt does not satisfy number (possibly missing ~ for int64 in number)
	// if type number is:
	// type number interface {
	//   int64 | float64
	// }
	fmt.Println("sumUnionInterface(customInts):", sumUnionInterface(customInts))
	fmt.Println("sumUnionInterface(castedInts):", sumUnionInterface(castedInts))
}

func Generics() {
	showSum()
	fmt.Println()

	showContains()
	fmt.Println()

	showAny()
	fmt.Println()

	unionInterfaceAndType()
	fmt.Println()

	typeApproximation()
}
