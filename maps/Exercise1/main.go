package main

import "fmt"

func main() {

	//review declaring slice -< intial value is nill

	var slice []int
	fmt.Println(slice)
	if slice == nil {
		fmt.Println("slice inialized with nil")
	}
	fmt.Println("____________________")
	//maps

	// var myMap map[int]string
	// fmt.Println(myMap)
	// myMap[1] = "nesma"
	// fmt.Println(myMap)

	// if myMap == nil {
	// 	fmt.Println("Map inialized with nil")
	// }

	// var myMap = make(map[int]string)

	// fmt.Println(myMap)
	// myMap[1] = "nesma"
	// fmt.Println(myMap)
	// myMap[-1] = "ahmed"
	// fmt.Println(myMap)
	// fmt.Println(myMap[-1])
	// fmt.Println(myMap[2])

	var family = make(map[string][]string)

	family["nesma"] = make([]string, 4)
	family["essam"] = make([]string, 3)
	family["jhon"] = make([]string, 3)

	family["nesma"][0] = "ahmed"
	family["nesma"][1] = "basma"
	family["nesma"][2] = "mohamed"
	family["nesma"][3] = "halla"

	family["essam"][0] = "ahmed"
	family["essam"][1] = "nesma"
	family["essam"][2] = "halla"

	fmt.Printf("fmaily available is : %v", family)
	fmt.Println()
	fmt.Printf("fmaily[nesma] is : %v", family["nesma"])

	fmt.Println()
	//checking if key exist
	if _, ok := family["john"]; !ok {

		fmt.Println("john dont have family")
	}

	// iterate over map
	for k, v := range family {
		fmt.Printf("family[%v] is %v", k, v)
		fmt.Println()
	}

	//delete from map
	delete(family, "jhon")

	fmt.Printf("fmaily available is : %v", family)

}
