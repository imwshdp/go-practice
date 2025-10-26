package structs

import "fmt"

type runner interface {
	Run() string
}

type swimmer interface {
	Swim() string
}

type flyer interface {
	Fly() string
}

type ducker interface {
	runner
	swimmer
	flyer
}

type human struct {
	Name string
}

func (h human) Run() string {
	return fmt.Sprintf("Human %s is running", h.Name)
}
func (h human) writeCode() string {
	return fmt.Sprintf("Human %s is coding", h.Name)
}

type duck struct {
	Name, Surname string
}

func (d duck) Run() string {
	return fmt.Sprintf("Duck %s is running", d.Name)
}
func (d duck) Swim() string {
	return fmt.Sprintf("Duck %s is swimming", d.Name)
}

func polymorphismRun(runner runner) {
	fmt.Println(runner.Run())
}

func typeAssertion(runner runner) {
	fmt.Printf("runner (%T): %#v\n", runner, runner)

	// if human, ok := runner.(*Human); ok {
	// 	fmt.Printf("human assertion (%T): %#v\n", human, human)
	// 	fmt.Println(human.writeCode())
	// }

	// if duck, ok := runner.(*Duck); ok {
	// 	fmt.Printf("duck assertion (%T): %#v\n", duck, duck)
	// 	fmt.Println(duck.Swim())
	// }

	// OR

	switch v := runner.(type) {
	case *human:
		fmt.Println(v.writeCode())
	case *duck:
		fmt.Println(v.Swim())
	default:
		fmt.Printf("v (%T): %#v\n", v, v)
	}
}

func typeAssertionAndPolymorphism() {
	fmt.Println("\n---Polymorphism---")

	var runner runner
	fmt.Printf("runner (%T): %#v\n", runner, runner)

	john := &human{Name: "John"}
	runner = john
	polymorphismRun(john)

	donald := &duck{Name: "Donald", Surname: "Duck"}
	runner = donald
	polymorphismRun(donald)

	fmt.Println("\n---Type Assertion---")

	typeAssertion(john)
	typeAssertion(donald)
}

func Interfaces() {
	var runner runner
	fmt.Printf("runner (%T): %#v\n", runner, runner)

	if runner == nil {
		fmt.Println("runner is nil")
	}

	var unnamedRunner *human
	fmt.Printf("\nunnamedRunner (%T): %#v\n", unnamedRunner, unnamedRunner)

	runner = unnamedRunner
	fmt.Printf("runner (%T): %#v\n", runner, runner)
	if runner == nil {
		fmt.Println("runner is nil")
	}

	namedRunner := &human{Name: "Runner"}
	fmt.Printf("\nnamedRunner (%T): %#v\n", namedRunner, namedRunner)

	runner = namedRunner
	fmt.Printf("runner (%T): %#v\n", runner, runner)
	if runner == nil {
		fmt.Println("runner is nil")
	}

	fmt.Println(namedRunner.Run())

	var emptyInterface interface{} = unnamedRunner
	fmt.Printf("\nemptyInterface (%T): %#v\n", emptyInterface, emptyInterface)

	emptyInterface = runner
	fmt.Printf("emptyInterface (%T): %#v\n", emptyInterface, emptyInterface)

	emptyInterface = namedRunner
	fmt.Printf("emptyInterface (%T): %#v\n", emptyInterface, emptyInterface)

	emptyInterface = int64(1)
	fmt.Printf("emptyInterface (%T): %#v\n", emptyInterface, emptyInterface)

	emptyInterface = true
	fmt.Printf("emptyInterface (%T): %#v\n", emptyInterface, emptyInterface)

	typeAssertionAndPolymorphism()
}
