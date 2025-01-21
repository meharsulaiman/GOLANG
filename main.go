package main

import (
	"fmt"
	"math"
	"math/rand"
)

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println("Hello, world", add(rand.Intn(10), 20), math.Pi)
}
