package main

import "fmt"

func main() {

	var conferenceName = "Go conference"
	const conferenceTickets = 50
	var remainingTicket = 50

	fmt.Printf("Wlcome to our %v booking application\n", conferenceName)
	fmt.Println("We have total of", conferenceTickets, "Tickets and remaining tickets are", remainingTicket)
	fmt.Println("Get your tickets to attend")

	var userName string
	var userTicket int

	//Ask for user

	userName = "Tom"
	userTicket = 2
	fmt.Println(userName)
	fmt.Printf("user %v booked %v tickets.\n", userName, userTicket)

}
