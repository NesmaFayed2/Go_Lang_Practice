package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
Challenge 1 – Mix Sleep & WaitGroup

Start 3 producers with WaitGroup.

Also keep a time.Sleep(2 * time.Second) in main().

Observe: does WaitGroup still wait correctly even if you add extra sleep?

*/

var (
	wg sync.WaitGroup

	s = rand.NewSource(time.Now().Unix())
	r = rand.New(s)
)

func main() {

	start := time.Now()

	wg.Add(3)

	go producer(1)
	go producer(2)
	go producer(3)

	wg.Wait()

	elapse := time.Since(start)
	fmt.Printf(" Challenge 1 took %v\n", elapse)

}

func producer(id int) {
	n := (r.Int() % 1000) + 1
	d := time.Duration(n) * time.Millisecond
	time.Sleep(d)
	fmt.Printf("Producer #%v finished in %v\n", id, d)

	wg.Done()
}
