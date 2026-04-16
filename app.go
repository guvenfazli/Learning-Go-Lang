package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to the profit calculator.")

	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Your Revenue:")
	fmt.Scan(&revenue)
	fmt.Print("Your Expenses:")
	fmt.Scan(&expenses)
	fmt.Print("Estimated Tax Rate:")
	fmt.Scan(&taxRate)

	fmt.Println("Your earnings before tax: ", revenue-expenses)
	fmt.Println("Your earnings after tax: ", (revenue-expenses)/taxRate)

}
