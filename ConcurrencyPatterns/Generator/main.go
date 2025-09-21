package main

import (
	"fmt"
	"math/rand"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().Unix()))

func main() {

	done := make(chan struct{})

	d := intGen(done)

	//consuming
	for i := 0; i < 10; i++ {
		fmt.Println(<-d)
	}

	close(done)

}

// intGen  take send only channel and  return recieve only channel

func intGen(done chan struct{}) chan int {
	out := make(chan int)

	// signal <-done
	// generate data tell done
	go func() {
		defer close(out)
		for {
			select {
			case <-done:
				return
			case out <- r.Int() % 200:

			}
		}
	}()

	return out

}
