package main

import (
	"fmt"
)

func main() {

	mySalary := 1000
	var expenses int
	const testValue int = 900

	fmt.Println("Your total expenses: ")
	fmt.Scan(&expenses)
	fmt.Println("Your current salary is: ", mySalary)
	fmt.Print("Please enter an imaginary salary: ")
	fmt.Scan(&mySalary)
	fmt.Println("With that, your total revenue will be: ", mySalary-expenses)

}
