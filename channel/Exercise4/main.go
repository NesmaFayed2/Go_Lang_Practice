package main

import (
	"fmt"
	"time"
)

func main() {
	// Producer generates data into a channel
	d := producer()

	// Consumer reads from the channel with timeout
	consumer(d)
}

// consumer reads messages from the channel
// and uses select to either:
// 1. Get data from channel
// 2. Timeout if waiting too long
func consumer(in chan string) {
	for {
		// timeout after 2ms if no data arrives
		alarm := time.After(2 * time.Millisecond)

		select {
		case m, ok := <-in:
			if !ok {
				fmt.Println("No more data from channel")
				return
			}
			fmt.Printf("Consumer got: %v", m)

		case <-alarm:
			fmt.Println("Timed out waiting for data")
			return
		}
	}
}

// producer sends 20 messages with a pause in the middle
func producer() chan string {
	out := make(chan string)
	go func() {
		count := 1

		// First batch of messages
		for i := 0; i < 10; i++ {
			out <- fmt.Sprintf("Producer message %v\n", count)
			count++
		}

		// Delay (causes consumer timeout)
		time.Sleep(3 * time.Millisecond)

		// Second batch of messages
		for i := 0; i < 10; i++ {
			out <- fmt.Sprintf("Producer message %v\n", count)
			count++
		}

		close(out) // close channel when done
	}()
	return out
}
