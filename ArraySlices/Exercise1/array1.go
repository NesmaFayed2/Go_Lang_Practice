package main

/*# Section 03 Exercise 01 - Declaring and Using Arrays

## Create a go program which does the following

### TODO 1

1. Declares an array of 10 to 20 elements of type float64 using a const for the array size

### TODO 2

1. Initialize the array with random currency values using 'input.GetRandFloat()'.
	* TIP:
		* See Section 02 - Lecture 10 for the currency type
		* Use a 'for' loop to initialize your array values
		* The 'input' package is github.com/striversity/glft/shared/input
		and was covered in Sec01-Lec12.

### TODO 3

1. Calculate total, max, min, avg*/

import (
	"fmt"

	"github.com/striversity/glft/shared/input"
)

type (
	Currency float64
)

func (c Currency) String() (s string) {
	s = fmt.Sprintf("$%.2f", c)
	return

}

const size = 10

func main() {

	var myArray [size]Currency

	//intialize the array
	for i := range myArray {
		myArray[i] = Currency(input.GetRandFloat(0, 10))
	}

	fmt.Println(myArray)

	// max min avg total

	var total, avg, max, min Currency

	for _, v := range myArray {

		total += v

		if v < min || min == 0 {
			min = v
		}
		if v > max {
			max = v
		}

	}

	avg = total / size

	fmt.Printf("total : %v , avg : %v , max : %v , min : %v", total, avg, max, min)

	// var arr [3]int

	// arr[0]=3
	// arr[1]=4
	// arr[2]=5

	// var sum  int
	// for i:=0;i<len(arr);i++{
	//  sum+=arr[i]

	// }

	// fmt.Printf("Average is %.2f",float64(sum/len(arr)))
	// fmt.Println("")

	// //range

	// var sum2 int
	// for _,v:= range arr{
	//    sum2+=v

	// }
	// fmt.Printf("Average is %.2f",float64(sum2/len(arr)))

	// fmt.Println("")

	// //assigning

	// var arr2 =[4]int{2,3,4,5}
	// fmt.Println(arr2)

	fmt.Println("")
}
