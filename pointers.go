package main

import (
	"fmt"
	"os"

	"example.com/learningGO/userAuth"
)

func main() {

	var createdAccount userAuth.UserAccount
	var loggedInAccount userAuth.LoggedInAccount
	var accountCreated bool
	var loggedIn bool

	usersChoice := greetUserAndNavigate()

	if usersChoice == 1 {
		fmt.Println("Login!")
		var alreadyExistedAccount = os.ReadFile("userList.txt")
		loggedInAccount, loggedIn = userAuth.LogInAccount()
	} else if usersChoice == 2 {
		fmt.Println("Create an account!")
		createdAccount, accountCreated = userAuth.CreateAnAccount()
		convertedData := fmt.Sprintf("%v", &createdAccount)
		os.WriteFile("userList.txt", []byte(convertedData), 0644)
	} else {
		fmt.Print("Unknown command selected! App is shutting down.")
		return
	}

	if accountCreated {
		fmt.Printf("Welcome to our bank %v!", createdAccount.Name)
	}

	if loggedIn {
		fmt.Printf("Welcome back %v, i'll be bringing the menu soon! Thank you for using us.", loggedInAccount.Name)
	}

}

func greetUserAndNavigate() int64 {
	var accountOption int64
	fmt.Println("Welcome!")
	fmt.Println("In order to use my application, please log in or create an account!")
	fmt.Println("1. I do have an account: ")
	fmt.Println("2. I do not have an account: ")
	fmt.Print("Please choose an option: ")
	fmt.Scan(&accountOption)
	return accountOption
}
