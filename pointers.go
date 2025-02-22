package main

import (
	"fmt"

	auth "example.com/learningGO/userAuth"
)

/* type UserAccount struct {
	name      string
	surname   string
	age       int64
	email     string
	password  string
	isMarried bool
}

type LoggedInAccount struct {
	name  string
	email string
	age   int64
} */

func main() {

	var createdAccount auth.UserAccount
	var loggedInAccount auth.LoggedInAccount
	var accountCreated bool
	var loggedIn bool

	usersChoice := greetUserAndNavigate()

	if usersChoice == 1 {
		fmt.Println("Login!")
		loggedInAccount, loggedIn = auth.LogInAccount()
	} else if usersChoice == 2 {
		fmt.Println("Create an account!")
		createdAccount, accountCreated = auth.CreateAnAccount()
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

/* func createAnAccount() (UserAccount, bool) {
	var userName string
	var userSurname string
	var userAge int64
	var userEmail string
	var userPassword string
	var userIsMarried bool

	var createdAccount UserAccount

	fmt.Println("Your name: ")
	fmt.Scan(&userName)
	fmt.Println("Your surname: ")
	fmt.Scan(&userSurname)
	fmt.Println("Your age: ")
	fmt.Scan(&userAge)
	fmt.Println("Your email: ")
	fmt.Scan(&userEmail)
	fmt.Println("Your password: ")
	fmt.Scan(&userPassword)
	fmt.Println("Are you married? ")
	fmt.Scan(&userIsMarried)

	createdAccount = UserAccount{
		userName,
		userSurname,
		userAge,
		userEmail,
		userPassword,
		userIsMarried,
	}
	accountCreated := true
	return createdAccount, accountCreated
}

func logInAccount() (LoggedInAccount, bool) {
	var userName string
	var userEmail string
	var userAge int64

	fmt.Println("Your name: ")
	fmt.Scan(&userName)
	fmt.Println("Your email: ")
	fmt.Scan(&userEmail)
	fmt.Println("Your age: ")
	fmt.Scan(&userAge)

	loggedInAccount := LoggedInAccount{
		userName,
		userEmail,
		userAge,
	}

	loggedIn := true

	return loggedInAccount, loggedIn

} */
