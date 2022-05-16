package main

import "fmt"

func main() {
	a, b := 5, 10
	fmt.Println("a: ", a, "b: ", b)

	// call swap here - pass pointers
	swap(&a, &b)
	fmt.Println("a: ", a, "b: ", b)
	fmt.Println(a == 10, b == 5)
}

// swap values
func swap(a *int, b *int) {
	*a, *b = *b, *a
}
