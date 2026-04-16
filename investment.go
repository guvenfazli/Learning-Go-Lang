package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64
	var myName string
	expectedReturnRate := 5.5
	var years float64

	fmt.Println("Please enter your investment amount: ")
	fmt.Scan(&investmentAmount)
	fmt.Println("Please also enter years: ")
	fmt.Scan(&years)
	fmt.Println("Please also enter expected return rate: ")
	fmt.Scan(&expectedReturnRate)
	fmt.Print("Last, please enter your name: ")
	fmt.Scan(&myName)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println("Mr.", myName)
	fmt.Println("Future Value: ", futureValue)
	fmt.Print("Future Real Value: ", futureRealValue)

}
