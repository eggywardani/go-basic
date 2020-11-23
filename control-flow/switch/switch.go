package main

import "fmt"

func main() {

	name := "Eggy"

	switch name {
	case "Eggy":
		fmt.Println(name)
	case "Rina":
		fmt.Println(name)
	default:
		fmt.Println("Boleh Kenalan?")
	}

	// switch length := len(name); length > 5 {
	// case true:
	// 	fmt.Println("Nama terlalu panjang")
	// case false:
	// 	fmt.Println("Nama sudah benar")
	// }

	length := len(name)

	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")

	}

}
