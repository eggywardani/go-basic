package main

import "fmt"

func main() {
	runApplication()

}

func runApplication() {
	defer logging()
	fmt.Println("Program Utama")
}
func logging() {
	fmt.Println("Selesai memanggil function")
}
