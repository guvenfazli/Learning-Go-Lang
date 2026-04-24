package main

import (
	"fmt"
)

func main() {

	mySalary := 1000
	isItDone := 0
	var expenses int

	fmt.Print("Your total expenses: ")
	fmt.Scan(&expenses)
	fmt.Println("Was it your last expense?")
	fmt.Println("1) Yes")
	fmt.Println("2) No")
	fmt.Scan(&isItDone)
	if isItDone == 2 {
		fmt.Println("Oh god...")
		return
	} else {
		fmt.Println("Your salary after promotion: ", mySalary)
		fmt.Print("Please enter an imaginary salary: ")
		fmt.Scan(&mySalary)
		fmt.Println("With that, your total revenue will be: ", mySalary-expenses)
	}

}
