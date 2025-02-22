package main

import (
	"fmt"
)

/* type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
} */

type CreatedItem struct {
	title  string
	price  string
	isSold string
}

func main() {

	itemTitle := createItem()
	itemPrice := createItem()
	itemSold := createItem()

	addedItem := CreatedItem{
		itemTitle,
		itemPrice,
		itemSold,
	}

	renderItem(addedItem)

	/* 	userFirstName := getUserData("Please enter your first name: ")
	   	userLastName := getUserData("Please enter your last name: ")
	   	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	   	var appUser User

	   	appUser = User{
	   		userFirstName,
	   		userLastName,
	   		userBirthdate,
	   		time.Now(),
	   	}

	   	// ... do something awesome with that gathered data!

	   	fmt.Println(userFirstName, userLastName, userBirthdate) */
}

func createItem() string {
	var enteredValue string
	fmt.Println("Please put a name for your item: ")
	fmt.Scan(&enteredValue)
	return enteredValue
}

func renderItem(itm CreatedItem) {
	fmt.Println("Your items title: ", itm.title)
	fmt.Println("Your items price: ", itm.price)
	fmt.Println("Is your item sold: ", itm.isSold)
}

/* func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
} */
