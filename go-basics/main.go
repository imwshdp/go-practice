package main

import (
	"go-basics/basics"
	"go-basics/collections"
	"go-basics/concurrency"
	"go-basics/funcs"
	"go-basics/generics"
	"go-basics/modules"
	"go-basics/pointers"
	"go-basics/statements"
	"go-basics/structs"
)

func main() {
	basics.HelloWorld()
	basics.Vars()

	funcs.Funcs()
	funcs.FuncsAdv()

	statements.Conditions()
	statements.Loops()

	pointers.Pointers()

	structs.Structs()
	structs.Methods()
	structs.Interfaces()
	structs.Embedding()

	collections.Arrays()
	collections.Slices()
	collections.Maps()

	concurrency.Goroutines()
	concurrency.Defer()
	concurrency.Panic()
	concurrency.Sync()
	concurrency.Channels()
	concurrency.Select()
	concurrency.Context()
	concurrency.ErrGroup()
	concurrency.Atomic()

	generics.Generics()

	modules.Modules()
}
