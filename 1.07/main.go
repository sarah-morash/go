package main

import "fmt"

var defaultOffset = 10

func main() {
	offset := 5
	fmt.Println(offset)
	offset = 10
	fmt.Println(offset)

	offset = defaultOffset
	fmt.Println(offset)
	offset = offset + defaultOffset
	fmt.Println(offset)
}
