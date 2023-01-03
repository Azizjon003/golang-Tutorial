package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Println("This is switch case statement")
	rand.Seed(time.Now().UnixNano())

	randNum := rand.Intn(10)+ 1

	fmt.Println("Random number is: ", randNum)


	switch {
	case randNum == 1:
		fmt.Println("One")
	case randNum == 2:
		fmt.Println("Two")
	case randNum == 3:
		fmt.Println("Three")
	case randNum == 4:
		fmt.Println("Four")
	case randNum == 5:
		fmt.Println("Five")
	case randNum == 6:
		fmt.Println("Six")
	case randNum == 7:
		fmt.Println("Seven")
	default: 
		fmt.Println("Invalid number")
	}
}