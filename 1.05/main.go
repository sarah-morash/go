package main

import (
	"fmt"
	"time"
)

func main() {
	// Shorthand notation :=
	// Initial value is required for shorthand
	Debug := false
	LogLevel := "info"
	startUpTime := time.Now()

	fmt.Println(Debug, LogLevel, startUpTime)
}
