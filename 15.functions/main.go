package main

import (
	"errors"
	"fmt"
	"reflect"
)

func main()  {
	 fmt.Println("Welcome to functions in GoLang")

	// greeter() void functions call
   x,_:= added(10, 20)
	 fmt.Println(x)
  sum := ProAdded(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)


	 fmt.Println(sum)

	
}

// void functions

func greeter()  {

	fmt.Println("Hello Functions")

}

func  added(a int, b int ) (int,error) {

	if reflect.TypeOf(a) != reflect.TypeOf(b) {
		return 0, errors.New("Type Mismatch")
	}

	
	return a + b,nil
}
func ProAdded(values ...int) int{

	total := 0
	for _, v := range values {
		total += v
	}

	fmt.Println(total)
	return total

}