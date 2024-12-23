package main

import (
	// explicitly importing custom package
	"booking-app/helper"
	"fmt"
)

// these are package level variables
//   - package level variables can only use this syntax
const conferenceTickets int = 50

var conferenceName string = "Go Conference"
var remainingTickets uint = 50

// since we're creating a slice of maps, we need to define an initial size of the slice (slices are dynamic so this will expand)
var bookings_slice = make([]UserData, 0)

// creating a struct
//   - similar to map key-value pair structure but instead allows us to save mixed value types
//   - need to use 'type' statement
//     |_> 'type' creates a new type, with the name you specify (in this case a new type named 'UserData')
//   - can be thought of as a lightweight class
type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

// entrypoint
func main() {

	// call function to greet users
	greetUsers()

	for {
		// call function for getting user input
		firstName, lastName, email, userTickets := getUserInput()

		// call function for user input validation
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

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
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
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
	// creating a struct for a user
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: uint(userTickets),
	}
	// add map to bookings slice
	bookings_slice = append(bookings_slice, userData)
	// to confirm bookings, print slice of bookings data
	fmt.Printf("List of bookings is %v\n", bookings_slice)

	fmt.Printf("Thank you %v %v for booking %v tickets.\nYou will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
