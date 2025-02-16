package main

import "fmt"

func main() {
	var userValue float64
	var infilationRate float64
	outputText("Please enter a value: ")
	fmt.Scanf("%f", &userValue)
	outputText("Please enter an infilation rate: ")
	fmt.Scan(&infilationRate)

	var result = fmt.Sprintf("Your future value with inflation will be: %.2f", userValue*infilationRate)

	fmt.Print(result)

}

func outputText(userText string) {
	fmt.Print(userText)
}
