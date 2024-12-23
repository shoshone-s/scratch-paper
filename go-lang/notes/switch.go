package main

import (
	"fmt"
	"strings"
)

// entrypoint
func main() {

	var conferenceName string = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	var bookings_slice []string

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still avaialable.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {
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

		// input validation; checking if first and last name variable is valid
		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		// input validation; checking if the email string has the '@' symbol
		isValidEmail := strings.Contains(email, "@")
		// input validation; checking if user didn't input a number of tickets that are negative or zero
		isValidTicketNumber := userTickets > 0 && userTickets <= int(remainingTickets)

		// execute the booking if all the validations are valid
		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets = remainingTickets - uint(userTickets)
			bookings_slice = append(bookings_slice, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets.\nYou will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			firstNames := []string{}
			for _, booking := range bookings_slice {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])

			}
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
				fmt.Printf("'%v' is an invalid number of tickets. Please try again\n", userTickets)
			}
		}

	}
	// switch statement
	// 	- allows a variable to be tested for equality against a list of values
	city := "London"
	switch city {
	case "New York":
		// execute code for booking New York conference tickets
	case "Singapore", "Hong Kong":
		// execute code for booking Singapore & Hong Kong conference tikcets
	case "London", "Berlin":
		// execute code for booking London & Berlin conference tikcets
	case "Mexico City":
		// execute code for booking Mexico City conference tikcets
	default:
		fmt.Print("No valid city selected")
	}

}
