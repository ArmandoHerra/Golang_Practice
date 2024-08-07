package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)

	if err != nil {
		return 1000, errors.New("failed to find balance file!")
	}

	balanceText := string(data)
	balance, _ := strconv.ParseFloat(balanceText, 64)
	return balance, nil
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func main() {

	var accountBalance, err = getBalanceFromFile()

	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		fmt.Println("----")
		// panic("Can't continue")
	}

	fmt.Println("Welcome to Golang Bank")

	for {
		fmt.Println("What do you want to do?")
		fmt.Println("Option 1: Check Balance")
		fmt.Println("Option 2: Deposit Money")
		fmt.Println("Option 3: Withdraw Money")
		fmt.Println("Option 4: Exit")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		// wantsCheckBalance := choice == 1 // You can assign variables like this

		switch choice {
		case 1:
			fmt.Println("Your balance is: ", accountBalance)
		case 2:
			for {
				fmt.Print("Enter the amount to deposit: ")
				var depositAmount float64
				fmt.Scan(&depositAmount)
				if depositAmount > 0 {
					accountBalance += depositAmount
					fmt.Println("Balance Updated!")
					fmt.Println("Balance Available: ", accountBalance)
					writeBalanceToFile(accountBalance)
					break
				}
				fmt.Println("Invalid deposit amount. Deposit more than $0 USD.")
			}
		case 3:
			for {
				fmt.Print("Enter the amount to withdraw: ")
				var withdrawAmount float64
				fmt.Scan(&withdrawAmount)
				if withdrawAmount <= accountBalance {
					accountBalance -= withdrawAmount
					fmt.Println("Balance Updated!")
					fmt.Println("Balance Available: ", accountBalance)
					writeBalanceToFile(accountBalance)
					break
				}
				fmt.Println("Attempted balance to withdraw is larger than your current account balance!")
			}
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing our bank!")
			return
		}
	}
}
