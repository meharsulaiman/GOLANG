package main

import (
	"fmt"
)

func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	// fmt.Println("Hello, world", add(rand.Intn(10), 20), math.Pi)
	a, b := swap("Hello", "World")

	fmt.Println(a, b)
}
