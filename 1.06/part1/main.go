package main

import (
	"fmt"
	"time"
)

func getConfig() (bool, string, time.Time) {
	return false, "info", time.Now()
}

func main() {
	// Declaring multiple varibles, shorthand notation
	// Debug, LogLevel, startUpTime := false, "info", time.Now()
	// fmt.Println(Debug, LogLevel, startUpTime)

	Debug, LogLevel, startUpTime := getConfig()
	fmt.Println(Debug, LogLevel, startUpTime)
}
