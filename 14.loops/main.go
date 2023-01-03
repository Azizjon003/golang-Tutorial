package main

import "fmt"
func main()  {
	 
	fmt.Println("Welcome to loops statement")
	// day := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	// for i := 0; i < len(day); i++ {
	// 	println("This is day",day[i])
	// }
	
	// for _,val := range day{
	// 	println("This is day",val)
	// }

	// for i,val := range day{
	// 	fmt.Printf("This is day %v   and index is %v \n",val,i)
	// }

	//while loop

	// i:=0
	// for i < 5 {
	// 	println("This is while loop")
	// 	i++

	// }
	// while loop with break
	// j := 0
	// for j<10 {
	// if j == 5 {
	// 	fmt.Println("This is break statement")
	// 	break
	// }
	// j++
	// }

	// while loop with continue
	// i:=0
	// for i<10 {
	// 	if i == 5 {
	// 		fmt.Println("This is continue statement ",i)
	// 		i++
	// 		continue
	// 	}
	// 	fmt.Println("This is while loop",i)
	// 	i++
	// }


	// for loop with goto
i:=0
for i<10 {
	if i == 5 {
		// fmt.Println("This is goto statement ",i)
		i++
		goto javlon
	}
	fmt.Println("This is for loop",i)
	i++
}

	javlon :
	     fmt.Println("This is Javlon Statement statement")
}