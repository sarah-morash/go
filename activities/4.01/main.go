package main

import "fmt"

func main() {
	var numbers [10]int

	for i := 0; i < len(numbers); i++ {
		numbers[i] = i + 1
	}

	fmt.Println(numbers)
}
