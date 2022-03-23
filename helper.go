package main

import "strings"

//with multiple returns you need ()
func validateUserInput (firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) > 1 && len(lastName) > 1
	isValidEmail := strings.Contains(email, "@") && strings.Contains(email, ".")
	isValidTickets := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTickets
}