package main

import "fmt"

type Square struct {
	Side int
}

func (s Square) Perimeter() {
	fmt.Printf("Square s(%T): %#v\n", s, s)
	fmt.Printf("Perimeter: %v\n\n", s.Side*4)
}

func (s *Square) Scale(multiplier int) {
	fmt.Printf("Square s(%T): %#v\n", s, s)
	s.Side *= multiplier
	fmt.Printf("Scaled: %v\n\n", s.Side)
}

func (s Square) WrongScale(multiplier int) {
	fmt.Printf("Square s(%T): %#v\n", s, s)
	s.Side *= multiplier
	fmt.Printf("Wrong scaled s(%T): %#v\n\n", s, s)
}

func Methods() {
	square := Square{Side: 10}
	pSquare := &square

	square2 := Square{Side: 2}

	fmt.Printf("===\nVALUE AND POINTER RECEIVERS:\n===\n")
	square.Perimeter()
	square2.Perimeter()
	pSquare.Scale(2)

	fmt.Printf("===\nCALL VALUE RECEIVERS WITH POINTER AND POINTER RECEIVERS WITH VALUE:\n===\n")
	pSquare.Perimeter()
	square.Scale(2)

	fmt.Printf("===\nWHY POINTERS (VALUE LOST):\n===\n")
	square3 := Square{Side: 5}
	square3.WrongScale(2)
	square3.Perimeter()
}
