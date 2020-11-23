package main

import "fmt"

func main() {
	letter := getGoodBye
	fmt.Println(letter("Rina Martha"))

}

func getGoodBye(name string) string {
	return "GoodBye " + name

}
