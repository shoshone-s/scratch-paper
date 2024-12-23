// NOTE:
// - Project or module needs to be initialized when first created (see notes)
// - All code must belong to a package; first statement in Go file must be 'package...' (see notes)
// - Need to declare an entrypoint

// declaring package
package main

// importing fmt package to utilize 'Print' function
import "fmt"

// entrypoint
func main() {
	// Can explicitly set the variable types

	// creating a variable
	// var conferenceName string = "Go Conference"
	// alt syntax for creating a variable
	// 	- Only applies to variables; cannot use with constants ('const')
	//  - Cannot explicitly state types
	conferenceName := "Go Conference"
	// creating a constant variable (amount of available tickets for the conference)
	const conferenceTickets int = 50
	// variable for available tickets remaining
	// vary type 'uint' won't allow for a negative integer
	var remainingTickets uint = 50

	// printing the type of the variables
	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	// Print formatted data ('Printf')
	// - Takes a template string that contains text that needs to be formatted
	// - '%v' - a placeholder for the variable we want to use in the string
	// - '\n' - adding in the newline
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still avaialable.\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets here to attend")

	// creating a null variable to be assigned later
	// 	- Need to also assign a data type to it (string, integer, bool, array, maps, etc.)
	var userName string
	var userTickets int
	// assigning value to 'userName' variable
	userName = "Tom"
	userTickets = 2
	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)

}
