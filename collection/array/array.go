package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Eggy"
	names[1] = "Andika"
	names[2] = "Wardani"

	fmt.Println(names[0], names[1], names[1])
	var values = [3]int{
		44,
		44,
		56,
	}
	fmt.Println(values)
	fmt.Println(len(values))
	fmt.Println(values[0], values[1], values[1])
}
