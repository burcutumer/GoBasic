package main

import (
	"fmt"
	"strings"
)

func main() {

	var conferenceName = "Go Conference"
	bookings := []string{}
	const conferenceTickets uint = 50
	var remaininTickets uint = 50

	fmt.Printf("Wellcome to %v booking application\n", conferenceName)
	fmt.Println("Get your tickets here to attend")
	fmt.Printf("we have %v tickets\n", conferenceTickets)

	for {
		var userName string
		var email string
		var userTickets uint

		fmt.Println("please enter your first name: ")
		fmt.Scan(&userName)

		fmt.Println("enter your email: ")
		fmt.Scan(&email)

		fmt.Println("enter number of tickets: ")
		fmt.Scan(&userTickets)

		var isValidName bool = len(userName) >= 2
		isValidEmail := strings.ContainsAny(email, "@")
		isValidTicket := userTickets > 0 && userTickets <= remaininTickets

		if isValidTicket && isValidEmail && isValidName{
			remaininTickets = remaininTickets - userTickets
			bookings = append(bookings, userName)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}

			fmt.Printf("the first name slice: %v\n", firstNames)
			fmt.Println("we have ", remaininTickets, "tickets still available")

			if remaininTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year")
				break
			}
		} else {
			fmt.Println("your data is invalid")
			continue
		}

	}

}
