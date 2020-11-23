package main

import "fmt"

func random() interface{} {
	return false
}
func main() {
	var result interface{} = random()
	// resultString := result.(string)
	// fmt.Println(resultString)

	switch value := result.(type) {
	case string:
		fmt.Println("string", value)
	case int:
		fmt.Println("int", value)
	default:
		fmt.Println("Unknown type")
	}

}
