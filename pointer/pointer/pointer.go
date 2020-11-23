package main

import (
	"fmt"
)

type Address struct {
	City, Province, Country string
}

func main() {

	//basic pointer
	number1 := 10
	number2 := &number1

	fmt.Println(number1)
	fmt.Println(number2)
	fmt.Println(*number2)

	fmt.Println("======================")

	// without pointer
	address1 := Address{"Kudus", "Jawa Tengah", "Indonesia"}
	address2 := address1
	address2.City = "Bandung"
	fmt.Println(address1)
	fmt.Println(address2)

	fmt.Println("======================")
	// with pointer
	address3 := Address{"Kudus", "Jawa Tengah", "Indonesia"}
	address4 := &address3
	address5 := &address3

	address4.City = "Semarang"
	address4 = &Address{"Palembang", "Sumatera Selatan", "Indonesia"}
	*address5 = Address{"Badung", "bali", "Indonesia"}

	// create pointer with new keyword
	var address6 = new(Address)
	address6.City = "Moscow"
	address6.Country = "Russia"

	fmt.Println(address3)
	fmt.Println(address4)
	fmt.Println(address5)
	fmt.Println(address6)

	fmt.Println("======================")
	number := 5
	fmt.Println("Nilai Awal", number)
	change(&number, 100)

	fmt.Println("Nilai Baru", number)

	fmt.Println("======================")
	var alamat = Address{
		"Ankara",
		"Hagia",
		"",
	}

	changeCountry(&alamat)
	fmt.Println(alamat)

}

func change(value *int, newValue int) {
	*value = newValue

}

func changeCountry(address *Address) {
	address.Country = "Turki"

}
