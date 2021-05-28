package main

import (
	"fmt"
)

type product struct {
	name string
	packSize int
}

type matchingBrands struct {
	clothes string
	shoes string
	accesories int
}

type superhero struct {
	name string
	age int
	power power
	address address
}

type power struct {
	flight bool
	superSpeed bool
	superStrenght bool
	fightSkills bool
}

type address struct {
	street int
	town string
	city string
}


func main() {
	product1:= product {
		name: "PRETIOX",
		packSize: 25,
	}
	product2:= product{
		name: "NATROSOL",
		packSize: 20,
	}

	fmt.Println(product2)
	fmt.Println(product1)
	fmt.Println(product1.name, product2.name)

	var shopping matchingBrands
	fmt.Println("\n")
	fmt.Println(shopping)
	shopping.accesories = 25
	shopping.clothes = "Versace"
	shopping.shoes = "Louis Vuttion"
	fmt.Println(shopping)
	fmt.Println("\n")

	supe1:= superhero {
		name: "Superman",
		age: 45,
		power: power {
			flight: true,
			superSpeed: true,
			superStrenght: true,
			fightSkills: true,
		},
		address: address {
			street: 10,
			town: "Gotham",
			city: "Wayne manor",
		},
	}
	fmt.Println(supe1)
	fmt.Println(supe1.power.superSpeed)
}