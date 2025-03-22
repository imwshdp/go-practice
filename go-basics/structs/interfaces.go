package structs

import "fmt"

type Runner interface {
	Run() string
}

type Swimmer interface {
	Swim() string
}

type Flyer interface {
	Fly() string
}

type Ducker interface {
	Runner
	Swimmer
	Flyer
}

type Human struct {
	Name string
}

func (h Human) Run() string {
	return fmt.Sprintf("Human %s is running", h.Name)
}
func (h Human) writeCode() string {
	return fmt.Sprintf("Human %s is coding", h.Name)
}

type Duck struct {
	Name, Surname string
}

func (d Duck) Run() string {
	return fmt.Sprintf("Duck %s is running", d.Name)
}
func (d Duck) Swim() string {
	return fmt.Sprintf("Duck %s is swimming", d.Name)
}

func polymorphismRun(runner Runner) {
	fmt.Println(runner.Run())
}

func typeAssertion(runner Runner) {
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
	case *Human:
		fmt.Println(v.writeCode())
	case *Duck:
		fmt.Println(v.Swim())
	default:
		fmt.Printf("v (%T): %#v\n", v, v)
	}
}

func typeAssertionAndPolymorphism() {
	fmt.Println("\n---Polymorphism---")

	var runner Runner
	fmt.Printf("runner (%T): %#v\n", runner, runner)

	john := &Human{Name: "John"}
	runner = john
	polymorphismRun(john)

	donald := &Duck{Name: "Donald", Surname: "Duck"}
	runner = donald
	polymorphismRun(donald)

	fmt.Println("\n---Type Assertion---")

	typeAssertion(john)
	typeAssertion(donald)
}

func Interfaces() {
	var runner Runner
	fmt.Printf("runner (%T): %#v\n", runner, runner)

	if runner == nil {
		fmt.Println("runner is nil")
	}

	var unnamedRunner *Human
	fmt.Printf("unnamedRunner (%T): %#v\n", unnamedRunner, unnamedRunner)

	runner = unnamedRunner
	fmt.Printf("runner (%T): %#v\n", runner, runner)
	if runner == nil {
		fmt.Println("runner is nil")
	}

	namedRunner := &Human{Name: "Runner"}
	fmt.Printf("namedRunner (%T): %#v\n", namedRunner, namedRunner)

	runner = namedRunner
	fmt.Printf("runner (%T): %#v\n", runner, runner)
	if runner == nil {
		fmt.Println("runner is nil")
	}

	namedRunner.Run()

	var emptyInterface interface{} = unnamedRunner
	fmt.Printf("emptyInterface (%T): %#v\n", emptyInterface, emptyInterface)

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
