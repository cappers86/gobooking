package main

import (
	"fmt"
	"strings"
)
func main() {
	conferenceName := "Flexera Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50 
	bookings := []string{} //slice  with out the set ammount in the array 

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to our booking application for the %v.\n", conferenceName )
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend the", conferenceName, "in June 2022!")

	//loops 
// infinite loop 
	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)
		// &var give us the pointer value of where the var is stored so we can reference it. 

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName + " " + lastName)

		fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confimation email to %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

// _ is used to a variable you do not want to use like not using Index below.

		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		} 
		fmt.Printf("The first names of the bookings are: %v\n", firstNames)
	}

	
}
