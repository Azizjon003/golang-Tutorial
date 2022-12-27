package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	fmt.Println("Welcome to our pizza App")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your age: ")
	age, _ := reader.ReadString('\n')

  fmt.Println("This is Age ", age)

	ageFloat,err := strconv.ParseFloat(strings.TrimSpace(age), 64)


	if(err != nil){
		fmt.Println("Error", err)
	}else{
		fmt.Println("This is Age ", ageFloat+1)
	}
}