package main

import (
	"errors"
	"fmt"
)

func pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh nol")
	}
	return nilai / pembagi, nil
}
func main() {

	// errorNih := errors.New("Ups Error")
	// fmt.Println(errorNih)

	hasil, err := pembagian(10, 0)
	if err == nil {
		fmt.Println("Hasil :", hasil)
	} else {
		fmt.Println("Error :", err.Error())
	}

}
