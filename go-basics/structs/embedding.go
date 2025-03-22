package structs

import "fmt"

type humanPerson struct {
	Name string
	Age  int
}

func (h humanPerson) printName() {
	fmt.Println(h.Name)
}

type workExperience struct {
	Name string
	Age  int
}

type woodBuilder struct {
	humanPerson
	Name string
}

func (wb woodBuilder) build() {
	fmt.Println("Building from woods")
}

type woodBuilderWithWE struct {
	humanPerson
	workExperience
	Name string
}

type brickBuilder struct {
	humanPerson
	Name string
}

func (bb brickBuilder) build() {
	fmt.Println("Building from bricks")
}

type builder interface {
	build()
}

type building struct {
	builder
	name string
}

func useCase() {
	woodBuilding := building{
		builder: woodBuilder{
			humanPerson{
				Name: "Bob",
				Age:  25,
			},
			"Bob The Builder",
		},
		name: "Wood building",
	}
	woodBuilding.build()

	brickBuilding := building{
		builder: brickBuilder{
			humanPerson{
				Name: "John",
				Age:  30,
			},
			"John The Builder",
		},
		name: "Brick building",
	}
	brickBuilding.build()
}

func Embedding() {
	builder := woodBuilder{
		humanPerson{
			Name: "John",
			Age:  25,
		},
		"builder",
	}

	fmt.Printf("builder(%T): %#v\n", builder, builder)

	// shorthands
	fmt.Println(builder.humanPerson.Age)
	fmt.Println(builder.Age)

	// shadowing
	fmt.Println(builder.Name)
	fmt.Println(builder.humanPerson.Name)

	// methods
	builder.printName()

	// colliding
	builder2 := woodBuilderWithWE{
		humanPerson{
			Name: "John",
			Age:  25,
		},
		workExperience{
			Name: "Builder",
			Age:  1,
		},
		"builder",
	}
	// fmt.Println(builder2.Age)
	fmt.Println(builder2.humanPerson.Age)
	fmt.Println(builder2.workExperience.Age)

	// embedding use case
	useCase()
}
