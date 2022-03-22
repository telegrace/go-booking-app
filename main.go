package main

import (
	"fmt"
	"strings"
)
func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50 
	var bookings []string //this is an abstraction of an array, slices 
	// var bookings = [50]string{} //this is an array, we can't mix.
	// var bookings = [50]string arrayType
	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	for remainingTickets > 0 && len(bookings) < 50{  
		firstName, lastName, email, userTickets := getUserInput()

		//validation 
		isValidName, isValidEmail, isValidTickets := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTickets {
			bookTickets(remainingTickets, userTickets, bookings, firstName, lastName, email, conferenceName)
			firstNames := getFirstNames(bookings)
			fmt.Printf("The first names of bookings are: %v\n", firstNames)
			
			if remainingTickets == 0{
				fmt.Println("Our conference is sold out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("Please enter a valid name. Your first and last name must be at least 2 characters.")
			}
			if !isValidEmail {
				fmt.Println("Please enter a valid email with an '@' sign.")
			}
			if !isValidTickets {
				fmt.Println("Number of tickets you entered is invalid")
			}
			fmt.Println("You filled out the input incorrectly. Please try again")
		}
	}

	// city := "London"

	// switch city {
	// 	case "New York":
	// 	case "Singapore":
	// 	case "London", "Berlin" :
	// 		// same logic for London and Berlin
	// 	case "Mexico City":
	// 	case "Hong Kong":
	// 	default:
	// 		fmt.Print("No valid city selected")
	// }
}

func greetUsers (conferenceName string, conferenceTickets int, remainingTickets uint)  {
	fmt.Printf("Welcome to our %v booking app!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")
}

func getFirstNames (bookings []string) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		var firstName = names[0]
		firstNames = append(firstNames, firstName)
	}
	return firstNames
}
//with multiple returns you need ()
func validateUserInput (firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) > 1 && len(lastName) > 1
	isValidEmail := strings.Contains(email, "@") && strings.Contains(email, ".")
	isValidTickets := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTickets

}

func getUserInput () (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Please enter your first name:")
	fmt.Scan(&firstName) //this is a pointer, a special variable

	fmt.Println("Please enter your last name:")
	fmt.Scan(&lastName) 

	fmt.Println("Please enter your email:")
	fmt.Scan(&email) 

	fmt.Println("Please number of tickets:")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}

func bookTickets(remainingTickets uint, userTickets uint, bookings []string, firstName string, lastName string, email string, conferenceName string) {
	remainingTickets = remainingTickets - userTickets 	
	// bookings[0]= firstName + " " + lastName
	bookings = append(bookings, firstName + " " + lastName) //no longer need to keep track of indices

	fmt.Printf("Thank you %v %v for booking %v ticket(s). We've sent a confirmation email at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}