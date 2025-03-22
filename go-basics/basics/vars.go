package basics

import (
	"fmt"
	"unsafe"
)

var globalVar string
var anotherGlobalVar string

func Vars() {
	var str string
	fmt.Printf("var str (%T): %v\n", str, str)

	var boolean bool
	fmt.Printf("var boolean (%T): %v\n", boolean, boolean)

	newBoolean := true
	fmt.Printf("var newBoolean (%T): %v\n", newBoolean, newBoolean)

	fmt.Printf("var globalVar (%T): %v\n", globalVar, globalVar)

	anotherGlobalVar := "scope shadowing"
	fmt.Printf("var anotherGlobalVar (%T): %v\n", anotherGlobalVar, anotherGlobalVar)

	const constant = "const"
	fmt.Printf("var constant (%T): %v\n", constant, constant)

	world := "world"
	helloWorld := fmt.Sprintf("Hello, %s!", world)
	_ = helloWorld
	fmt.Printf("Formatted string: %v\n", helloWorld)

	var num1 int64 = 15
	var num2 int = 15

	fmt.Printf("\n%v + %v = %v\n", num1, num2, num1+int64(num2))

	fmt.Println("uint8 size:", unsafe.Sizeof(uint8(1)))
	fmt.Println("uint32 size:", unsafe.Sizeof(uint32(1)))
	fmt.Println("string size:", unsafe.Sizeof("string"))
	fmt.Println("bool size:", unsafe.Sizeof(true))
}
