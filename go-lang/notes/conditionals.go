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
	// go only uses the 'for' loop
	// for loop to continue the application after a booking is made with conditional statement
	for {
		var firstName string
		var lastName string
		var email string
		var userTickets int
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)
		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)
		fmt.Println("Enter your email: ")
		fmt.Scan(&email)
		fmt.Println("How many tickets would you like? ")
		fmt.Scan(&userTickets)

		// conditional for if the user tries to book more tickets than are available or less than what's avaialable
		if userTickets < int(remainingTickets) {
			remainingTickets = remainingTickets - uint(userTickets)
			bookings_slice = append(bookings_slice, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets.\nYou will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			firstNames := []string{}
			for _, booking := range bookings_slice {
				// this function will split the string with white space as a separator and returns a slice with the split elements
				var names = strings.Fields(booking)
				// adding an element to a slice
				firstNames = append(firstNames, names[0])

			}
			fmt.Printf("The first names of bookings ar: %v\n", firstNames)
			// creating if statement to check the number of tickets left
			if remainingTickets == 0 {
				// if condition is met then end the program
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			fmt.Printf("We only have %v tickets reamining, so you can't book %v tickets\n", remainingTickets, userTickets)
		}

	}
}
