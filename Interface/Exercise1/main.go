package main

import "fmt"

type (
	// Reader interface {
	// 	Read(buf []byte) (n int, err error)
	// }
	// Writer interface {
	// 	write(buf []byte) (n int, err error)
	// }
	// //cirlce and rect have area so we ade another intefrce and imbed it into them
	// Shape interface {
	// 	Area() float64
	// }
	// Circle interface {
	// 	Raduis() float64
	// 	Shape
	// }
	// Rect interface {
	// 	Width() float64
	// 	length() float64
	// 	Shape
	// }
	Stringer interface {
		String() string
	}

	Currency float64
)

func main() {

	// type Empty interface {
	// }
	// var e Empty
	// fmt.Printf("e's value: %v, type : %T", e, e)

	// var dataSource Reader
	// var printer Writer

	// fmt.Printf("dataSource's value : %v,type: %T\n", dataSource, dataSource)
	// fmt.Printf("printer's value : %v,type: %T\n", printer, printer)

	//____________
	//interfcae stores value and type -> vlaue of it is a type that implement of the interfcae methods
	// var c Currency = 11.3
	// fmt.Printf("currecy is : %v", c.String())

	// //you make var of the interface and you can assign that type that implement all interface method to it

	// var mainStringer Stringer
	// mainStringer = c
	// fmt.Printf("mainStringer's value : %v,type: %T\n", mainStringer, mainStringer)

	//pointer to the interface stringer
	var c = new(Currency)
	*c = 11.04
	fmt.Printf("currecy is : %v \n", c.String())
	fmt.Println(c)

	//you make var of the interface and you can assign that type that implement all interface method to it
	//impllement method for pointer vs fo a value concept
	var mainStringer Stringer
	mainStringer = c
	fmt.Printf("mainStringer's value : %v,type: %T\n", mainStringer, mainStringer)

}

func (c Currency) String() string {
	return fmt.Sprintf("$%.2f", float64(c))
}
