package main

/*


 */

import (
	"fmt"

	"github.com/striversity/glft/shared/input"
)






const(
	dataSize=10
)


func main(){

	var array [dataSize] int



	for i:=range array{
		array[i]= input.GetRandInt(-10,10)
	}

	fmt.Printf("Data set:%v\n",array)
	fmt.Printf("Min Value is : %v", min(array))

	fmt.Println("")

	//sorting using bubble sort
	// fmt.Printf("Array  : %v",array)
	// fmt.Println("")
	// sort(array)

	fmt.Printf("Array after sorting : %v",sort(array))


}

func min(arr[10]int) (min int){

	min =arr[0]
	for _,v:=range arr{
        if(v<min){
			min=v
		}
	}
	return

}

func sort(d [10] int) [10]int{

	for i:=range d{
		for j:=i+1;j<len(d);j++{
			if(d[j]<d[i]){
				d[j],d[i]=d[i],d[j]
			}
		}
	}
	return d

  
}