package helper

import "strings"

func ValidateUserInput(userName string, email string, userTickets uint, remaininTickets uint) (bool, bool, bool) {
	var isValidName bool = len(userName) >= 2
	isValidEmail := strings.ContainsAny(email, "@")
	isValidTicket := userTickets > 0 && userTickets <= remaininTickets

	return isValidTicket, isValidEmail, isValidName
}