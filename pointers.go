package main

import "fmt"

func main() {
	age := 32 // regular variable

	agePointer := &age

	var name string

	userName := &name
	fmt.Println(*userName)

	fmt.Println("Age:", agePointer)
	adultYears := getAdultYears(*agePointer)
	fmt.Println(adultYears)

}

func getAdultYears(age int) int {
	return age - 18
}
