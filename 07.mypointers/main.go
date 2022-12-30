package main

import "fmt"
func main() {
fmt.Println("Welcome to pointers in Go")
// var ptr *int8

// fmt.Println("Value of ptr is", ptr)
myNum := 13
 var ptr = &myNum // ptr ga 13 ni manzili beriladi
 fmt.Println("Memory  of ptr is", ptr) // results  0xc0000a4008
 fmt.Println("Value of ptr is", *ptr) // results 13
 myNum = myNum + 15
 fmt.Println("Memory  of ptr is", ptr) // results  0xc0000a4008
 fmt.Println("Value of ptr is", *ptr) // results 28
 *ptr = *ptr + 2
 fmt.Println("Memory  of ptr is", myNum) // results  30
 fmt.Println("Value of ptr is", *ptr) // results 30
}