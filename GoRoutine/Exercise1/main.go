package main

import (
	"fmt"
	"time"
)

func main() {

	go producer(1)
	go producer(2)
	go func() {
		fmt.Println("hello")
	}()

	go func() {
		fmt.Println("world")
	}()

	time.Sleep(1 * time.Second)

}

func producer(id int) {
	for i := 0; i < 5; i++ {
		fmt.Printf("message from %v: %v\n", id, i)

		//creates natural context switches: while one goroutine waits, the scheduler runs the other.
		//By adding time.Sleep, we create a blocking point, so the scheduler naturally gives a turn to the other goroutines.
		time.Sleep(100 * time.Millisecond)

	}
}
