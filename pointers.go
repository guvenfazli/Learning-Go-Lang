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

func main() {

	var createdAccount UserAccount
	var accountCreated bool
	usersChoice := greetUserAndNavigate()

	switch usersChoice {
	case 1:
		{
			fmt.Print("Login!")

		}
	case 2:
		{
			fmt.Print("Create an account!")
			createdAccount, accountCreated := createAnAccount()

		}
	default:
		{
			fmt.Print("Unknown command selected! App is shutting down.")
			return
		}
	}

	if accountCreated {

	}

}

func greetUserAndNavigate() int64 {
	var accountOption int64
	fmt.Println("Welcome!")
	fmt.Println("In order to use my application, please log in or create an account!")
	fmt.Println("1. I do not have an account: ")
	fmt.Println("2. I do have an account: ")
	fmt.Print("Please choose an option: ")
	fmt.Scan(&accountOption)
	return accountOption
}

func createAnAccount() (UserAccount, bool) {
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

func logInAccount(createdAccount UserAccount, createdAccountEmail, createdAccountPassword string) string {
	var userEmail string
	var userPassword string

	fmt.Println("Please enter your fresh created accounts email: ")
	fmt.Scan(&userEmail)
	fmt.Println("Password Please: ")
	fmt.Scan(&userPassword)

	if userEmail != createdAccountEmail || userPassword != createdAccountPassword {
		errorMessage := "Username or password is wrong, please restart the app and try again!"
		return errorMessage
	} else {

		greetMessage := "Welcome!"

		return greetMessage

	}

}
