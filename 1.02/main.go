package main

import "fmt"

var foo string = "HELLO" // Package level

func main() {
	var bar string = "WORLD" // Function level
	fmt.Println(foo, bar)
}
