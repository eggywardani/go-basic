package main

import "fmt"

func main() {

	person := map[string]string{
		"name":    "Eggy Andika Wardani",
		"address": "Kudus",
	}

	person["title"] = "Student"

	fmt.Println(person)

	fmt.Println(person["name"])
	fmt.Println(person["address"])

	book := make(map[string]string)
	book["title"] = "Arbain"
	book["year"] = "1988"

	fmt.Println(book)
	delete(book, "year")
	fmt.Println(book)

}
