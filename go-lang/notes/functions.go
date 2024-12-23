package main

import (
	"fmt"
	"strings"
)

// these are package level variables
//   - package level variables can only use this syntax
const conferenceTickets int = 50

var conferenceName string = "Go Conference"
var remainingTickets uint = 50
var bookings_slice []string

// entrypoint
func main() {

	// call function to greet users
	greetUsers()

	for {
		// call function for getting user input
		firstName, lastName, email, userTickets := getUserInput()

		// call function for user input validation
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		// execute the booking if all the validations are valid
		if isValidName && isValidEmail && isValidTicketNumber {
			// call function for booking tickets
			bookTicket(userTickets, firstName, lastName, email)

			// call function to print first names
			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings ar: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				if len(firstName) < 2 {
					fmt.Printf("'%v' is an invalid name last name length. Please try again\n", firstName)
				}
				if len(lastName) < 2 {
					fmt.Printf("'%v' is an invalid name last name length. Please try again\n", lastName)
				}
			}
			if !isValidEmail {
				fmt.Printf("'%v' does not contain '@' sign. Please try again\n", email)
			}
			if !isValidTicketNumber {
				if userTickets > int(remainingTickets) {
					fmt.Printf("We're sorry, but there are only %v tickets remaining for this conference.\n", remainingTickets)
				} else {
					fmt.Printf("'%v' is an invalid number of tickets. Please try again\n", userTickets)
				}
			}
		}

	}
}

// function for greeting users
func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still avaialable.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

// function for printing the first names of the users who booked tickets
//   - since the function is returning a slice of strings, then it needs to be stated in the function definition
func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings_slice {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])

	}
	return firstNames
}

// function for validating user input
func validateUserInput(firstName string, lastName string, email string, userTickets int) (bool, bool, bool) {
	// input validation; checking if first and last name variable is valid
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	// input validation; checking if the email string has the '@' symbol
	isValidEmail := strings.Contains(email, "@")
	// input validation; checking if user didn't input a number of tickets that are negative or zero
	isValidTicketNumber := userTickets > 0 && userTickets <= int(remainingTickets)

	// returining multiple variables from a function
	return isValidName, isValidEmail, isValidTicketNumber
}

// function for getting user input
func getUserInput() (string, string, string, int) {
	var firstName string
	var lastName string
	var email string
	var userTickets int
	// var city string

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("How many tickets would you like? ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

// function for booking tickets
func bookTicket(userTickets int, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - uint(userTickets)
	bookings_slice = append(bookings_slice, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets.\nYou will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
