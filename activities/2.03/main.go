package main

import "fmt"

func main() {
	// Define a slice with unsorted numbers in it.
	nums := []int{5, 8, 2, 4, 0, 1, 3, 7, 9, 6}

	// Print this slice to the console.
	fmt.Println("Before:", nums)

	// Sort the values using swapping.
	for swapped := true; swapped; {
		swapped = false
		for i := 1; i < len(nums); i++ {
			if nums[i-1] > nums[i] {
				nums[i], nums[i-1] = nums[i-1], nums[i]
				swapped = true
			}
		}
	}

	// Once done, print the now sorted numbers to the console.
	fmt.Println("After :", nums)
}
