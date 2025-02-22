package main

import "fmt"

type UserAccount struct {
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
}

func CreateAnAccount() (UserAccount, bool) {
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

func LogInAccount() (LoggedInAccount, bool) {
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

}
