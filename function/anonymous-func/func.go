package main

import "fmt"

//BlackList type function
type BlackList func(string) bool

func main() {
	user1 := func(name string) bool {
		return name == "admin"
	}
	registerUser("admin", user1)

	registerUser("user2", func(name string) bool {
		return name == "user"
	})

}

func registerUser(name string, bl BlackList) {
	if bl(name) {
		fmt.Println("you are blocked")

	} else {
		fmt.Println("you are my Member")

	}
}
