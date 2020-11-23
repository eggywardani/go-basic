package main

import "fmt"

func main() {
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}
	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	months[5] = "EGGY"
	fmt.Println(slice1)

	slice1[0] = "ANDIKA"
	fmt.Println(months)

	slice2 := months[10:]
	fmt.Println(slice2)

	slice3 := append(slice2, "WARDANI")
	fmt.Println(slice3)
	slice3[1] = "Bukan Desember"

	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(months)

	newSlice := make([]string, 2, 5)
	copySlice := make([]string, len(newSlice), cap(newSlice))

	newSlice[0] = "EGGY"
	newSlice[1] = "WARDANI"
	fmt.Println(newSlice)

	copy(copySlice, newSlice)

	fmt.Println(copySlice)

	iniArray := [5]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(iniArray)
	fmt.Println(iniSlice)

}
