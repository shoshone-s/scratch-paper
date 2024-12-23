package main

import (
	"fmt"
)

// entrypoint
func main() {

	var conferenceName string = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still avaialable.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	// array for all the bookings
	// 	- Need to define the size of the array (Arrays in Go have fixed sizes)
	// 	- Need to define the data type of the elements expected in the array
	// 		- Cannot mix data types in Go arrays
	// var bookings = [50]string{} // blank array
	// var bookings [50]string // <-- blank array creation alt syntax
	// var bookings = [50]string{"Nana", "Nicole", "Peter"} <-- Creation of an array of strings

	// slice in Go
	// 	- An abstraction of an array; similar structure to an array but has more dynamic size
	// 	- More flexible and powerful with variable-length or get a sub-array of its own
	// 	- Index-based and have a size, but is resized when needed
	var bookings_slice []string
	// alt syntact of creating slices
	// var bookings_slice = []string{}
	// bookings_slice := []string{}

	var firstName string
	var lastName string
	var email string
	var userTickets int
	// ask user for their info
	// 	- In order to get user input need a pointer '&'
	// 	- The pointer is a variable that points to the memory address of another variable where the value is stored
	// 	- Pointers are also referred to as 'special variables'
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email: ")
	fmt.Scan(&email)
	fmt.Println("How many tickets would you like? ")
	fmt.Scan(&userTickets)
	// need to convert the type of 'userTickets' into 'uint' type
	remainingTickets = remainingTickets - uint(userTickets)
	// populating the array
	// bookings[0] = firstName + " " + lastName
	// populating the slice
	bookings_slice = append(bookings_slice, firstName+" "+lastName)

	// print out a thank you to the user
	fmt.Printf("Thank you %v %v for booking %v tickets.\nYou will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

	fmt.Printf("These are all our bookings: %v\n", bookings_slice)
}
