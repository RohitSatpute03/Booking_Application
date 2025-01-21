package main
import "strings"

func ValidInformation(firstName string, lastName string, emailId string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) > 2 && len(lastName) > 2
	isValidEmail := strings.Contains(emailId, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}