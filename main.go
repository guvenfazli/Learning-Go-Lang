package main

import "fmt"

func main() {
	var balance float64 = 1000

	fmt.Println("Welcome to GO Bank!")

	for i := 0; i < 200; i++ {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Please make a choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance: ", balance)
		case 2:
			fmt.Print("Your Deposit: ")
			var depositValue float64
			fmt.Scan(&depositValue)
			if depositValue <= 0 {
				fmt.Println("Value should be higher than 0!")
				continue
			}
			balance += depositValue
			fmt.Println("Your New Balance: ", balance)
		case 3:
			fmt.Print("Your Withdraw: ")
			var withdrawMoney float64
			fmt.Scan(&withdrawMoney)
			if withdrawMoney <= 0 {
				fmt.Println("Value should be higher than 0!")

				continue
			} else if withdrawMoney > balance {
				fmt.Println("You can not withdraw the money that does not exist!")

				continue
			}

			balance -= withdrawMoney
			fmt.Println("Your New Balance: ", balance)
		default:
			fmt.Print("Good Bye!")
			return
		}
	}
}
