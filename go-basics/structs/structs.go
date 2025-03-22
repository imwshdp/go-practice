package structs

import (
	"fmt"
	"time"
)

type OurString string
type OurInt int

type Person struct {
	Name string
	Age  int
}

func Structs() {
	var customString OurString
	fmt.Printf("var customString (%T): %#v\n", customString, customString)

	customString = "Hello, I'am a custom string"
	fmt.Printf("var customString (%T): %#v\n\n", customString, customString)

	customInt := OurInt(10)
	fmt.Printf("var customInt (%T): %#v\n", customInt, customInt)
	fmt.Printf("var int(customInt) (%T): %#v\n\n", int(customInt), int(customInt))

	var John Person
	fmt.Printf("var John (%T): %#v\n", John, John)

	John = Person{}
	fmt.Printf("var John{} (%T): %#v\n\n", John, John)

	John.Name = "John"
	John.Age = 20
	fmt.Printf("John: %#v\n", John)

	Brad := Person{
		Name: "Brad",
		Age:  30,
	}
	fmt.Printf("Brad: %#v\n", Brad)

	Sam := Person{"Sam", 25}
	fmt.Printf("Sam: %#v\n\n", Sam)

	pointerSam := &Sam
	fmt.Println("Age of Sam:", (*pointerSam).Age)
	fmt.Println("Age of Sam from pointer:", pointerSam.Age)

	pointerKate := &Person{"Kate", 18}
	fmt.Printf("pointerKate: %#v and her age from pointer: %#v\n\n", pointerKate, pointerKate.Age)

	unnamedStruct := struct {
		Name, LastName, BirthDate string
	}{
		Name:      "NoName",
		LastName:  "User",
		BirthDate: fmt.Sprintf("%s", time.Now()),
	}

	fmt.Printf("var unnamedStruct (%T): %v", unnamedStruct, unnamedStruct)
}
