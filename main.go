package main

import "fmt"

func main() {
	var userValue float64
	var infilationRate float64
	fmt.Println("Please enter your value: ")
	fmt.Scanf("%f", &userValue)
	fmt.Println("Please provide an infilation rate: ")
	fmt.Scan(&infilationRate)

	var result = fmt.Sprintf("Your future value with inflation will be: %.2f", userValue*infilationRate)

	fmt.Print(result)

}
