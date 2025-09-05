package main

/*
Given the function 'getUserAge()', don't worry about how it is implemented, write a program that does the following:

* _NOTE_: I have provided the function 'getUserAge()', which gets the user's age.
Complete the program using the value in the variable 'userAge'.

TODO 1 - Use 'if' statement to print the following appropriate message
printing "You are too young to sign up." if 'userAge' is less than or equal to 16,
printing "Come back when you are 18.", if 'userAge' is equal to 17
printing "Welcome, you are old enough to sign up!", if 'userAge' is greater than or equal to 18
TODO 2 - Run the program a few times to see different results

*/

import (
	"fmt"
	"math/rand"
	"time"
)
func main(){

	var age= getUserAge()
	fmt.Printf("your age is %d \n",age)
	if age<=16{
		fmt.Printf("You are too young to sign up.")
	}else if age==17{
		fmt.Printf("Come back when you are 18.") 
	}else if age>=18{
		fmt.Printf("Welcome, you are old enough to sign up!")
	}


  fmt.Printf("")
}

const (
	minAge = 1
	maxAge = 30
)


// getUserAge returns a value representing a user's age between minAge to maxAge
func getUserAge() (age uint8) {
	rand.Seed(time.Now().Unix())
	// get a random int value
	v := rand.Uint32()
	age = uint8(v%maxAge) + minAge
	return}
	