package main

import (
	"fmt" // Formatted I/O
	"math/rand"
	"strings"
	"time"
)

/* Main Function
 * Seed the random number generator
 * Generate a random number between 1 - 5
 * Use string repeater to show the number of random stars
 */
func main() {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(5) + 1
	stars := strings.Repeat("*", r)
	fmt.Println(stars)
}
