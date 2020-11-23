package main

import "fmt"

type HashName interface {
	GetName() string
}

func sayHello(hashName HashName) {
	fmt.Println("Hello", hashName.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func Ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	}
	return "Ups"
}
func main() {
	var eggy Person
	eggy.Name = "Eggy Andika Wardani"

	cat := Animal{"Pussy"}

	sayHello(eggy)

	sayHello(cat)

	// Interface Kosong
	var data interface{} = Ups(8)
	fmt.Println(data)
}
