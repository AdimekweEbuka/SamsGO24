package main

import (
	"fmt"
)

func ifstatement(v int) {

	if v == 1000 {
		fmt.Println("Yes it is!")
	} else if v == 2000 {
		fmt.Println("You know nothing Jon Snow")
	} else if v == 10000 {
		fmt.Println("Lets work!")
	} else {
		fmt.Println("The number is not in this function")
	}
}

func passwordAuthentication(b string) {
	inputPassword := "THOMAS1998#"
	if inputPassword == b {
		fmt.Println("WELCOME EBUKA!")
	} else {
		fmt.Println("WRONG PASSWORD")
	}
}

func taxesPayment(income int) {
	//THIS INDICATES HOW MUCH TAXES WOULD BE PAYED BASED ON MONTHLY INCOME
	if income >= 1000000 {
		taxes := 0.1 * float32(income)
		fmt.Println("You have to pay", taxes, "in taxes my boy")

	} else if income <= 999999 {
		taxes := 0.01 * float32(income)
		fmt.Println("You have to pay", taxes, "Naira in taxes my boy")
	}
}

func twoFactorAuthen(pin int, name, ownersName string) {
	originalPin := 1546
	originalname := "THOMAS1998"

	if pin == originalPin && name == originalname {
		fmt.Println("Welcome", ownersName)
	} else {
		fmt.Println("You're not", ownersName)
	}
}

func pretioxStock(request int) {
	amountLeftkg := 10000

	if request <= amountLeftkg {
		fmt.Println("Yes! We have", request, "kg of Pretiox")
	} else {
		fmt.Println("I'm so sorry, we do not have that amount in stock yet")
	}
}

func orothickStock(amountLeftkg int) {
	switch amountLeftkg {
	case 10000:
		fmt.Println("We have that in stock")
	case 1000000:
		fmt.Println("We have that sir")
	case 40000:
		fmt.Println("Yeah nIgga e dey nah")
	default:
		fmt.Println("Let me check my roster again")
	}
}

func forLoopEG() {
	number := 1800
	// Multiple defer statements can be written to be executed as soon as the code is done but they run from the bottom up i.e. the codes run from the last written defer statement to the first defer statement
	defer fmt.Println("I AM AN INDICATOR THAT THE CODE IS DONE")
	defer fmt.Println("I am the first indicator")
	for number > 800 { //if this statement is true the code runs until the statement is falsethen it exits SIMPLE
		number = number - 29
		fmt.Println("The number is'nt right, keep going")
		fmt.Println(number)
	}
	fmt.Println("Fianl number", number, "is correct, well done")
	//This can also be written as
	defer fmt.Println("NO I AM THE FIRST")
	var food int
	for food = 1000; food > 200; {
		food = food - 79
		fmt.Println("The food is not enogh")
		fmt.Println(food)
	}
	fmt.Println("Thank you, I've had", food, "plates, im full")
}
func hourFiveExercises() {
	simple:= 708
	if simple > 200 {
		fmt.Println("Simple is greater")
	}

	if simple > 5 && simple < 10 {
		fmt.Println("THis is too easy")
	} else {
		fmt.Println("I wan sleep abeg")
	}

	scoops:= 0
	for scoops < 40 {
		scoops = scoops + 9
		fmt.Println("I just had", scoops, "scoops, please add more")
		fmt.Println(scoops)
	}
	fmt.Printf("I have had %v amount of scoops", scoops)
}
func main() {
	// ifstatement(10000)
	// ifstatement(2000)
	// ifstatement(10000)
	// ifstatement(24554545)
	// passwordAuthentication("THOMAS1998")
	// taxesPayment(99000)
	// twoFactorAuthen(1546, "THOMAS1998", "JACOB")
	// pretioxStock(100000000000000)
	// pretioxStock(100)
	// orothickStock(400)
	forLoopEG()
	hourFiveExercises()
}
