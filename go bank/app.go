package main

import (
	"example.com/bank/fileops"
	"fmt"
)

const balanceFileName = "balance.txt"

func main() {
	accountBalance, err := fileops.ReadFloatFromFile(balanceFileName, 36880.0)

	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		fmt.Println("------------")
	}

	fmt.Println("Welcome to the Go Bank!.")

	for i := 0; i < 10; i++ {
		communication()

		var choice int
		fmt.Print("Please enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Printf("Your account balance is: %0.2f\n", accountBalance)
		case 2:
			var depositAmount float64
			fmt.Print("Enter the amount to deposit: ")
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid deposit amount. Please enter a positive value.")
				continue
			}

			accountBalance += depositAmount
			fmt.Printf("Deposit successful! Your new balance is: %.2f\n", accountBalance)
			fileops.WriteFloatToFile(balanceFileName, accountBalance)
		case 3:
			var withdrawalAmount float64
			fmt.Print("Enter the amount to withdraw: ")
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount <= 0 {
				fmt.Println("Invalid withdrawal amount. Please enter a positive value.")
				continue
			}

			if withdrawalAmount > accountBalance {
				fmt.Println("Insufficient funds!")
				continue
			}

			accountBalance -= withdrawalAmount
			fmt.Printf("Withdrawal successful! Your new balance is: %.2f\n", accountBalance)
			fileops.WriteFloatToFile(balanceFileName, accountBalance)
		case 4:
			fmt.Println("Goodbye!")
			fmt.Printf("Your final account balance is: %.2f\n", accountBalance)
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
