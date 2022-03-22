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


	fmt.Printf("Welcome to our %v booking app!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")

	for {
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
		remainingTickets = remainingTickets - userTickets 	
		// bookings[0]= firstName + " " + lastName
		bookings = append(bookings, firstName + " " + lastName) //no longer need to keep track of indices

		fmt.Printf("Thank you %v %v for booking %v ticket(s). We've sent a confirmation email at %v.\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			var firstName = names[0]
			firstNames = append(firstNames, firstName)
		}
		fmt.Printf("The first names of bookings are: %v\n", firstNames)
		
		if remainingTickets == 0{
			fmt.Println("Our conference is sold out. Come back next year.")
			break
		}
	}

	

}