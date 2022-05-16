package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var Debug bool
	var LogLevel = "info"
	var startUpTime = time.Now()

	fmt.Println(Debug, LogLevel, startUpTime)

	var seed int64 = 1234456789 // need int64 or else it will crash
	rand.Seed(seed)
}
