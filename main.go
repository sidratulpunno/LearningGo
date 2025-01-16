package main

import (
	"fmt"
	"strings"
)

func main() {

	conferenceName := "Go conference"
	const conferenceTickets = 50
	var remainingTicket uint = 50

	fmt.Printf("Wlcome to our %v booking application\n", conferenceName)
	fmt.Println("We have total of", conferenceTickets, "Tickets and remaining tickets are", remainingTicket)
	fmt.Println("Get your tickets to attend")
	bookings := []string{}

	for {
		var firstName string
		var lastName string
		var email string

		var userTicket uint
		fmt.Println("Enter your first name")
		fmt.Scan(&firstName)
		fmt.Println("Enter your last name")

		fmt.Scan(&lastName)
		fmt.Println("Enter your email")

		fmt.Scan(&email)
		fmt.Println("Number of tickets you want to buy")

		fmt.Scan(&userTicket)

		//Ask for user
		remainingTicket = remainingTicket - userTicket
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thank you %v %v for buying %v tickets. You will recieve confiramtion email at %v\n", firstName, lastName, userTicket, email)
		fmt.Printf("Remaining tickets %v\n", remainingTicket)

		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("The first name of the bookings are: %v\n", firstName)
	}

}
