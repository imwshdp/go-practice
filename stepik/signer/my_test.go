package main

import (
	"fmt"
	"testing"
)

func TestSimple(t *testing.T) {
	jobs := []job{
		func(in, out chan interface{}) {
			out <- 1
			out <- 2
			out <- 3
		},
		func(in, out chan interface{}) {
			for v := range in {
				out <- v.(int) * 2
			}
		},
		func(in, out chan interface{}) {
			for v := range in {
				fmt.Println(v) // 2, 4, 6
			}
		},
	}
	ExecutePipeline(jobs...)
}
