package main

import "fmt"

func main() {

	var balance float64 = 1000

	fmt.Println("Welcome to GO Bank!")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit Money")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. Exit")

	var choice int
	fmt.Print("Please make a choice: ")
	fmt.Scan(&choice)

	balanceCheck := choice == 1
	depositMoney := choice == 2
	withdrawMoney := choice == 3

	if balanceCheck {
		fmt.Println("Your balance: ", balance)
	} else if depositMoney {
		fmt.Print("Your Deposit: ")
		var depositValue float64
		fmt.Scan(&depositValue)
		if depositValue <= 0 {
			fmt.Print("Value should be higher than 0!")
			return
		}
		balance += depositValue
		fmt.Println("Your New Balance: ", balance)
	} else if withdrawMoney {
		fmt.Print("Your Withdraw: ")
		var withdrawMoney float64
		fmt.Scan(&withdrawMoney)
		if withdrawMoney <= 0 {
			fmt.Print("Value should be higher than 0!")
			return
		} else if withdrawMoney > balance {
			fmt.Print("You can not withdraw the money that does not exist!")
			return

		}
		balance -= withdrawMoney
		fmt.Println("Your New Balance: ", balance)
	} else {
		fmt.Print("Good Bye!")
	}
}
