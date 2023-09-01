//--Summary:
//  The existing program is used to restrict access to a resource
//  based on day of the week and user role. Currently, the program
//  allows any user to access the resource. Implement the functionality
//  needed to grant resource access using any combination of `if`, `else if`,
//  and `else`.
//
//--Requirements:

//* Use the accessGranted() and accessDenied() functions to display
//  informational messages

package main

import "fmt"

// Days of the week
const (
	Monday    = 0
	Tuesday   = 1
	Wednesday = 2
	Thursday  = 3
	Friday    = 4
	Saturday  = 5
	Sunday    = 6
)

// User roles
const (
	Admin      = 10
	Manager    = 20
	Contractor = 30
	Member     = 40
	Guest      = 50
)

func accessGranted() {
	fmt.Println("Granted")
}

func accessDenied() {
	fmt.Println("Denied")
}

func weekday(day int) bool {
	return day <= 4
}

func main() {
	// The day and role. Change these to check your work.
	today, role := Monday, Contractor

	if role == Admin || role == Manager {
		fmt.Println("role is Admin or Manager")
		accessGranted()
	} else if role == Contractor && !weekday(today) {
		fmt.Println("today is weekends and role is Contractor")
		accessGranted()
	} else if role == Member && weekday(today) {
		fmt.Println("today is not weekends and role is Member")
		accessGranted()
	} else if role == Guest && (today == Monday || today == Wednesday || today == Friday) {
		fmt.Println("today is in the weekdays and role is Guest")
		accessGranted()
	} else {
		accessDenied()
	}
}
