package cart

/*
1. Initializes an array representing the total purchase of 20 shoppers.
   - TIP: Use sub-package 'cart.getTotal()'
2. For each shopper, calculate a 5% discount and 8% sales tax
3. Print the result

*/

import "fmt"


var totals = getTotal()
func main(){

  
  fmt.Printf("")
}

func Shop(){
	for i, total := range totals {
		// خصم 5%
		discounted := total * 0.95
		// ضريبة 8%
		final := discounted * 1.08

		// 3. نطبع النتيجة
		fmt.Printf("Shopper %d -> Original: %.2f, Final: %.2f\n", i+1, total, final)
	}
}


func getTotal() [20]float64 {
	return [20]float64{
		100, 200, 150, 90, 300, 50, 75, 400, 250, 180,
		120, 600, 80, 220, 330, 450, 30, 500, 275, 95,
	}
}
