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

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	// fmt.Println("Hello, world", add(rand.Intn(10), 20), math.Pi)
	a, b := swap("Hello", "World")
	x, y := split(10)
	fmt.Println(x, y)
	fmt.Println(a, b)
}
