//Summary:
//  Print basic information to the terminal using various variable
//  creation techniques. The information may be printed using any
//  formatting you like.
//
//Requirements:

//
//Notes:
//* Use fmt.Println() to print out information
//* Basic math operations are:
//    Subtraction    -
// 	  Addition       +
// 	  Multiplication *
// 	  Division       /

package main

import "fmt"

func main() {
	//* Store your favorite color in a variable using the `var` keyword
	var favoriteColor = "blue"
	fmt.Println("My favorite color is", favoriteColor)

	//* Store your birth year and age (in years) in two variables using
	//  compound assignment
	birthYear, ageInYears := 2000, 23
	fmt.Println("I was born in", birthYear, "and I'm", ageInYears)

	//* Store your first & last initials in two variables using block assignment
	var (
		firstName = "T"
		lastName = "D"
	)
	fmt.Println("My initials are:", firstName, lastName)

	//* Declare (but don't assign!) a variable for your age (in days),
	//  then assign it on the next line by multiplying 365 with the age
	// 	variable created earlier
	var myAgeInDays int
	myAgeInDays = 365 * ageInYears
	fmt.Println("I'm", myAgeInDays, "years old")
}
