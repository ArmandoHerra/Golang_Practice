package main

import (
	"fmt"

	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main() {

	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		fmt.Println("----")
		// panic("Can't continue")
	}

	fmt.Println("Welcome to Golang Bank")
	fmt.Println("Reach us 24/7 at: ", randomdata.PhoneNumber())

	for {

		presentOptions()

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
					fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
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
					fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
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
