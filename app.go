package main

import (
	"fmt"
)

func main() {
	firstValue := 1628.8946267774415
	secondValue := 1272.4898790181621

	//	fmt.Println("First Value:", firstValue)
	//	fmt.Println("Second Value:", secondValue)

	fmt.Printf("First value: %.82f\nSecond Value: %.92f", firstValue, secondValue)

}
