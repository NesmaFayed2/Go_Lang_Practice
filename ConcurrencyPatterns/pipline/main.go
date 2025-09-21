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
	d = passThroughCounter(d)
	d = filter(d, func(x int) bool {
		if x%2 == 0 {
			return true
		}
		return false
	})

	//consuming ...
	consume(d)

	time.Sleep(2 * time.Second)
	close(done)
	wg.Wait()

}

func passThroughCounter(nums chan int) (out chan int) {
	wg.Add(1)
	out = make(chan int)

	go func() {
		defer wg.Done()
		defer close(out)
		start := time.Now()
		var count int
		for v := range nums {
			count++
			out <- v
		}
		fmt.Printf("Counter Process : %v in %v seconeds", count, time.Since(start))

	}()
	return out

}

// filter
func filter(in chan int, even func(x int) bool) (out chan int) {

	wg.Add(1)
	out = make(chan int)

	if even == nil {
		even = func(int) bool { return true }
	}

	go func() {
		defer close(out)
		defer wg.Done()
		for v := range in {
			if even(v) {
				out <- v
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
		fmt.Printf("Counter Process : %v in %v seconeds", count, time.Since(start))

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
