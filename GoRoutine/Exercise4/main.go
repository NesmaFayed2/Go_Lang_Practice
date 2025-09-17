package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	s  = rand.NewSource(time.Now().Unix())
	r  = rand.New(s)
	wg sync.WaitGroup
)

func main() {
	start := time.Now()

	// Correct: wg.Add matches number of goroutines
	wg.Add(2)
	go producer(1)
	go producer(2)

	wg.Wait()

	elapse := time.Since(start)
	fmt.Printf("Challenge 1 finished in %v\n", elapse)
}

func producer(id int) {
	n := (r.Int() % 1000) + 1
	d := time.Duration(n) * time.Millisecond
	time.Sleep(d)
	fmt.Printf("Producer #%v ran for %v\n", id, d)
	wg.Done()
}
