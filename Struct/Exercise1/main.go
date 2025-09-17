package main

import "fmt"

func main() {

	type Person struct {
		fn          string
		ln          string
		age         uint8
		middleNames []string
		phones      map[string]string
	}

	var nesma Person

	nesma.fn = "nesma"
	nesma.ln = "fayed"
	nesma.age = 20
	//slice
	middlenames := make([]string, 2)
	middlenames[0] = "mohamed"
	middlenames[1] = "essam"
	nesma.middleNames = middlenames

	//map
	phones := make(map[string]string)
	phones["dial"] = "01003721863"
	nesma.phones = phones

	fmt.Println(nesma)

}
