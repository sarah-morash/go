package main

import "fmt"

var level = "pkg"

func main() {
	// fmt.Println("Main start :", level)
	// if true {
	// 	fmt.Println("Block start :", level)
	// 	funcA()
	// }

	// fmt.Println("Main start :", level)
	// // Create a shadow variable
	// level := 42
	// if true {
	// 	fmt.Println("Block start :", level)
	// 	funcA()
	// }
	// fmt.Println("Main end :", level)

	{
		level := "Nest 1"
		fmt.Println("Block end :", level)
	}
	// Error: undefined: level
	//fmt.Println("Main end  :", level)
}

func funcA() {
	fmt.Println("funcA start :", level)
}
