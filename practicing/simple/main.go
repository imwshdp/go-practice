package main

import (
	"fmt"
	"tasks-practice/internal/conveyer"
	"tasks-practice/internal/intersection"
	"tasks-practice/internal/mergeChan"
	"tasks-practice/internal/randGen"
	"tasks-practice/internal/workerPool"
	"time"
)

func main() {
	// 1. Пересечение слайсов
	//
	// На вход подаются два неупорядоченных слайса любой длины. Надо написать функцию, которая возвращает их пересечение

	first, second := []int{1, 2, 3}, []int{3, 4, 5}

	setRes := intersection.SolutionWithSet(
		first,
		second,
	)
	fmt.Printf("1.1. intersection (using set) of %v and %v: %v\n", first,
		second, setRes)

	sortRes := intersection.SolutionWithSort(
		first,
		second,
	)
	fmt.Printf("1.2. intersection (using sort) of %v and %v: %v\n", first,
		second, sortRes)

	// 2. Слить N каналов в один

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
	fmt.Printf("2. chan merge result: %v\n", values)

	// 3. Написать генератор случайных чисел

	genSize := 5
	gen := randGen.New(genSize)

	genResult := make([]int, 0, genSize)
	for val := range gen {
		genResult = append(genResult, val)
	}
	fmt.Printf("3. rand generator with chan: %v\n", genResult)

	// 4. Сделать конвейер чисел
	//
	// Даны два канала.
	// В первый пишутся числа.
	// Нужно, чтобы числа читались из первого по мере поступления,
	// что-то с ними происходило (допустим, возводились в квадрат) и результат записывался во второй канал.

	conveyerSize := 5
	in := make(chan int)

	go func() {
		for inx := range conveyerSize {
			in <- inx
		}
		close(in)
	}()

	out := conveyer.New(in, func(val int) int {
		return val * val
	})

	conveyerResult := make([]int, 0, conveyerSize)
	for val := range out {
		conveyerResult = append(conveyerResult, val)
	}
	fmt.Printf("4. conveyer result: %v\n", conveyerResult)

	// 5. Написать WorkerPool с заданной функцией
	//
	// Нам нужно разбить процессы на несколько горутин.
	// При этом не создавать новую горутину каждый раз, а просто переиспользовать уже имеющиеся.
	//
	// package workerPool
	const numJobs = 5

	jobs := make(chan *workerPool.Job, numJobs)
	multiplier := func(x int) int {
		return x * 10
	}

	poolResults := workerPool.New(jobs, numJobs)

	go func() {
		defer close(jobs)

		for inx := range numJobs {
			jobs <- &workerPool.Job{
				Value: inx + 1,
				Func:  multiplier,
			}
			time.Sleep(time.Second)
		}
	}()

	fmt.Println("5. worker pool result:")
	for val := range poolResults {
		fmt.Println(val)
	}
}
