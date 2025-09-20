package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	chCap = 10
)

var (
	s = rand.NewSource(time.Now().Unix())

	r = rand.New(s)
)

func main() {

	//using generator returning channel from a function
	d := generator()
	//processor:
	d = counter(d)
	d = adder(d, 5)
	//consuming after proseccing
	consumer(d)

}

// adder for adding values to every elent in the channel before consume it
func adder(in chan int, num int) (out chan int) {

	out = make(chan int, chCap)
	for v := range in {
		out <- v + num
	}
	close(out)
	return
}

// we take channel in and copies to channel out

func counter(in chan int) (out chan int) {

	out = make(chan int, chCap)
	count := 0

	for v := range in {
		out <- v
		count++
	}
	close(out)
	fmt.Printf("Counted elements : %v\n", count)

	return
}

// generate channel that intailize channels and send random values to it
func generator() (out chan int) {
	fmt.Println("Generator of random int")
	out = make(chan int, chCap)
	producer(out)
	return

}

func producer(nums chan int) {
	l := r.Int()%cap(nums) + 1
	for i := 0; i < l; i++ {
		nums <- r.Int() % 200
	}
	//close to ssay to the consumer i close it dont expect more form it
	//so you can range over it
	close(nums)

}

func consumer(nums chan int) {
	for v := range nums {
		fmt.Printf("Consumer got : %v\n", v)
	}

}
