package main

import (
	"fmt"
	"sort"
	"strings"
	"sync"
)

const MultiHashSteps = 6

var md5MutexChan = make(chan struct{}, 1)

func SingleHash(in, out chan interface{}) {
	wg := &sync.WaitGroup{}

	for input := range in {
		wg.Add(1)

		go func(inputStr string) {
			defer wg.Done()

			md5MutexChan <- struct{}{}
			md5 := DataSignerMd5(inputStr)
			<-md5MutexChan

			crc32Chan := make(chan string)
			crc32md5Chan := make(chan string)

			go func(crc32md5Chan chan<- string) { crc32md5Chan <- DataSignerCrc32(md5) }(crc32md5Chan)
			go func(crc32Chan chan<- string) { crc32Chan <- DataSignerCrc32(inputStr) }(crc32Chan)

			out <- (<-crc32Chan) + "~" + (<-crc32md5Chan)
		}(fmt.Sprint(input))
	}

	wg.Wait()
}

func MultiHash(in, out chan interface{}) {
	wg := &sync.WaitGroup{}

	for input := range in {
		wg.Add(1)

		go func(inputStr string) {
			defer wg.Done()

			resArray := [MultiHashSteps]string{}

			innerWg := &sync.WaitGroup{}
			innerWg.Add(MultiHashSteps)

			for th := 0; th < MultiHashSteps; th++ {
				var mu sync.Mutex

				go func(th int) {
					defer innerWg.Done()

					res := DataSignerCrc32(fmt.Sprint(th) + inputStr)
					mu.Lock()
					resArray[th] = res
					mu.Unlock()
				}(th)
			}

			innerWg.Wait()

			result := ""
			for inx := 0; inx < MultiHashSteps; inx++ {
				result += resArray[inx]
			}
			out <- result

		}(fmt.Sprint(input))
	}

	wg.Wait()
}

func CombineResults(in, out chan interface{}) {
	resultSlice := make([]string, 0, MaxInputDataLen)

	for inputData := range in {
		strInput := fmt.Sprint(inputData)
		resultSlice = append(resultSlice, strInput)
	}

	sort.Strings(resultSlice)
	out <- strings.Join(resultSlice, "_")
}

func ExecutePipeline(jobs ...job) {
	channels := make([]chan interface{}, len(jobs)+1)
	for inx := range channels {
		channels[inx] = make(chan interface{}, MaxInputDataLen)
	}

	wg := &sync.WaitGroup{}

	for inx, curJob := range jobs {
		wg.Add(1)

		go func(inx int, curJob job) {
			defer func() {
				close(channels[inx+1])
				wg.Done()
			}()

			curJob(channels[inx], channels[inx+1])
		}(inx, curJob)
	}

	close(channels[0])
	wg.Wait()
}
