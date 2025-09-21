package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

var (
	r  = rand.New(rand.NewSource(time.Now().Unix()))
	wg sync.WaitGroup
)

func main() {

	//signal channel
	done := make(chan struct{})

	d := intGen(done)

	//consuming ...
	consume(d)

	time.Sleep(2 * time.Second)
	close(done)
	wg.Wait()

}

// consume and count how many data consumed in what time
func consume(nums chan int) {
	wg.Add(1)

	go func() {
		defer wg.Done()

		log.Info("")
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
