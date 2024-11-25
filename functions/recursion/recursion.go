package recursion

import "fmt"

func main() {
	fact := factorialRecursion(5)
	fmt.Println(fact)
}

// Recursion
func factorialRecursion(number int) int {

	if number == 0 {
		return 1
	}

	return number * factorialLoop(number-1)
}

// Loop
func factorialLoop(number int) int {
	result := 1

	for i := 1; i <= number; i++ {
		result = result * i
	}

	return result
}
