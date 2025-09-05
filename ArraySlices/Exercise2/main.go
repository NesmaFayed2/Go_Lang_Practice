package main

/*
TODO 1
Declares an array of 10 to 20 elements of type float64 using a const for the array size
TODO 2
Initialize the array with random currency values using 'input.GetRandFloat()'.
TIP:
See Section 02 - Lecture 10 for the currency type
Use a 'for' loop to initialize your array values
The 'input' package is github.com/striversity/glft/shared/input and was covered in Sec01-Lec12.
TODO 3
Calculate total, max, min, avg

*/

import (
	"fmt"

	"github.com/striversity/glft/shared/input"
)


type(
	currency float64
)

func (c currency) String() string{
	s:= fmt.Sprintf("$%.2f",float64(c))
	return s
}




const(
	dataSize=10
)


func main(){

	var array [dataSize] currency

	//TODO 2 - Initialize the array with random currency values using 'input.GetRandFloat()'.
	for i:=range array{
		array[i]= currency(input.GetRandFloat(20,100))
	}

		//	TODO 3 - Calculate total, max, min, avg
     var totalPrice ,minPrice, maxPrice, avgPrice currency
	for _,v:=range array{
		totalPrice+=v
		if(v<minPrice ||minPrice==0){
            minPrice=v
		}
		if v>maxPrice{
			maxPrice=v
		}

	}
	avgPrice=totalPrice/dataSize

	fmt.Printf("Total price: %v\n", totalPrice)
	fmt.Printf("Min price: %v\n", minPrice)
	fmt.Printf("Max price: %v\n", maxPrice)
	fmt.Printf("Avg price: %v\n", avgPrice)




  fmt.Printf("")
}