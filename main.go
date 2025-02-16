package main

import (
	"errors"
	"fmt"
	"os"
)

// Goals
// 1) Validate user input
//    => Show error message & exit if invalid input is provided
//    - No negative numbers
//    - Not 0
// 2) Store calculated results into file

func validateUserInputs(revenue, expenses, taxRate float64) error {

	var invalidValue error

	if revenue <= 0 || expenses <= 0 || taxRate <= 0 {
		invalidValue = errors.New("value can not be negative")
	}

	return invalidValue
}

func saveTheResults(ebt, profit, ratio float64) error {
	var fileError error

	generalResults := fmt.Sprint("Your calculations are: \n", ebt, profit, ratio)

	fileWritingError := os.WriteFile("calculatedResults.txt", []byte(generalResults), 0064)

	if fileWritingError != nil {
		fileError = errors.New("can not write the file")
	}

	return fileError
}

func main() {
	revenue := getUserInput("Revenue: ")
	expenses := getUserInput("Expenses: ")
	taxRate := getUserInput("Tax Rate: ")

	newError := validateUserInputs(revenue, expenses, taxRate)

	if newError != nil {
		fmt.Print("ERROR")
		panic(newError)
	}

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fileError := saveTheResults(ebt, profit, ratio)

	if fileError != nil {
		fmt.Print("ERROR")
		panic(fileError)
	}

	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.3f\n", ratio)
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func getUserInput(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	return userInput
}
