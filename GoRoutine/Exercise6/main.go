package main

/*
Challenge 1: Race Condition Demonstration

Write a program that:

Starts 1000 goroutines.

Each goroutine increments a shared counter (counter++).

Print the final counter value.

Expected: The result will not always be 1000 because of the race condition.
*/

import (
	"fmt"
	"sync"
)

var (
	counter int
	wg      sync.WaitGroup
)

func main() {

	numberOfWorkers := 1000
	wg.Add(numberOfWorkers)
	for i := 0; i < numberOfWorkers; i++ {
		go func() {
			counter++
			wg.Done()

		}()
	}
	wg.Wait()
	fmt.Printf("Final Counter (Race Condition): %v\n", counter)

}
