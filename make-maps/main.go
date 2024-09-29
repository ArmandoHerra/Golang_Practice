package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	// Pre-allocate space and capacity for an array we want to create
	// Making memory management more efficient
	userNames := make([]string, 2, 5)

	userNames[0] = "Jorge"
	userNames[1] = "Fernando"

	userNames = append(userNames, "Armando")
	userNames = append(userNames, "Alexis")

	fmt.Println(userNames)

	// Make function also works for pre-allocating space for maps
	// Only 1 argument is valid for maps
	courseRatings := make(floatMap, 5)

	courseRatings["go"] = 4.87
	courseRatings["react"] = 4.8

	courseRatings.output()

	// Looping over arrays, slices or maps
	for index, value := range userNames {
		fmt.Println("Index: ", index)
		fmt.Println("Value: ", value)
	}
}
