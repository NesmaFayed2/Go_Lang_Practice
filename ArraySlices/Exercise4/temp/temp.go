package temp

import (
	"fmt"

	"github.com/striversity/glft/shared/input"
)

const dataSize = 10

type (
	dataSet [dataSize]int
)

var data dataSet

func init() {
	for i := 0; i < len(data); i++ {

		data[i] = input.GetRandInt(-20, 120)
	}
}

func sort(d dataSet) dataSet {
	for i := range d {
		for j := i + 1; j < len(d); j++ {
			if d[j] <= d[i] {

				d[i], d[j] = d[j], d[i]
			}
		}
	}
	return d

}

func max(d dataSet) int {
	return d[len(d)-1]
}

func min(d dataSet) int {
	return d[0]
}

func sum(d dataSet) (sum int) {
	for _, v := range d {
		sum += v
	}
	return
}

func avg(d dataSet) (avg int) {
	sum := sum(d)
	avg = sum / len(d)
	return
}

func Print() {
	fmt.Println("Data: ", data)
	sdata := sort(data)
	fmt.Println("Sorted: ", sdata)
	fmt.Println("Max: ", max(sdata))
	fmt.Println("Min: ", min(sdata))
	fmt.Println("avg : ", avg(sdata))
}

func main() {
	fmt.Println()

}

// package temp

// /*
// Section 03 - Exercise 02 : Arrays and Functions
// TODO 1 - Stats on temperature values
// 	1) declare an array of about a few hundred 'int'
// 	2) assign a random integer between -20 and 120 to data[i]
// 	  - TIP: you can use input.GetRandInt(low, max) int
// 	3) print sorted list of temperature data
// 	4) print that max temperature
// 	5) print the min temperature
// 	6) print the average temperature
// TODO 2 - Calculate discount and sales tax:
// 1. Initializes an array representing the total purchase of 20 shoppers.
//    - TIP: Use sub-package 'cart.getTotal()'
// 2. For each shopper, calculate a 5% discount and 8% sales tax
// 3. Print the result

// */

// import (
// 	"fmt"

// 	"github.com/striversity/glft/shared/input"
// )

// var data [100]int

// func init() {
// 	for i := range data {
// 		data[i] = input.GetRandInt(-20, 120)
// 	}

// }

// func main() {

// 	/*
// 		TODO 1 - Stats on temperature values
// 		1) declare an array of about a few hundred 'int'
// 		2) assign a random integer between -20 and 120 to data[i]
// 		  - TIP: you can use input.GetRandInt(low, max) int
// 		3) print sorted list of temperature data
// 		4) print that max temperature
// 		5) print the min temperature
// 		6) print the average temperature

// 	*/

// 	fmt.Printf("")
// }

// func Print() {
// 	fmt.Printf("Sorted temp : %v", sort(data))
// 	fmt.Println("")
// 	fmt.Printf("max of temp : %v", max(sort(data)))
// 	fmt.Println("")
// 	fmt.Printf("min of temp : %v", min(sort(data)))
// 	fmt.Println("")
// 	fmt.Printf("avrage of temp : %v", avg(sort(data)))
// }

// func sort(d [100]int) [100]int {

// 	for i := range d {
// 		for j := i + 1; j < len(d); j++ {
// 			if d[j] < d[i] {
// 				d[j], d[i] = d[i], d[j]
// 			}
// 		}
// 	}
// 	return d

// }

// func max(d [100]int) (max int) {
// 	max = d[0]
// 	for i := range d {
// 		if d[i] > max {
// 			max = d[i]
// 		}
// 	}
// 	return

// }

// func min(d [100]int) (min int) {
// 	min = d[0]
// 	for i := range d {
// 		if d[i] < min {
// 			min = d[i]
// 		}
// 	}
// 	return

// }

// func avg(d [100]int) (avg float64) {
// 	var sum int
// 	for _, v := range d {
// 		sum += v
// 	}
// 	return float64(sum / len(d))

// }
