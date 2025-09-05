package main

/*
# Section 02 - Exercise 07

Write a program that calculates the sum of ODD numbers between 1 to 10001.

## TODO 1

Calculate the sum of ODD numbers between 1 to 10001

    * HINT: Use the modulo operator: https://www.khanacademy.org/computing/computer-science/cryptography/modarithmetic/a/what-is-modular-arithmetic

ODD numbers  -> 1, 3, 5, 7, 9, 11, .... 9999, 10001

The result should be 25010001.

### REQUIREMENTS

1. You *must* use 'continue' in at least one of your solutions.

## TODO 2

Calculate the 'integer' average of all numbers from 1 to 10001, excluding the numbers:
   10, 19, 21, 39, 309, 431, 643, 942, 1209, 7981, 8888, 9910
Result should be: 5003

*/

import (
	"fmt"
)
func main(){

    var sum =0 
	for i:=1;i<=1001;i++{
		if(i%2==0){
			continue;
		}
		sum+=i


	}
	fmt.Printf("Sum of ODD numbers is: %d",sum)

	fmt.Println("")


	//problem 2:

    var sum2 ,count int
	for i:=1;i<=1001;i++{

		switch i {
		case 10, 19, 21, 39, 309, 431, 643, 942, 1209, 7981, 8888, 9910:
			continue;
		default:
			sum2+=i
			count++

		}	
	}

	fmt.Printf("Avergae of 1 to %v, excluding 10, 19, 21, 39, 309, 431, 643, 942, 1209, 7981, 8888, 9910 is equal to %v\n",
		10001, sum/count)
	

  fmt.Println("")
}

	