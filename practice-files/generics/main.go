package main

import "fmt"

func main() {
	result := add(1, 2)
	fmt.Println(result)
	result2 := add("Hello ", "World")
	fmt.Println(result2)
}

func add[T int | float64 | string](a, b T) T {
	return a + b
}
