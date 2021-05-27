package main

import (
	"fmt"
)

var instock [5]string
var newStock = make([]string, 3)


func main() {

	instock[0] = "PRETIOX"
	instock[1] = "COMBIZELL"
	instock[2] = "OROTHICK"
	instock[3] = "OROTHICK"
	instock[4] = "AGITAN80"

	fmt.Println(instock[3])
	fmt.Println(instock)
	fmt.Println("\n")

	newStock[0] = "MEKO"
	newStock[1] = "MERGAL"
	newStock[2] = "AGITA185"

	fmt.Println(newStock[2])
	fmt.Println(newStock)
	fmt.Println("\n")

	newStock = append(newStock, "TEXANOL", "BLUE PASTE") //This adds new strings to the slice and adjusts it accordingly
	fmt.Println(newStock[4])
	fmt.Println(newStock)
	fmt.Println("\n")

	newStock = append(newStock[:3], newStock[3+1:]...)  //This removes accordingly as well
	fmt.Println(newStock)
	fmt.Println("\n")

	fmt.Println(newStock[3])

	var newStockArrivals = make([]string, 4)
	copy(newStockArrivals, newStock[1:])

	fmt.Println("\n")
	fmt.Println(newStockArrivals)
	fmt.Println(newStockArrivals[0])
	fmt.Println(newStockArrivals[1])
	fmt.Println(newStockArrivals[2])
	//fmt.Println(newStockArrivals[3])


}
