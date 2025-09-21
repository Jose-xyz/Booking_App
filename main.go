package main

import "fmt"

func main() {

	var conferenceName = "Go Conference 2025"
	const conferenceTickets = 50 // when const is used the variable cannot be changed down the road
	var remainingTickets = 50

	fmt.Println("Welcome to", conferenceName, "booking app")
	fmt.Println("Total tickets:", conferenceTickets, "tickets and", remainingTickets, "left")
	fmt.Println("Get your tickets to attend")

}
