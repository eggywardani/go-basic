package main

import "fmt"

func main() {
	runApplication(true)

}

func runApplication(error bool) {
	defer endApp()
	if error {
		panic("Aplikasi Error")
	}

	fmt.Println("Application Running")
}
func endApp() {
	fmt.Println("Application Ended")
}
