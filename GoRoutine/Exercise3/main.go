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
	// Run sequential
	startSeq := time.Now()
	for i := 1; i <= 4; i++ {
		producerSequential(i)
	}
	elapseSeq := time.Since(startSeq)
	fmt.Printf("Sequential run took %v\n\n", elapseSeq)

	// Run concurrent
	startConc := time.Now()
	wg.Add(4)
	for i := 1; i <= 4; i++ {
		go producerConcurrent(i)
	}
	wg.Wait()
	elapseConc := time.Since(startConc)
	fmt.Printf("Concurrent run took %v\n", elapseConc)
}

func producerSequential(id int) {
	n := (r.Int() % 1000) + 1
	d := time.Duration(n) * time.Millisecond
	time.Sleep(d)
	fmt.Printf("[Sequential] Producer #%v finished in %v\n", id, d)
}

func producerConcurrent(id int) {
	n := (r.Int() % 1000) + 1
	d := time.Duration(n) * time.Millisecond
	time.Sleep(d)
	fmt.Printf("[Concurrent] Producer #%v finished in %v\n", id, d)
	wg.Done()
}
