package main

import "fmt"

func main() {
	productNames := [4]string{"A Book"}
	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
	fmt.Println(prices)

	productNames[2] = "A Carpet"
	fmt.Println(productNames)
	fmt.Println("Prices[2]: ", prices[2])

	// Create slice of array (slice is a reference of a part of an array)
	featuredPrices := prices[1:]
	// Overwrite array value
	featuredPrices[0] = 199.99
	// Create a slice of another slice
	highlightedPrices := featuredPrices[:1]

	fmt.Println("highlightedPrices: ", highlightedPrices)
	fmt.Println("prices: ", prices)
	// Length gives the number of the elements in the slice/array
	fmt.Printf("Length of highlightedPrices: %v\n", len(highlightedPrices))
	// Capacity indicates the capacity left in the original array where it's based off of, the no. of elements left in the array.
	fmt.Printf("Capacity of highlightedPrices: %v\n", cap(highlightedPrices))

	highlightedPrices = highlightedPrices[:3]
	fmt.Println("highlightedPrices: ", highlightedPrices)
	fmt.Printf("Length of highlightedPrices: %v\n", len(highlightedPrices))
	fmt.Printf("Capacity of highlightedPrices: %v\n", cap(highlightedPrices))

	dynamicPrices := []float64{10.99, 8.99}
	fmt.Println(dynamicPrices[0:1])
	dynamicPrices[1] = 9.99

	updatedPrices := append(dynamicPrices, 5.99)
	fmt.Println(updatedPrices)

	secondStorePrices := []float64{4.99, 5.99, 9.99}
	thirdStorePrices := []float64{12.99, 15.99, 19.99}

	mergedStorePrices := append(secondStorePrices, thirdStorePrices...)
	fmt.Println("Merged Store Prices: ", mergedStorePrices)
}
