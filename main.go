package main

import (
	"fmt"
	"go-tutorial/helper"
	"sync"
	"time"
)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	// var bookings [50]string
	// var bookings []string // slice, no length define
	// bookings := make([]map[string]string, 0) // dictionary
	bookings := make([]UserData, 0) // struct

	// fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)
	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	// for remainingTickets > 0 && len(bookings) < 50 {

	firstName, lastName, email, userTickets := getUserInput()

	isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	// fmt.Println(remainingTickets)
	// fmt.Println(&remainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {

		remainingTickets -= userTickets
		// bookings[0] = firstName + " " + lastName

		// var userData = make(map[string]string)
		var userData = UserData{
			firstName:       firstName,
			lastName:        lastName,
			email:           email,
			numberOfTickets: userTickets,
		}
		// userData["firstName"] = firstName
		// userData["lastName"] = lastName
		// userData["email"] = email
		// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

		bookings = append(bookings, userData)

		// fmt.Printf("Bookings: %v\n", bookings)
		// fmt.Printf("The first value: %v\n", bookings[0])
		// fmt.Printf("Bookings type: %T\n", bookings)
		// fmt.Printf("Total Bookings: %v\n", len(bookings))

		fmt.Printf("User %v %v booked %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)

		firstNames := getFirstNames(bookings)
		fmt.Printf("The first names of bookings are: %v\n", firstNames)

		// noTicketsRemaining := remainingTickets == 0
		// if noTicketsRemaining {
		if remainingTickets == 0 {
			// end the program
			fmt.Println("Our conference is booked out. Come back next year.")
		}

	} else {
		// invalid input
		// fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
		// continue
		fmt.Println("Invalid input data")
	}
	// }
	wg.Wait()
}

func greetUsers(confName string, confTickets int, remainingTickets uint) {
	fmt.Printf("Welcome to %v booking application\n", confName)
	fmt.Printf("We have total of %v tickets and %v are still available", confTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames(bookings []UserData) []string {
	firstNames := []string{}
	for _, booking := range bookings { // range provides the index and value for each element
		// names := strings.Fields(booking)
		// firstNames = append(firstNames, names[0])
		// firstNames = append(firstNames, booking["firstName"])
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// Ask user for their first name
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	// Ask user for their last name
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	// Ask user for their email
	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	// Ask user for number of tickets
	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(50 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("##########################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("##########################")
	wg.Done()
}
