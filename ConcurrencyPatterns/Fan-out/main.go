package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	r  = rand.New(rand.NewSource(time.Now().Unix()))
	wg sync.WaitGroup
)

func main() {

	//signal channel
	done := make(chan struct{})

	d := intGen(done)
	s0, s1, s2 := fanOut3(d)

	//consuming ...
	consume(s0)
	consume(s1)
	consume(s2)

	time.Sleep(2 * time.Second)
	close(done)
	wg.Wait()

}

func fanOut3(in chan int) (out0, out1, out2 chan int) {

	wg.Add(1)
	out0 = make(chan int)
	out1 = make(chan int)
	out2 = make(chan int)
	go func() {
		defer wg.Done()
		defer close(out0)
		defer close(out1)
		defer close(out2)
		for v := range in {

			select {
			case out0 <- v:
			case out1 <- v:
			case out2 <- v:

			}
		}

	}()
	return
}

// consume and count how many data consumed in what time
func consume(nums chan int) {
	wg.Add(1)

	go func() {
		defer wg.Done()
		start := time.Now()
		var count int
		for range nums {
			count++
		}
		fmt.Printf("Counter Process : %v in %v seconeds \n", count, time.Since(start))

	}()

}

func intGen(done chan struct{}) chan int {
	wg.Add(1)
	out := make(chan int)

	go func() {
		defer close(out)
		defer wg.Done()
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
