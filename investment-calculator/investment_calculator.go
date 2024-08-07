package main

import (
	"fmt"
	"math"
)

const inflationRate float64 = 2.5

func main() {

	var expectedReturnRate float64
	var investmentAmount float64
	var years float64

	fmt.Println("---INVESTMENT CALCULATOR---")
	fmt.Println(" ")

	outputText("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	outputText("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	outputText("Investment Length (Years): ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)

	fmt.Printf(`Future Value: %.1f
Future Value After Inflation %.1f`, futureValue, futureRealValue)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount float64, expectedReturnRate float64, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv = fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
}
