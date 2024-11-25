package anonymous

import "fmt"

func main() {
	numbers := []int{1, 2, 3}
	quintiple := createTransformer(5)

	transformed := transformNumbers(&numbers, func(number int) int {
		return number * 2
	})

	quintipled := transformNumbers(&numbers, quintiple)

	fmt.Println(transformed)
	fmt.Println("Quintiple: ", quintipled)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}
