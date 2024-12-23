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

	// Using variable in string
	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Println("We have a total of", conferenceTickets, "tickets and", remainingTickets, "are still avaialable.")
	fmt.Println("Get your tickets here to attend")
}
