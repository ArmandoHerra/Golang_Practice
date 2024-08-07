package main

import (
	"errors"
	"fmt"
	"os"
)

// Goals
// 1) Validate user input
//	=> Show error message & exit if invalid input is provided
//  => No negative numbers
//  => No 0
// 2) Store calculated results into a file

func main() {
	revenue, err1 := getUserInput("Enter the Revenue: ")
	expenses, err2 := getUserInput("Enter the Expenses: ")
	taxRate, err3 := getUserInput("Enter the Tax Rate: ")

	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println(err1)
		return
	}

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf(`EBT: %.1f
Profit: %.1f
Ratio: %.1f`, ebt, profit, ratio)
	storeResults(ebt, profit, ratio)
}

func storeResults(ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\n Ratio: %.3f\n", ebt, profit, ratio)
	os.WriteFile("results.txt", []byte(results), 0644)
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("value must be a positive number")
	}

	return userInput, nil
}

func calculateFinancials(revenue float64, expenses float64, taxRate float64) (ebt float64, profit float64, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit
	return ebt, profit, ratio
}
