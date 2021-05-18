package main

import (
	"fmt"
	"reflect"
	"strconv"
)

var b bool
var str string = "my name is "
var beetels [4]string



func sayHello(s string) string {
	return "Hello " + s
}

func addition(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println("Hello World")
	fmt.Println(sayHello("Ebuka"))
	fmt.Println(addition(10, 43))
	fmt.Println(b)
	b = true
	fmt.Println(b)
	str += "Ebuka"
	fmt.Println(str)
	beetels[0] = "John"
	beetels[1] = "Paul"
	beetels[2] = "Ringo"
	beetels[3] = "Ebuka"
	fmt.Println(beetels)
	fmt.Println(reflect.TypeOf(beetels))
	fmt.Println(reflect.TypeOf(str))
	fmt.Println(reflect.TypeOf(b))
	var conv string = strconv.FormatBool(b)
	fmt.Println(reflect.TypeOf(conv))
	fmt.Println(conv)
	Bconv, _:= strconv.ParseBool(conv)
	fmt.Println(reflect.TypeOf(Bconv))
	var number int
	number, _ = strconv.Atoi(str)
	fmt.Println(reflect.TypeOf(number))
	fmt.Println(number)
}
