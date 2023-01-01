package main

import (
	"fmt"
	"sort"
)
 func main()  {
	
fruitList  := []string{"apple", "banana", "orange"}
	fmt.Printf("This is fruitList %T \n",fruitList)

	fruitList = append(fruitList,"grape","mango","pineapple")
fruitList = fruitList[4:]
	fmt.Println("This is fruitList",fruitList)


	// Slices make function
	hightScores := make([]int,4)
	hightScores[0] = 100
	hightScores[1] = 200
	hightScores[2] = 300
	hightScores[3] = 400
	fmt.Println("This is hightScores",hightScores)
	
	
	hightScores = append(hightScores,900,800,50)
	
	// hightScores = hightScores[3:]
	fmt.Println("This is hightScores",hightScores)
	fmt.Printf("This is hightScores %T \n",hightScores)
	sort.Ints(hightScores)

	fmt.Println("\nThis is hightScores is sorted",hightScores)
	fmt.Println("This is hightScores length :(",len(hightScores),")")
	fmt.Println("This is hightScores Is Sorted :(",sort.IntsAreSorted(hightScores),")")

	// remove slice element
	var lang = []string{"go","java","python","c++","c#","ruby","php"}
	index := 3
	lang = append(lang[:index],lang[index+1:]...)
	fmt.Println("This is lang",lang)



	var langs = make([]string,3)
	langs[0] = "go"
	langs[1] = "java"
	langs[2] = "python"
	fmt.Println("This is langs",langs)
	indx := 1
	langs = append(langs[:indx],langs[indx+1:]...)

fmt.Println("This is langs",langs)
}