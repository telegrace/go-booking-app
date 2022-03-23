package main

import (
	"booking-app/helper"
	"fmt"
	"strconv"
)

// to avoid reptition we can define variables shares among functions (package level variables) and they can't use :=
const conferenceTickets = 50
var conferenceName = "Go Conference"
var remainingTickets uint = 50 
var bookings = make([]map[string]string, 0) //empty slice of maps needs to have a size

func main() {
	// no need to pass those variables, they can be accessed from the package level
	greetUsers()

	for remainingTickets > 0 && len(bookings) < 50{  
		firstName, lastName, email, userTickets := getUserInput()

		//validation 
		isValidName, isValidEmail, isValidTickets := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTickets {
			bookTickets(remainingTickets, firstName, lastName, email)
			firstNames := getFirstNames()
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
}
// no need to pass those values as input 
func greetUsers ()  {
	fmt.Printf("Welcome to our %v booking app!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")
}

func getFirstNames () []string {
	firstNames := []string{}
	//each booking is now a map and not string
	for _, booking := range bookings {
		// append(what_we_are_appending_to, what_we_are_appending)
		firstNames = append(firstNames, booking["firstName"])
	}
	return firstNames
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

func bookTickets(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets 	
	
	//create a map for the a user 
	var userMap = make(map[string]string)
	userMap["firstName"] = firstName
	userMap["lastName"] = lastName
	userMap["email"] = email
	userMap["numberOfTickets"] = 	strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userMap) 
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v ticket(s). We've sent a confirmation email at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}