package main

import "fmt"

func main() {
	fmt.Println(sum(22, 33, 11, 1, 2, 31, 34, 12, 299))

	numbers := []int{1, 2, 3, 55, 12, 12142, 34}
	total := sum(numbers...)
	fmt.Println(total)

}

func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}
