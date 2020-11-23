package main

import "fmt"

func getFullName() (firstName, lastName string, age int) {
	firstName = "Eggy Andika"
	lastName = "Wardani"
	age = 21
	return firstName, lastName, age
}
func main() {
	firstName, lastName, age := getFullName()
	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(age)

}
