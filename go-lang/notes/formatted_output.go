// importing fmt package to utilize 'Print' function
import "fmt"

// entrypoint
func main() {
	// creating a variable
	var conferenceName = "Go Conference"
	// creating a constant variable (amount of available tickets for the conference)
	const conferenceTickets = 50
	// variable for available tickets remaining
	var remainingTickets = 50

	// Print formatted data ('Printf')
	// - Takes a template string that contains text that needs to be formatted
	// - '%v' - a placeholder for the variable we want to use in the string
	// - '\n' - adding in the newline
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still avaialable.\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets here to attend")

}