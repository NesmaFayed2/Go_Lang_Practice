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
	// Signal channel to stop generators
	done := make(chan struct{})

	// Two generators (numbers + letters)
	numbers := intGen(done)
	letters := alphaGen(done)

	// Fan-In: merge numbers + letters into one channel
	merged := fanIn(numbers, letters)

	// Print everything from merged channel
	printer(merged)

	// Run for a short time, then stop
	time.Sleep(50 * time.Millisecond)
	close(done)
	wg.Wait()
}

// Fan-In: merge two channels into one
func fanIn(nums chan int, letters chan rune) chan string {
	out := make(chan string)
	wg.Add(1)

	go func() {
		defer wg.Done()
		defer close(out)
		for {
			select {
			case n, ok := <-nums:
				if ok {
					out <- fmt.Sprintf("Number: %d", n)
				} else {
					return
				}
			case l, ok := <-letters:
				if ok {
					out <- fmt.Sprintf("Letter: %c", l)
				} else {
					return
				}
			}
		}
	}()
	return out
}

// Printer: just consumes and prints
func printer(in chan string) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range in {
			fmt.Println("Consumer got ->", v)
		}
	}()
}

// Random numbers generator
func intGen(done chan struct{}) chan int {
	out := make(chan int)
	wg.Add(1)

	go func() {
		defer wg.Done()
		defer close(out)
		for {
			select {
			case <-done:
				return
			case out <- r.Int() % 100:
			}
		}
	}()
	return out
}

// Random letters generator
func alphaGen(done chan struct{}) chan rune {
	out := make(chan rune)
	wg.Add(1)

	go func() {
		defer wg.Done()
		defer close(out)
		letters := []rune("abcdefghijklmnopqrstuvwxyz")
		for {
			select {
			case <-done:
				return
			case out <- letters[r.Int()%len(letters)]:
			}
		}
	}()
	return out
}
