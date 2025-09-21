package main

import (
	"fmt"
	"sync"
	"time"
)

const numPhilosophers = 5

type Chopstick struct{ sync.Mutex }

type Philosopher struct {
	id             int
	leftChopstick  *Chopstick
	rightChopstick *Chopstick
}

func (p Philosopher) eat(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 3; i++ { // Each philosopher eats 3 times
		// Pick up left then right (to avoid deadlock, enforce order by ID)
		if p.id%2 == 0 {
			p.leftChopstick.Lock()
			p.rightChopstick.Lock()
		} else {
			p.rightChopstick.Lock()
			p.leftChopstick.Lock()
		}

		// Eating
		fmt.Printf("Philosopher %d is eating\n", p.id)
		time.Sleep(time.Millisecond * 500)

		// Put down chopsticks
		p.leftChopstick.Unlock()
		p.rightChopstick.Unlock()

		fmt.Printf("Philosopher %d finished eating\n", p.id)
		time.Sleep(time.Millisecond * 500) // Thinking
	}
}

func main() {
	// Create chopsticks
	chopsticks := make([]*Chopstick, numPhilosophers)
	for i := 0; i < numPhilosophers; i++ {
		chopsticks[i] = new(Chopstick)
	}

	// Create philosophers
	philosophers := make([]*Philosopher, numPhilosophers)
	for i := 0; i < numPhilosophers; i++ {
		philosophers[i] = &Philosopher{
			id:             i,
			leftChopstick:  chopsticks[i],
			rightChopstick: chopsticks[(i+1)%numPhilosophers],
		}
	}

	// Run all philosophers concurrently
	var wg sync.WaitGroup
	for i := 0; i < numPhilosophers; i++ {
		wg.Add(1)
		go philosophers[i].eat(&wg)
	}
	wg.Wait()
}
