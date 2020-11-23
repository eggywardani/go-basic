package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

type Group struct {
	Name      string
	Premium   Customer
	Customers []Customer
}

func main() {
	var eggy Customer
	eggy.Name = "Eggy Andika Wardani"
	eggy.Address = "Kudus"
	eggy.Age = 21

	andika := Customer{"Andika", "Bandung", 24}
	users := []Customer{eggy, andika}

	group := Group{"Andika Wardani", eggy, users}
	fmt.Println(group)

}
