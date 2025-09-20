package main

import "fmt"

var (
	ch chan int
)

func fillCh(c, l int) {
	ch = make(chan int, c)
	for i := 0; i < l; i++ {
		ch <- i
	}
}

func main() {
	//declaring channel
	// var ch chan int
	// fmt.Printf("ch: %v , len: %v, cap: %v\n", ch, len(ch), cap(ch))

	// //sending to channel (dont sent to nil channel)
	// ch <- 4

	// // recieve form nil channel  (dont recieve to nil channel)
	// <-ch

	//making channel un buffer doesnot have capasity
	// ch = make(chan int)
	// fmt.Printf("ch: %v , len: %v, cap: %v\n", ch, len(ch), cap(ch))

	// //dont send  to un buffer channel (without reciever)
	// ch <- 4
	// //dont recieve to un buffer channel (without sender)
	// <-ch

	//actually send
	// ch = make(chan int, 1)

	// fmt.Printf("ch: %v , len: %v, cap: %v\n", ch, len(ch), cap(ch))
	// ch <- 4
	// fmt.Printf("ch: %v , len: %v, cap: %v\n", ch, len(ch), cap(ch))
	// v := <-ch
	// fmt.Printf("data recieved from a channel : %v \n", v)
	// fmt.Printf("ch: %v , len: %v, cap: %v\n", ch, len(ch), cap(ch))

	// //read only  and write only channel
	// chs := make(chan string, 5)
	// //only preduce channel
	// var out chan<- string // only send to this channel
	// out = chs
	// out <- "nesma"
	// out <- "nesma again"
	// // only consume channel

	// var in <-chan string // only recieve from this channel
	// in = chs
	// <-in
	// <-in

	//close channel

	// ch := make(chan string, 5)
	// ch <- "first"
	// close(ch)
	// fmt.Printf("ch: %v , len: %v, cap: %v\n", ch, len(ch), cap(ch))
	// //sending to close channel (panic sending to close channel )
	// // ch <- "seconed"
	// //reciving from close channel (does not panic)
	// <-ch
	// v := <-ch

	// fmt.Printf("ch: %v , len: %v, cap: %v last: %v\n", ch, len(ch), cap(ch), v)

	//tesing for values sent before closure to recive it and defer between empty values and the values that doesnot exist
	// ch := make(chan string, 10)

	// ch <- "first"
	// ch <- "seconed"
	// ch <- ""
	// close(ch)
	// v, ok := <-ch
	// //ok=> not nil
	// fmt.Printf("1st value from channel : %v , ok: %v\n", v, ok)
	// v, ok = <-ch

	// fmt.Printf("2nd value from channel : %v , ok: %v\n", v, ok)
	// v, ok = <-ch

	// fmt.Printf("3rd value from channel : %v , ok: %v\n", v, ok)

	// // it gives us the zero value because this values doesnot exist
	// // but if we try to do this before the channel close it will give us deadlock

	// v, ok = <-ch

	// fmt.Printf("4th value from channel : %v , ok: %v\n", v, ok)

	// dont iterate with cap !
	// fillCh(5, 3)
	// for i := 0; i < cap(ch); i++ {

	// 	fmt.Println(<-ch)
	// }

	// dont iterate with len ! len is decreasing by iteration after recieving from the channel
	// fillCh(5, 3)
	// for i := 0; i < len(ch); i++ {

	// 	fmt.Println(<-ch)
	// }

	//if we take l in a variable before iteration (kinda work)
	// because length can increase while we took it as const
	// fillCh(5, 3)
	// l := len(ch)
	// for i := 0; i < l; i++ {
	// 	fmt.Println(<-ch)

	// }

	//range operator
	//even if it knos the length of the channel it still waiting for additional values
	//dead lock with the same issue that length can  increse
	// fillCh(5, 3)
	// for v := range ch {
	// 	fmt.Println(v)

	// }

	//range operator after close channel (finally the solution)
	// fillCh(5, 3)
	// close(ch)
	// for v := range ch {
	// 	fmt.Println(v)

	// }
	fmt.Println()

}
