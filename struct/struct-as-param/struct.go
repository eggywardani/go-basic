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

	fmt.Println(eggy.Name)
	fmt.Println(eggy.Address)
	fmt.Println(eggy.Age)

	user := showCustomer(eggy)
	fmt.Println(user)

}

func showCustomer(customer Customer) string {
	result := fmt.Sprintf("Name: %s, Address: %s, Age: %d", customer.Name, customer.Address, customer.Age)
	return result
}
