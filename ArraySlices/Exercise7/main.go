package main

/*

 */

import (
	"fmt"

	"github.com/striversity/glft/shared/input"
)

func main() {

	var a = [5]int{10, 20, 30, 40, 50}

	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	// create
	s1 := a[1:4]
	s1[0] = -1

	//shallow copy
	s3 := s1
	s3[1] = -1
	fmt.Println(a)

	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	//expand
	var s2 []string
	s2 = append(s2, "cairo", "washnton")

	fmt.Println(s2)

	//deep copy
	s4 := make([]string, 10)

	copy(s4, s2)

	fmt.Println(s4)

	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>")

	var slice = make([]int, 5)
	for i := range slice {
		slice[i] = input.GetRandInt(10, 20)
	}
	fmt.Printf("Slice is %v", slice)

	doubled := myDouble(slice...)
	fmt.Println("")

	fmt.Printf("Dobled slice : %v", doubled)

	even := even(slice...)
	fmt.Println("")

	fmt.Printf("even slice : %v", even)

	fmt.Printf("")
}

func myDouble(num ...int) (doubleSlice []int) {

	doubleSlice = make([]int, len(num))
	for i, v := range num {
		doubleSlice[i] = v * 2
	}
	return

}

func even(num ...int) (even []int) {
	even = []int{}
	for _, v := range num {
		if v%2 != 0 {
			continue
		}
		even = append(even, v)
	}

	return

}
