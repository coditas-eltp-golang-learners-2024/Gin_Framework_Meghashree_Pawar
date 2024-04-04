package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func squareWorker(nums []int, ch chan int, start, end int) {
	defer wg.Done()
	for i := start; i < end; i++ {
		ch <- nums[i] * nums[i]
	}
}

func aggregateSquares(ch chan int, numResults int) {
	defer wg.Done()
	sum := 0
	for i := 0; i < numResults; i++ {
		sum += <-ch
	}
	fmt.Println("Aggregated squared result:", sum)
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8}
	numWorkers := 3
	ch := make(chan int)

	wg.Add(numWorkers + 1)

	numElementsPerWorker := (len(nums) + numWorkers - 1) / numWorkers

	for i := 0; i < numWorkers; i++ {
		start := i * numElementsPerWorker
		end := (i + 1) * numElementsPerWorker
		if end > len(nums) {
			end = len(nums)
		}
		go squareWorker(nums, ch, start, end)
	}

	go aggregateSquares(ch, len(nums))

	wg.Wait()
}
