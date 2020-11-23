package main

import "fmt"

func main() {

	sayHelloFilter("anjing", spamFilter)

}
func sayHelloFilter(name string, filter func(string) string) {
	nameFiltered := filter(name)
	fmt.Println("Halo", nameFiltered)
}

func spamFilter(name string) string {
	if name == "anjing" {
		return "..."
	}

	return name

}
