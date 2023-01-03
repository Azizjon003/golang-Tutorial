package main

import "fmt"

func main() {
 fmt.Println("Welcome to Maps")
 // Map is a collection of key value pairs
 var myMaps = make(map[string]int)
 myMaps["go"] = 100
 myMaps["java"] = 200
 myMaps["python"] = 300
 myMaps["c++"] = 400
 fmt.Println("This is myMaps",myMaps)
fmt.Println("This is myMaps length",len(myMaps))

fmt.Println("This is myMaps go value",myMaps["go"])


 delete(myMaps,"go")


fmt.Println("This is myMaps",myMaps)
// my maps

timeZones := map[string]string{
	"IST":"India Standard Time",
	"UTC":"Universal Time Coordinated",
	"GMT":"Greenwich Mean Time",
}
fmt.Println("This is timeZones",timeZones) // map[IST:India Standard Time UTC:Universal Time Coordinated GMT:Greenwich Mean Time]

_,is := timeZones["Azizjon"] // bunda qiymat bor yoki yo'qmi tekshirish is := boolean qiymat qaytaradi


fmt.Println("This is val",is) // false
}