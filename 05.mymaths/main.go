package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	// "math/rand"
)
 func main(){
	 fmt.Println("Welcome to maths in Golang")
	// numbers
	 //  var myNumberOne int = 10
	//  var myNumberTwo float64 = 20.5
	//  fmt.Println("The sum of two numbers is", myNumberOne +int(myNumberTwo)) 

	 // Random numbers
	  // rand.Seed(time.Now().UnixNano())
		
		// fmt.Println("Random number is", rand.Intn(100))
	//  fmt.Println("Random number is", rand.NormFloat64())
	//  fmt.Println("Random number is", rand.Float32())

	// Crypto random numbers
	myRandNum,_ :=rand.Int(rand.Reader,big.NewInt(100))
	fmt.Println("Random number is", myRandNum)
 }