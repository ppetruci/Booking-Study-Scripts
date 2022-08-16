package main

import (
	helper "booking-app/support"
	"fmt"
	"strings"
)

const conferenceTickets int = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var Bookings = []string{}

func main() {

	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	for {

		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTickets := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTickets {

			bookTickets(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are: %v \n", firstNames)

			if remainingTickets == 0 {
				// end the program
				fmt.Println("Our conference is booked out. Come back next year!")
				break

			} else {
				if !isValidName {
					fmt.Println("First name or last name is wrong")
				}
				if !isValidEmail {
					fmt.Println("Email is not correct")
				}
				if !isValidTickets {
					fmt.Println("Number of tickets is not valid")
				}

			}
		}
	}

}

func greetUsers(conferenceName string, conferenceTickets int, remainingTickets uint) {
	fmt.Printf("Welcome to our %v!\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, Bookings := range Bookings {
		var names = strings.Fields(Bookings)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
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

	return firstName, lastName, email, userTickets
}

func bookTickets(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	Bookings = append(Bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booked %v tickets. You receive a confirmation email at %v! \n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets reamining for %v. \n", remainingTickets, conferenceName)
}
