package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Welcome to time Study in Golang")
	
	timeNow := time.Now()
	fmt.Println("The time now is", timeNow)

	fmt.Println("The year is", timeNow.Format("2006/01/02"))

}