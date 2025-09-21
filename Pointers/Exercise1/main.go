package main

import "fmt"

type Foo func(float64) (int, bool, string)

func main() {
	// var pCount *int
	// fmt.Printf("pCount: %v pCount : %p \n", pCount, pCount)

	// //more examples of declaring pointers
	// var pBool *bool
	// var pName *string
	// var pInt64 *int64
	// var pMap *map[string]int
	// var pChan *chan int
	// fmt.Printf("bool : %v , string: %v , int: %v ,msp: %v, chan : %v ", pBool, pName, pInt64, pMap, pChan)

	// // DEL=CLARE POINTER TO USER DEIFIENED TYPE
	// type SSN string
	// var pSsn *SSN
	// fmt.Printf("pSsn: %v\n", pSsn)

	// type Person struct {
	// 	fname string
	// 	lname string
	// 	age   uint8
	// }
	// var pPerson *Person
	// fmt.Printf("pPerson: %v\n", pPerson)

	//functions and pointers
	// var pFunc func(int)
	// fmt.Printf("pFunc: %v \n", pFunc)
	// pFunc(7)
	// pFunc := goo
	// pFunc(7)

	//& address
	// var pCount *int
	// var count int
	// fmt.Printf("pCount: %v pCount : %p \n", pCount, pCount)
	// pCount = &count
	// fmt.Printf("pCount: %v pCount : %p , type: %T\n", pCount, pCount, pCount)

	// name := "nesma"
	// fmt.Printf("&name : %v -> name : %v", &name, name)

	//pointer to pointer example

	//new function

	pCount := new(int)
	fmt.Printf("pCount: %v pCount : %p , type: %T\n", pCount, pCount, pCount)

	pName := new(string)
	fmt.Printf("pName: %v pName : %p , type: %T\n", pName, pName, pName)

}

// func main() {
// 	v := 0
// 	fmt.Println(&v)
// 	v = inc(v)
// 	fmt.Println(v, &v)
// }

// func inc(n int) int {
// 	return n + 1
// }

// func goo(x int) {
// 	fmt.Println(x)
// }
