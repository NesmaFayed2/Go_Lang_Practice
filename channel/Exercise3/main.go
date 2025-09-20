package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

var (
	s          = rand.NewSource(time.Now().Unix())
	r          = rand.New(s)
	wgProducer sync.WaitGroup
	wgConsumer sync.WaitGroup
)

func main() {

	//buffered
	// d := make(chan string, 5)
	// go producer(1, d)

	// consumer(d)

	//unbuffered channel multiple producer and single consumer
	// d := make(chan string)
	// producer(1, d)
	// producer(2, d)
	// producer(3, d)
	// go consumer(d)
	// wgProducer.Wait()
	// close(d)
	// wgConsumer.Wait()

	//!! multiple producer and multiple consumer

	d := make(chan string)
	producer(1, d)
	producer(2, d)
	producer(3, d)
	consumer(1, d)
	consumer(2, d)
	wgProducer.Wait()
	close(d)
	wgConsumer.Wait()

}

func consumer(id int, channel chan string) {
	wgConsumer.Add(1)
	go func() {

		// count := 0
		// for v := range channel {
		// 	count++
		// 	fmt.Printf("Consumer: %v - got : %v\n", id, v)
		// }

		// if count == 0 {
		// 	fmt.Println("No Data recieved")
		// 	return
		// }
		// fmt.Printf("Counted : %v", count)

		db := make(map[string]int)
		var fields []string
		for v := range channel {
			fields = strings.Split(v, ",")
			db[fields[0]]++
		}
		for k, v := range db {

			fmt.Printf("Consumer %v - process %v items for producer %v \n ", id, v, k)
		}

		wgConsumer.Done()
	}()
}

func producer(id int, channel chan string) {
	wgProducer.Add(1)

	go func() {
		l := r.Int() % 5
		for i := 0; i < l; i++ {
			channel <- fmt.Sprintf("Producer: %v, item: %v", id, i)
		}
		wgProducer.Done()
	}()

}
