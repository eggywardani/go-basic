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

	eggy.sayHello("Nida")
	andika.sayHello("Lina")
	wardani.sayHello("Syifa")

	fmt.Println(eggy.showCustomer())
	fmt.Println(andika.showCustomer())
	fmt.Println(wardani.showCustomer())

}

func (customer Customer) sayHello(name string) {

	fmt.Println("Hello", name, "My name is", customer.Name)
}

func (customer Customer) showCustomer() string {
	result := fmt.Sprintf("Name: %s, Address: %s, Age: %d", customer.Name, customer.Address, customer.Age)

	return result

}
