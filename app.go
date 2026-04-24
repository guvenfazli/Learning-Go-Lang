package main

import (
	"fmt"
)

func main() {
	firstValue := 1628.8946267774415
	secondValue := 1272.4898790181621

	formattedFV := fmt.Sprintf("First Value: %.1f\n", firstValue)
	formattedSV := fmt.Sprintf("Second Value: %.1f\n", secondValue)
	fmt.Printf("First Type: %v\nSecond Type: %v\n", formattedFV, formattedSV)

	//	fmt.Println("First Value:", firstValue)
	//	fmt.Println("Second Value:", secondValue)

	// fmt.Printf("First value: %.82f\nSecond Value: %.92f", firstValue, secondValue)

}
