//--Summary:
//  Create a program that can create a report of rune information from
//  lines of text.
//
//--Requirements:
//* Create a single function to iterate over each line of text that is
//  provided in main().
//  - The function must return nothing and must execute a closure
//* Using closures, determine the following information about the text and
//  print a report to the terminal:
//  - Number of letters
//  - Number of digits
//  - Number of spaces
//  - Number of punctuation marks
//
//--Notes:
//* The `unicode` stdlib package provides functionality for rune classification

package main

import (
	"fmt"
	"unicode"
)

func numOfLetters(t string) int {
	total := 0

	for _, char := range t {
		if unicode.IsLetter(char) {
			total += 1
		}
	}

	return total
}

func numOfDigits(t string) int {
	total := 0

	for _, char := range t {
		if unicode.IsDigit(char) {
			total += 1
		}
	}

	return total
}

func numOfSpaces(t string) int {
	total := 0

	for _, char := range t {
		if unicode.IsSpace(char) {
			total += 1
		}
	}

	return total
}

func numOfPunct(t string) int {
	total := 0

	for _, char := range t {
		if unicode.IsPunct(char) {
			total += 1
		}
	}

	return total
}

func main() {
	lines := []string{
		"There are",
		"68 letters,",
		"five digits,",
		"12 spaces,",
		"and 4 punctuation marks in these lines of text!",
	}

	totalLetters, totalDigits := 0, 0
	totalSpaces, totalPunct := 0, 0

	for _, text := range lines {
		totalDigits += numOfDigits(text)
		totalLetters += numOfLetters(text)
		totalSpaces += numOfSpaces(text)
		totalPunct += numOfPunct(text)
	}

	fmt.Printf("The total of letters is %v\n", totalLetters)
	fmt.Printf("The total of digits is %v\n", totalDigits)
	fmt.Printf("The total of spaces is %v\n", totalSpaces)
	fmt.Printf("The total of punctuations is %v\n", totalPunct)
}
