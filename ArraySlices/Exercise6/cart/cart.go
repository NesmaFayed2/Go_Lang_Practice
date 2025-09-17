package cart

/*
1. Initializes an array representing the total purchase of 20 shoppers.
   - TIP: Use sub-package 'cart.getTotal()'
2. For each shopper, calculate a 5% discount and 8% sales tax
3. Print the result

*/

import (
	"fmt"
	"math/rand"
)

var array [20]float64
var sliceOfArray = array[:]

func main() {

	fmt.Printf("")
}

func Shop() {
	//intialize slice of shooper amounts
	for i := range sliceOfArray {
		sliceOfArray[i] = getTotal()
	}
	//printing
	for i, purchase := range sliceOfArray {
		discount := purchase * 0.05
		afterDiscount := purchase - discount
		tax := afterDiscount * 0.08
		finalPrice := afterDiscount + tax

		fmt.Printf("Shopper %d:\n", i+1)
		fmt.Printf("  Original: $%.2f\n", purchase)
		fmt.Printf("  After 5%% discount: $%.2f\n", afterDiscount)
		fmt.Printf("  After 8%% tax: $%.2f\n", finalPrice)
		fmt.Println("---------------------------")
	}

}

func getTotal() float64 {
	// قيمة عشوائية بين 50 و 500
	return 50 + rand.Float64()*(500-50)
}
