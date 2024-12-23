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
	// for loop to continue the application after a booking is made
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
		remainingTickets = remainingTickets - uint(userTickets)
		bookings_slice = append(bookings_slice, firstName+" "+lastName)

		fmt.Printf("Thank you %v %v for booking %v tickets.\nYou will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

		firstNames := []string{}
		// creating a 'for-each' loop
		// - 'range' allows us to iterate through different elements for different data structures (not just arrays and slices)
		// - but with arrays and slices, it will give us the index and element values
		// for index, booking := range bookings_slice { <-- original for-each loop but we need to use a blank identifier for the 'index' because it's not being used
		// - the blank identifier here is to ignore the 'index' variable
		for _, booking := range bookings_slice {
			// this function will split the string with white space as a separator and returns a slice with the split elements
			var names = strings.Fields(booking)
			// adding an element to a slice
			firstNames = append(firstNames, names[0])

		}
		fmt.Printf("The first names of bookings ar: %v\n", firstNames)
	}
}
