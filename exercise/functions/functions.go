//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:

package main

import "fmt"

//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
func greet(name string) {
	fmt.Println("Greeting", name)
}

//* Write a function that returns any message, and call it from within
//  fmt.Println()
func printMessage() string {
	return "Golang is awesome!"
}

//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
func add(first, second, third int) int {
	return first + second + third
}

//* Write a function that returns any number
func returnANumber() int {
	return 3
}

//* Write a function that returns any two numbers
func returnAnyTwoNum() (int, int) {
	return 1, 7
}

func main() {
	greet("takanome-dev")

	fmt.Println(printMessage())

	sum := add(2, 5, 8)
	fmt.Println("The sum is", sum)

	first, second := returnAnyTwoNum()
	//* Add three numbers together using any combination of the existing functions.
	result := add(returnANumber(), first, second)
	
	//  * Print the result
	fmt.Println("The result is", result)
}
