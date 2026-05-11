package main

func main() {
	// 1. operation in parallel
	// res := square.Slice([]int{1, 2, 3, 4, 5})
	// fmt.Println(res)

	// 2. semaphore
	// semaphore.Run([]int{1, 2, 3, 4, 5}, 2)
	// res := semaphore.RunWithCounter([]int{1, 2, 3, 4, 5}, 2)
	// print(res)

	// 3. find and fix data race
	// race.SolutionMutex()
	// race.SolutionAtomic()
	// race.SolutionChan()

	// 4. fan-in channels
	// fanin.RunLoop()
	// fanin.RunSliceLoop()

	// 5. worker pool
	// pool.Demo()

	// 6. http requests concurrent limiter
	// limiter.Demo()
}
