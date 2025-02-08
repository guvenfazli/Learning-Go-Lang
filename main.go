package main

import "fmt"

func main() {
	var userValue float64
	var multiplier float64

	fmt.Print("Please enter your value: ")
	fmt.Scan(&userValue)
	fmt.Print("Please enter your multiplier: ")
	fmt.Scan(&multiplier)
	fmt.Print(userValue * multiplier)
}
