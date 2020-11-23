package main

import "fmt"

func getHello(name string) string {
	return "Halo " + name
}

func add(number1 int, number2 int) int {
	return number1 + number2
}
func main() {

	name := getHello("Eggy Andika Wardani")
	fmt.Println(name)

	number := add(199, 22)
	fmt.Println(number)
}
