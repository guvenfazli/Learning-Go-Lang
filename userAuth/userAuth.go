package userAuth

import (
	"fmt"
	"os"
)

type UserAccount struct {
	Name      string
	Surname   string
	Age       int64
	Email     string
	Password  string
	IsMarried bool
}

type LoggedInAccount struct {
	Name  string
	Email string
	Age   int64
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

	registeredAccount, error := os.ReadFile("userList.txt")
	convertedData := string(registeredAccount)
	fmt.Print(convertedData[0])

	if error != nil {
		fmt.Print("ERROR")
		panic(error)
	}

	createdAccount = UserAccount{
		Name:      userName,
		Surname:   userSurname,
		Age:       userAge,
		Email:     userEmail,
		Password:  userPassword,
		IsMarried: userIsMarried,
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
