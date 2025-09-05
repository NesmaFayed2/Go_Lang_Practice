package main

import "fmt"

func main(){
   

var arr [3]int

arr[0]=3
arr[1]=4
arr[2]=5

var sum  int
for i:=0;i<len(arr);i++{
 sum+=arr[i]

}

fmt.Printf("Average is %.2f",float64(sum/len(arr)))
fmt.Println("")

//range 

var sum2 int
for _,v:= range arr{
   sum2+=v
    
}
fmt.Printf("Average is %.2f",float64(sum2/len(arr)))

fmt.Println("")

//assigning

var arr2 =[4]int{2,3,4,5}
fmt.Println(arr2)

 fmt.Println("")
}