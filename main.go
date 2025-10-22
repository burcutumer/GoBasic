package main

import (
	"booking-go/helper"
	"fmt"
	"time"
)

const conferenceTickets int = 50

var conferenceName = "Go Conference"
var remaininTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	userName        string
	email           string
	numberOfTickets uint
}

func main() {

	greetUsers()

	for {

		userName, email, userTickets := getUserInput()

		isValidTicket, isValidEmail, isValidName := helper.ValidateUserInput(userName, email, userTickets, remaininTickets)

		if isValidTicket && isValidEmail && isValidName {

			bookTicket(userTickets, email, userName)
			go sendTicket(userTickets, userName, email)

			var firstNames = getFirstName()
			fmt.Printf("the first names of bookings are: %v\n", firstNames)

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

func greetUsers() {
	fmt.Printf("Welcome to our %v \n", conferenceName)
	fmt.Printf("we have total of %v tickets and %v are still available.\n", conferenceTickets, remaininTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstName() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.userName)
	}
	return firstNames
}

func getUserInput() (string, string, uint) {
	var userName string
	var email string
	var userTickets uint

	fmt.Println("please enter your first name: ")
	fmt.Scan(&userName)

	fmt.Println("enter your email: ")
	fmt.Scan(&email)

	fmt.Println("enter number of tickets: ")
	fmt.Scan(&userTickets)

	return userName, email, userTickets
}

func bookTicket(userTickets uint, email string, userName string) {
	remaininTickets = remaininTickets - userTickets

	var userData = UserData{
		userName:        userName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("list of bookings: %v\n", bookings)

	fmt.Printf("Thank you %v for booking %v tickets. you will recieve email at %v mail address\n", userName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remaininTickets, conferenceName)
}

func sendTicket(userTickets uint, userName string, email string) {
	time.Sleep(5 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v\n", userTickets, userName)
	fmt.Println("**********")
	fmt.Printf("Sending ticket:\n %v to email address %v\n", ticket, email)
	fmt.Println("**********")
}
