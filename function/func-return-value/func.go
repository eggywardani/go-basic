package main

import "fmt"

func getFullName() (string, string, int) {
	return "Eggy", "Rinda", 21
}
func main() {
	firstName, _, age := getFullName()
	fmt.Println(firstName)
	fmt.Println(age)

}
