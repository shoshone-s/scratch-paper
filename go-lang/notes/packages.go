// this is a package that belongs to the main package
// 	- includes helper functions to the main package

// defining which package this belongs to
package helper

import "strings"

// - capitalizing a variable and defining in the package makes it exportable
// var Test string

// function for validating user input
// - capitalizing the function makes it exportable
func ValidateUserInput(firstName string, lastName string, email string, userTickets int, remainingTickets uint) (bool, bool, bool) {
	// input validation; checking if first and last name variable is valid
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	// input validation; checking if the email string has the '@' symbol
	isValidEmail := strings.Contains(email, "@")
	// input validation; checking if user didn't input a number of tickets that are negative or zero
	isValidTicketNumber := userTickets > 0 && userTickets <= int(remainingTickets)

	// returining multiple variables from a function
	return isValidName, isValidEmail, isValidTicketNumber
}

// from the file using the package
// import
// import (
// 	// explicitly importing custom package
// 	"booking-app/helper"
// 	"fmt"
// 	"strings"
// )

// // how to call function from package
// isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)
