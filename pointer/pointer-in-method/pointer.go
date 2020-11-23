package main

import (
	"fmt"
)

type Man struct {
	Name string
}

func (man *Man) married() {
	man.Name = "Mr. " + man.Name
}

func main() {

	eggy := Man{"Eggy"}
	eggy.married()
	fmt.Println(eggy.Name)

}
