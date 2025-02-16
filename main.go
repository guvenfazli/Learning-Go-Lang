package main

import "fmt"

func main() {
	var userValue float64
	var userYear int
	var infilationRate float64

	outputText("Please enter a value: ")
	fmt.Scanf("%f", &userValue)
	outputText("Please provide years: ")
	fmt.Scan(&userYear)
	outputText("Please enter an infilation rate: ")
	fmt.Scan(&infilationRate)

	withInfilation, withoutInfilation := calculateFutureValue(userValue, infilationRate, userYear)

	/* 	var result = fmt.Sprintf("Your future value with inflation will be: %.2f", userValue*infilationRate)
	 */

	fmt.Println("Your value with Infilation: ", withInfilation)
	fmt.Print("Your value without Infilation: ", withoutInfilation)

}

func outputText(userText string) {
	fmt.Print(userText)
}

func calculateFutureValue(userValue, infilationRate float64, year int) (float64, float64) {
	futureValueWithoutInfilation := userValue * float64(year)
	futureValueWithInfilation := (userValue * float64(year)) / infilationRate

	return futureValueWithInfilation, futureValueWithoutInfilation
}
