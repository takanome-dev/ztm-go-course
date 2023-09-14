//--Summary:
//  Create a program to manage lending of library books.
//
//--Requirements:
//* The library must have books and members, and must include:
//  - Which books have been checked out
//  - What time the books were checked out
//  - What time the books were returned
//* Perform the following:
//  - Add at least 4 books and at least 3 members to the library
//  - Check out a book
//  - Check in a book
//  - Print out initial library information, and after each change
//* There must only ever be one copy of the library in memory at any time
//
//--Notes:
//* Use the `time` package from the standard library for check in/out times
//* Liberal use of type aliases, structs, and maps will help organize this project

package main

import (
	"fmt"
	"time"
)

type Name string
type Title string

type LendAudit struct {
	checkOut time.Time
	checkIn  time.Time
}

type Member struct {
	name  string
	books map[Title]LendAudit
}

type BookEntry struct {
	total  int // total books owned by the lib
	lended int // total books lended out
}

type Library struct {
	members map[Name]Member
	books   map[Title]BookEntry
}

func printMemberAudit(member *Member) {
	for title, audit := range member.books {
		var returnTime string
		if audit.checkIn.IsZero() {
			returnTime = "[not returned yet]"
		} else {
			returnTime = audit.checkIn.String()
		}

		fmt.Println(member.name, ":", title, ":", audit.checkOut.String(), "through", returnTime)
	}
}

func printAudits(library *Library) {
	for _, member := range library.members {
		printMemberAudit(&member)
	}
}

func printLib(library *Library) {
	fmt.Println()
	for title, entry := range library.books {
		fmt.Println("For the book:", title, ", the total is", entry.total, "and", entry.lended, "has been lended out.")
	}
	fmt.Println()
}

func checkoutBook(library *Library, title Title, member *Member) bool {
	book, found := library.books[title]
	if !found {
		fmt.Println("Book", title, "is not found in the lib")
		return false
	}

	if book.lended == book.total {
		fmt.Println("Book", title, "has been lended out")
		return false
	}

	book.lended += 1
	library.books[title] = book

	member.books[title] = LendAudit{checkOut: time.Now()}
	return true
}

func returnBook(library *Library, title Title, member *Member) bool {
	book, found := library.books[title]
	if !found {
		fmt.Println("Book", title, "is not found in the lib")
		return false
	}

	audit, found := member.books[title]
	if !found {
		fmt.Println("Member", member.name, "didn't check out the book", title)
		return false
	}

	book.lended -= 1
	library.books[title] = book

	audit.checkIn = time.Now()
	member.books[title] = audit
	return true
}

func main() {
	library := Library{
		books:   make(map[Title]BookEntry),
		members: make(map[Name]Member),
	}

	library.books["Webapps in Go"] = BookEntry{
		total:  4,
		lended: 0,
	}
	library.books["The little Go Book"] = BookEntry{
		total:  3,
		lended: 0,
	}
	library.books["Let's learn Go"] = BookEntry{
		total:  2,
		lended: 0,
	}
	library.books["Go bootcamp"] = BookEntry{
		total:  1,
		lended: 0,
	}

	library.members["John"] = Member{"John", make(map[Title]LendAudit)}
	library.members["Anna"] = Member{"Anna", make(map[Title]LendAudit)}
	library.members["Mosh"] = Member{"Mosh", make(map[Title]LendAudit)}

	fmt.Println("\nInitial:")
	printLib(&library)
	printAudits(&library)

	member := library.members["John"]
	bookName := "Go bootcamp"

	fmt.Println(member.name, "is checking out a book...")
	checkedOut := checkoutBook(&library, Title(bookName), &member)

	if checkedOut {
		printLib(&library)
		printAudits(&library)
	}

	fmt.Println(member.name, "is returning the book...")
	returned := returnBook(&library, Title(bookName), &member)
	if returned {
		printLib(&library)
		printAudits(&library)
	}
}
