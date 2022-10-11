package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "Go conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	var bookings []string

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend!")

	for {
		var firstName string
		var lastName string
		var email string
		var UserTickets uint
		// ask user for their name
		fmt.Println("Enter your first name")
		fmt.Scan(&firstName)
		fmt.Println("Enter your last name")
		fmt.Scan(&lastName)
		fmt.Println("Enter your email")
		fmt.Scan(&email)
		fmt.Println("Enter N tickets")
		fmt.Scan(&UserTickets)

		//isValidName := len(firstName) > 1 && len(lastName) > 1
		//isValidEmail := strings.Contains(email, "@")
		//isValidTicketNumber := UserTickets > 0

		if remainingTickets < UserTickets {
			fmt.Printf("We hawe only %v tickets remaining, so you can't book %v tickets.\n", remainingTickets, UserTickets)
			continue
		}
		remainingTickets = remainingTickets - UserTickets
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("The whole array: %v\n", bookings)
		fmt.Printf("The first value: %v\n", bookings[0])
		fmt.Printf("Array type: %T\n", bookings)
		fmt.Printf("Array length: %v\n", len(bookings))

		fmt.Printf("Thank you %v %v for booking %v tickets/ You will recive a confirmation email at %v.\n", firstName, lastName, UserTickets, email)

		fmt.Printf("User %v booked %v tickets.\n", firstName, UserTickets)
		fmt.Printf("%v tickets remaining for %v.\n", remainingTickets, conferenceName)

		var firstNames []string
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("The first names of bookings are: %v\n", firstNames)

		if remainingTickets == 0 {
			//end the program
			fmt.Printf("Our conference is booked up")
			break
		}
	}
}
