package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName string = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	var Bookings []string

	fmt.Printf("Welcome to our %v Booking App\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint
		// ask user for their name
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email: ")
		fmt.Scan(&email)

		fmt.Println("Enter the number of tickets: ")
		fmt.Scan(&userTickets)

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTickets := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTickets {
			remainingTickets = remainingTickets - userTickets
			Bookings = append(Bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booked %v tickets. You receive a confirmation email at %v! \n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets reamining for %v. \n", remainingTickets, conferenceName)

			firstNames := []string{}
			for _, Bookings := range Bookings {
				var names = strings.Fields(Bookings)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("The first names of bookings are: %v \n", firstNames)

			if remainingTickets == 0 {
				// end the program
				fmt.Println("Our conference is booked out. Come back next year!")
				break

			} else {
				fmt.Printf("We only have %v tickets remaining, so you can`t book %v tickets \n", remainingTickets, userTickets)
			}

		}
	}

}
