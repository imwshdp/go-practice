package structs

import (
	"fmt"
	"time"
)

type ourString string
type ourInt int

type person struct {
	Name string
	Age  int
}

func Structs() {
	var customString ourString
	fmt.Printf("var customString (%T): %#v\n", customString, customString)

	customString = "Hello, I'am a custom string"
	fmt.Printf("var customString (%T): %#v\n\n", customString, customString)

	customInt := ourInt(10)
	fmt.Printf("var customInt (%T): %#v\n", customInt, customInt)
	fmt.Printf("var int(customInt) (%T): %#v\n\n", int(customInt), int(customInt))

	var John person
	fmt.Printf("var John (%T): %#v\n", John, John)

	John = person{}
	fmt.Printf("var John{} (%T): %#v\n\n", John, John)

	John.Name = "John"
	John.Age = 20
	fmt.Printf("John: %#v\n", John)

	Brad := person{
		Name: "Brad",
		Age:  30,
	}
	fmt.Printf("Brad: %#v\n", Brad)

	Sam := person{"Sam", 25}
	fmt.Printf("Sam: %#v\n\n", Sam)

	pointerSam := &Sam
	fmt.Println("Age of Sam:", (*pointerSam).Age)
	fmt.Println("Age of Sam from pointer:", pointerSam.Age)

	pointerKate := &person{"Kate", 18}
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
