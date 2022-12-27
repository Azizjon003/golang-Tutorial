package main

import "fmt"
const  constVar string  = "Hello World"
func main() {
var username string = "Aliqulov"
fmt.Println("My username is ", username)
fmt.Printf("this variable is of type %T \n", username)

var isBool bool =  true;
fmt.Println(isBool)
fmt.Printf("this variable is of type %T \n", isBool)

var smalVal uint8 =  254;
fmt.Println(smalVal)
fmt.Printf("this variable is of type %T \n", smalVal)

// implicit type
var implicitVal =  254;
fmt.Println(implicitVal)
fmt.Printf("this variable is of type %T \n", implicitVal)


// constant
fmt.Println(constVar)
fmt.Printf("this variable is of type %T \n", constVar)

}
