package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	revenue, error := userInput("Total Revenue: ")
	if error != nil {
		fmt.Println(error)
		return
	}

	expenses, error := userInput("Total Expenses: ")
	if error != nil {
		fmt.Println(error)
		return
	}

	taxes, error := userInput("Taxes: ")
	if error != nil {
		fmt.Println(error)
		return
	}

	ebt, profit, ratio := calculateProfits(revenue, expenses, taxes)

	fmt.Printf("%.2f\n", ebt)
	fmt.Printf("%.2f\n", profit)
	fmt.Printf("%.2f\n", ratio)
}

func userInput(infoText string) (float64, error) {
	var inputVal float64
	fmt.Print(infoText)
	fmt.Scan(&inputVal)

	if inputVal <= 0 {
		return 0, errors.New("value must be positive")
	}

	return inputVal, nil
}

func calculateProfits(revenue, expenses, taxes float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxes/100)
	ratio := ebt / profit
	writeResultsToFile(ebt, profit, ratio)
	return ebt, profit, ratio
}

func writeResultsToFile(ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT: %.2f\nProfit: %.2f\nRatio: %.2f\n", ebt, profit, ratio)
	os.WriteFile("results.txt", []byte(results), 0644)
}
