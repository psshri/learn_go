package helper

import "strings"

func IsValid(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	var isValidName bool = len(firstName) > 2 && len(lastName) > 2
	// isValidName := len(firstName) > 2 && len(lastName) > 2
	var isValidEmail bool = strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}
