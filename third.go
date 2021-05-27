package main

import (
	"fmt"
)

func add(x, y int) int {
	return x + y
}

func isEven(g int) bool {
	var even bool
	if g%2 == 0 {
		even = true
	} else {
		even = false
	}
	return even
}

func prize() (int, string) {
	i := 100
	s := "goats"
	return i, s
}
func sumNumbers(numbers ...int) int {
	total := 0
	for _, number := range numbers { //come back when you're done with five
		total += number
	}
	return total
}

func greet() {
	cc := "Hello"
	yy := "Motherfucker"
	fmt.Println(cc, yy)

}

func recursive(portion, eaten int) int {
	haveEaten := portion + eaten
	if haveEaten >= 5 {
		fmt.Println("I have had enough to eat, i ate ", haveEaten, " portions")
		return haveEaten //I noticed that the return function acts like kind of a stop keyword and stops the function from running continuously
	}
	fmt.Println("Can i have some more please, I've only had ", haveEaten, "portions")
	return recursive(portion, haveEaten)
}

func anotherFunction(dd func()string) string {
	return dd()
}

func wtmd() {
	fdd := "GLOCK GLOCK 3000!"
	fmt.Println(fdd)

}

func printOut() {
	fmt.Println("You're clunch game weak")
}

func hotBoySummer(numbersGotten int) int {
	numbersGotten2:= numbersGotten + 2
	if numbersGotten >= 10 {
		fmt.Println("I've met enough babes for the night")
		return numbersGotten  // I NEED MORE KNOWLEDGE ON RETURN STATEMENT
	} else {
		fmt.Println("Meet more babes")
	}

	return hotBoySummer(numbersGotten2)
}

func eatenYet(x, y int) (int, int, string) {
	var haveEaten string
	if x >= 10 && y >= 15 {
		haveEaten = "Yes I'm full"
		fmt.Println(haveEaten, "I had", x, "apples and", y, "oranges")
	}else {
		haveEaten = "I need more"
	}
	return x, y, haveEaten
}
func main() {
	// fmt.Println(isEven(10))
	// fmt.Println(isEven(5))
	// quantity, item := prize()
	// fmt.Println("The prize is ", quantity , item)
	// fmt.Println(sumNumbers(1,2,3,4,5,90))
	//greet()
	//recursive(1,0)
	//fmt.Println(anotherFunction(wtmd))
	hotBoySummer(0)
	// wtmd()
	// eatenYet(10,15)

}
