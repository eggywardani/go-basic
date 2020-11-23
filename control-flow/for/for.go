package main

import "fmt"

func main() {

	// i := 1
	// for i <= 10 {
	// 	fmt.Println(i)
	// 	i++
	// }

	for i := 1; i < 2; i++ {
		fmt.Println("No", i)
	}

	slice := []string{"Eggy", "Andika", "Wardani"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])

	}

	for _, value := range slice {
		// fmt.Println("Index", i, "=", value)
		fmt.Println(value)

	}
}
