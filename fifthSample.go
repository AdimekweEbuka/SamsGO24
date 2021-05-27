package main

import (
	"fmt"
)

//SLICES
var warehouseStock = make([]string, 3)
var incomingStock = make([]string, 3)

//ARRAYS
var arrayTest [5]string

//MAPS
var mapsTest = make(map[string]int)
var mapsTest2 = make(map[int]string)
var HM = make(map[string]string)

func stockArrival(WS []string, IS []string) {
	WS = append(WS, IS...)
	fmt.Println(WS)
}


	
	


func main() {
	warehouseStock[0] = "PRETIOX"
	warehouseStock[1] = "COMBIZELL"
	warehouseStock[2] = "MERGAL"
	fmt.Println(warehouseStock)
	fmt.Println("\n")

	incomingStock[0] = "MEKO"
	incomingStock[1] = "CALCIUM"
	incomingStock[2] = "ZIRCO"
	fmt.Println(incomingStock)
	fmt.Println("\n")

	// copy(warehouseStock, incomingStock[1:1]) //How does the number work? the numbers in the [] after incominstock
	// fmt.Println(warehouseStock)

	stockArrival(warehouseStock, incomingStock)

	mapsTest["Pretiox"] = 90
	mapsTest["Combizell"] = 100
	mapsTest["OROTHICK"] = 200

	fmt.Println("\n")
	fmt.Println(mapsTest["Pretiox"])
	fmt.Println(mapsTest["OROTHICK"])
	fmt.Println(mapsTest)

	mapsTest2[1998] = "2 Million Dollars"
	mapsTest2[2304] = "7.5 Million DOllars"
	mapsTest2[4599897] = "150 Billion POunds"

	fmt.Println("\n")
	fmt.Println(mapsTest2[1998])
	fmt.Println(mapsTest2[4599897])
	fmt.Println(mapsTest2)
	fmt.Println("\n")

	delete(mapsTest2, 1998)
	fmt.Println(mapsTest2)
	fmt.Println("\n")


	HM["P"] = "Paragraph"
	HM["img"] = "Image"
	HM["h1"] = "Heading one"
	HM["h2"] = "Heading two"
	fmt.Println(HM)
	fmt.Println(HM["h2"])
}
