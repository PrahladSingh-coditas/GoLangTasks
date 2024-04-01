package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func squareWorker(numList []int) {
	ch := make(chan int)

	wg.Add(1)
	go func() {
		defer close(ch) //Basically cloase the channel when all the send operations are completed
		for _, number := range numList {
			ch <- (number * number)
		}

	}()

	aggregateSquares(ch)
}

func aggregateSquares(ch chan int) {
	defer wg.Done()

	sum := 0

	//The range creates an implicit loop that iterates over the channel's values until the channel is closed.
	//Here the range does not work the way it used to work for slices
	for receivingVal := range ch { //Here we are receiving the output from the channel
		sum += receivingVal
	}

	fmt.Println("The aggregation of the square numbers is: ", sum)
}

func main() {
	numList := []int{1, 2, 3}
	squareWorker(numList)

	wg.Wait()
}
