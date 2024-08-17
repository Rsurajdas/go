package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 6.0
	var investedAmt float64 = 1000
	expectedReturnRate := 15.3
	years := 10.0

	fmt.Print("Invested Amount: ")
	fmt.Scan(&investedAmt)
	fmt.Print("Years: ")
	fmt.Scan(&years)
	fmt.Print("Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	result, inflationAdjustedResult := calculateResult(investedAmt, expectedReturnRate, years, inflationRate)
	// result := investedAmt * math.Pow(1+expectedReturnRate/100, years)
	// inflationAdjustedResult := result / math.Pow(1+inflationRate/100, years)
	formattedResult := fmt.Sprintf("Future value:- %.2f\n", result)
	formattedAdjustedResult := fmt.Sprintf("Future value(adjusted for inflation):- %.2f\n", inflationAdjustedResult)
	// fmt.Printf(`Future value:- %.2f
	// Future value(adjusted for inflation):- %.2f`, result, inflationAdjustedResult)
	// fmt.Println("Future value:- ",result)
	// fmt.Println("Future value(adjusted for inflation):- ",inflationAdjustedResult)
	fmt.Print(formattedResult, formattedAdjustedResult)
}

func calculateResult(investedAmt, expectedReturnRate, years, inflationRate float64) (result float64, inflationAdjustedResult float64) {
	result = investedAmt * math.Pow(1+expectedReturnRate/100, years)
	inflationAdjustedResult = result / math.Pow(1+inflationRate/100, years)
	return result, inflationAdjustedResult
}
