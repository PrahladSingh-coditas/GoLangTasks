package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func squareWorker(numList []int, ch chan int) {
	defer wg.Done() // indicate that this goroutine is done
	defer close(ch) // close the channel when all the send operations are completed

	for _, number := range numList {
		ch <- (number * number)
	}
}

func aggregateSquares(ch chan int) {
	defer wg.Done()

	sum := 0
	for receivingVal := range ch {
		sum += receivingVal
	}

	fmt.Println("The aggregation of the square numbers is:", sum)
}

func main() {
	start := time.Now() // Record the start time

	numList := []int{1, 2, 3}
	ch := make(chan int) // Create channel in main function scope

	wg.Add(2) // Add two goroutines to wait for

	go squareWorker(numList, ch)
	go aggregateSquares(ch)

	fmt.Println(time.Since(start))

	wg.Wait()
}
