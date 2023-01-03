package main

import "fmt"

func main() {
	fmt.Println("Welcome to Structure")
hitesh := User{"Ma'ruf",23,"maruf@gmail.com",true}
fmt.Println("This is hitesh",hitesh)

fmt.Printf("This is hitesh %+v \n",hitesh)
fmt.Printf("This is hitesh %#v \n",hitesh)



fmt.Printf(" Name : %v \n Age : %v \n Email : %v \n IsActive : %v \n",hitesh.Name,hitesh.Age,hitesh.Email,hitesh.IsActive)

}

type User struct {
	Name string
	Age int
	Email string
	IsActive bool
}