package main

import "fmt"
func main()  {
	fmt.Println("Hello arrays")

	// Arrays
	var appleArr = [4]string{"apple", "banana", "orange"}
	fmt.Println("This is appleArr",appleArr)
	fmt.Println("This is appleArr length :(",len(appleArr))
	 fruitArr := [5]int{1, 2, 3, 4, 5}
	fmt.Println("THis is fruit array",fruitArr)

	var arr = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("This is arr",arr[0])
	fmt.Println("This is arr length :(",len(arr),")")
	// fmt.Println(Sum(&arr))s
}