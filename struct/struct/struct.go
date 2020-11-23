package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var eggy Customer
	eggy.Name = "Eggy Andika Wardani"
	eggy.Address = "Kudus"
	eggy.Age = 21

	andika := Customer{
		Name:    "Andika",
		Address: "Malang",
		Age:     34,
	}

	wardani := Customer{"Wardani", "Bali", 29}

	fmt.Println(eggy.Name)
	fmt.Println(eggy.Address)
	fmt.Println(eggy.Age)

	fmt.Println(andika)
	fmt.Println(wardani)

}
