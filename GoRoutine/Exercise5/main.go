package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	s = rand.NewSource(time.Now().Unix())
	r = rand.New(s)
)

func main() {
	var wg sync.WaitGroup

	start := time.Now()

	// launch 5 producers with shared WaitGroup
	for i := 1; i <= 5; i++ {
		producer(&wg, i)
	}

	wg.Wait()

	elapse := time.Since(start)
	fmt.Printf("Challenge 3 finished in %v\n", elapse)
}

// Correct: pass *sync.WaitGroup, not copy
func producer(wg *sync.WaitGroup, id int) {
	wg.Add(1)
	go func() {
		n := (r.Int() % 1000) + 1
		d := time.Duration(n) * time.Millisecond
		time.Sleep(d)
		fmt.Printf("Producer #%v ran for %v\n", id, d)
		wg.Done()
	}()
}
