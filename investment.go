package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount float64 = 1000
	expectedReturnRate := 5.5
	var years float64 = 10

	if years > 41 {
		fmt.Println("YES!")
	} else if years > 43 {
		fmt.Println("YES!")
	} else {
		fmt.Println("Woohoo!")
	}

	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	fmt.Print(futureValue)

}
