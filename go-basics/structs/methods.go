package structs

import "fmt"

type square struct {
	Side int
}

func (s square) Perimeter() {
	fmt.Printf("mySquar s(%T): %#v\n", s, s)
	fmt.Printf("Perimeter: %v\n\n", s.Side*4)
}

func (s *square) Scale(multiplier int) {
	fmt.Printf("mySquar s(%T): %#v\n", s, s)
	s.Side *= multiplier
	fmt.Printf("Scaled: %v\n\n", s.Side)
}

func (s square) WrongScale(multiplier int) {
	fmt.Printf("mySquar s(%T): %#v\n", s, s)
	s.Side *= multiplier
	fmt.Printf("Wrong scaled s(%T): %#v\n\n", s, s)
}

func Methods() {
	mySquare := square{Side: 10}
	pMySquare := &mySquare

	mySquare2 := square{Side: 2}

	fmt.Printf("===\nVALUE AND POINTER RECEIVERS:\n===\n")
	mySquare.Perimeter()
	mySquare2.Perimeter()
	pMySquare.Scale(2)

	fmt.Printf("===\nCALL VALUE RECEIVERS WITH POINTER AND POINTER RECEIVERS WITH VALUE:\n===\n")
	pMySquare.Perimeter()
	mySquare.Scale(2)

	fmt.Printf("===\nWHY POINTERS (VALUE LOST):\n===\n")
	mySquare3 := square{Side: 5}
	mySquare3.WrongScale(2)
	mySquare3.Perimeter()
}
