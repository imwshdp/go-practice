package main

import (
	"fmt"
	"tasks-practice/internal/conveyer"
	"tasks-practice/internal/intersection"
	"tasks-practice/internal/mergeChan"
	"tasks-practice/internal/randGen"
)

func main() {
	// Пересечение слайсов
	//
	// На вход подаются два неупорядоченных слайса любой длины. Надо написать функцию, которая возвращает их пересечение
	//
	// package intersection
	sortRes := intersection.SolutionWithSort(
		[]int{1, 2, 3},
		[]int{3, 4, 5},
	)
	fmt.Println(sortRes)

	setRes := intersection.SolutionWithSet(
		[]int{1, 2, 3},
		[]int{3, 4, 5},
	)
	fmt.Println(setRes)

	// Слить N каналов в один
	//
	// package mergeChan
	ch1 := make(chan int, 2)
	ch2 := make(chan int, 2)
	ch3 := make(chan int, 2)

	ch1 <- 1
	ch1 <- 2
	close(ch1)

	ch2 <- 3
	ch2 <- 4
	close(ch2)

	ch3 <- 5
	ch3 <- 6
	close(ch3)

	result := mergeChan.Merge(ch1, ch2, ch3)

	var values []int
	for val := range result {
		values = append(values, val)
	}
	fmt.Println(values)

	// Написать генератор случайных чисел
	//
	// package randGen
	gen := randGen.New(5)

	genResult := make([]int, 0, 5)
	for val := range gen {
		genResult = append(genResult, val)
	}
	fmt.Println(genResult)

	// Сделать конвейер чисел
	//
	// Даны два канала.
	// В первый пишутся числа.
	// Нужно, чтобы числа читались из первого по мере поступления,
	// что-то с ними происходило (допустим, возводились в квадрат) и результат записывался во второй канал.
	//
	// package conveyer
	in := make(chan int)

	go func() {
		for inx := range 5 {
			in <- inx
		}
		close(in)
	}()

	out := conveyer.New(in, func(val int) int {
		return val * val
	})

	for val := range out {
		fmt.Print(val)
	}
}
