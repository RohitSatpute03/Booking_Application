package main

import (
	"fmt"
	"sync"
	"time"
)

var conferenceName = "Go conference"

const conferenceTickets uint = 50

var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {
	greetUsers()

	for {
		// Get user input
		firstName, lastName, emailId, userTickets := getUserInput()

		// Validate inputs
		isValidName, isValidEmail, isValidTicketNumber := ValidInformation(firstName, lastName, emailId, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			// Book tickets
			BookTickets(userTickets, firstName, lastName, emailId)

			wg.Add(1)

			go sendTicket(userTickets, firstName, lastName, emailId)

			// Print first names of all bookings
			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			// Check if tickets are sold out
			if remainingTickets == 0 {
				fmt.Println("Tickets for the conference are sold out. Come back next year!")
				break
			}
		} else {
			// Handle invalid inputs
			if !isValidName {
				fmt.Println("First name or last name must be longer than 2 characters.")
			}
			if !isValidEmail {
				fmt.Println("The email address you entered is invalid. Please include an '@' symbol.")
			}
			if !isValidTicketNumber {
				fmt.Printf("We only have %v tickets remaining. You cannot book %v tickets.\n", remainingTickets, userTickets)
			}
			fmt.Println("Please try again.")
		}
	}
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets, and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend!")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var emailId string
	var userTickets uint

	// Ask user for details
	fmt.Print("Enter first name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email ID: ")
	fmt.Scan(&emailId)

	fmt.Print("Number of tickets you want: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, emailId, userTickets
}

func BookTickets(userTickets uint, firstName string, lastName string, emailId string) {
	remainingTickets -= userTickets

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           emailId,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is: %v\n", bookings)

	fmt.Printf("Thank you %v %v! You have booked %v tickets. You will receive your confirmation at %v.\n", firstName, lastName, userTickets, emailId)
	fmt.Printf("%v tickets remaining for %v.\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(5 * time.Second)
	var ticket = fmt.Sprintf("%v tickets booked by %v %v\n", userTickets, firstName, lastName)
	fmt.Println("########")
	fmt.Printf("Sending ticket %v to email address %v\n", ticket, email)
	fmt.Println("##########")
	wg.Done()
}
