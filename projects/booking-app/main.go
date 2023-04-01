package main

import (
	"fmt"
	"strings"
)

// defining variables at package level so that all the functions can access them

var conferenceName = "Go Conference"

const conferenceTickets = 50

var remainingTickets uint = 50
var bookings []string

func main() {
	// defining all the variables /////////////////////////////////////////////////////////////////
	// var conferenceName = "Go Conference"
	// the above line could also be written like below line
	// var conferenceName string = "Go Conference"
	// and the above line can also be written like below line
	// conferenceName := "Go Conference"
	// const conferenceTickets = 50
	// var remainingTickets uint = 50 // uint makes sure that the variable remainingTickets will never be assigned/assume a negative value
	// var bookings = [50]string{"pss", "nana"} // this is how you declare array and assign elements to it
	// you define an array in go using the line below
	// var bookings [50]string // this is how you just define the array of size 50 and data type string
	// var bookings []string
	// you define a slice in go using the line above
	// var bookings = []string{} // this is how you define and assign values to a slice
	// bookings := []string{}  // aise bhi define and assign kar sakte hain
	///////////////////////////////////////////////////////////////////////////////////////////////

	// writing intro message //////////////////////////////////////////////////////////////////////
	greetUsers()
	///////////////////////////////////////////////////////////////////////////////////////////////

	// for remainingTickets > 0 && len(bookings) < 50 {
	// example of infinte loop
	for {

		firstName, lastName, email, userTickets := userInput()

		// var firstName string
		// var lastName string
		// var email string
		// var userTickets uint // this has to be uint, otherwise we won't be able to substract this value from remaining ticket, data type has to be same in Go for calculation purpose

		// // fmt.Println(remainingTickets) // prints the value of remainingTickets
		// // fmt.Println(&remainingTickets) // prints the address of remainingTickets

		// fmt.Println("Enter your first name: ")
		// fmt.Scan(&firstName)
		// fmt.Println("Enter your last name: ")
		// fmt.Scan(&lastName)
		// fmt.Println("Enter your email address: ")
		// fmt.Scan(&email)
		// fmt.Println("Enter the number of tickets you wish to buy: ")
		// fmt.Scan(&userTickets)

		isValidName, isValidEmail, isValidTicketNumber := isValid(firstName, lastName, email, userTickets)

		// var isValidName bool = len(firstName) > 2 && len(lastName) > 2
		// // isValidName := len(firstName) > 2 && len(lastName) > 2
		// var isValidEmail bool = strings.Contains(email, "@")
		// isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketNumber {

			createBooking(userTickets, firstName, lastName, email)

			// remainingTickets = remainingTickets - userTickets

			// // bookings[0] = firstName + " " + lastName // this is how you add elements in an array
			// bookings = append(bookings, firstName+" "+lastName) // this is how you add elements in a slice in go

			// // fmt.Printf("The complete slice is : %v\n", bookings)
			// // fmt.Printf("The first element of slice is : %v\n", bookings[0])
			// // fmt.Printf("The size of slice is : %T\n", bookings)
			// // fmt.Printf("The length of slice is : %v\n", len(bookings))

			// fmt.Printf("Thank you %v %v for booking %v tickets, a confirmation email will be sent to %v\n", firstName, lastName, userTickets, email)
			// fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			firstNames := getFirstNames(bookings)
			fmt.Printf("Bookings are made by: %v\n", firstNames)

			// firstNames := []string{}
			// // this is how you iterate over an array in go
			// for _, booking := range bookings {
			// 	var names = strings.Fields(booking)
			// 	var firstName = names[0]
			// 	firstNames = append(firstNames, firstName)
			// }
			// fmt.Printf("Bookings are made by: %v\n", firstNames)

			noTicketsRemaining := remainingTickets == 0
			// if remainingTickets == 0 {
			if noTicketsRemaining {
				// end program
				fmt.Printf("Our %v is booked out, come back next year!\n", conferenceName)
				break
			}
		} else if userTickets == remainingTickets {
			// do something else
		} else {
			if !isValidName {
				fmt.Println("First name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("Email address you entered doesn't contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets you entered is invalid")
			}
			// fmt.Println("Your input data in invalid, try again")
			continue
		}
		// var noTicketsRemaining bool = remainingTickets == 0
	}

	// city := "London"

	// switch city {
	// case "New York":
	// 	// execute code for New York
	// case "Berlin", "Shanghai":
	// 	// execute code for Berlin, Shanghai
	// case "London":
	// 	// execute code for London
	// default:
	// 	fmt.Println("No valid city selected")
	// }
}

func userInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint // this has to be uint, otherwise we won't be able to substract this value from remaining ticket, data type has to be same in Go for calculation purpose

	// fmt.Println(remainingTickets) // prints the value of remainingTickets
	// fmt.Println(&remainingTickets) // prints the address of remainingTickets

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)
	fmt.Println("Enter the number of tickets you wish to buy: ")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}

func greetUsers() {
	// Printf doesn't add a new line explicitly, hence we need to add \n
	fmt.Printf("conferenceName is %T, conferenceTickets is %T and remainingTickets is %T\n", conferenceName, conferenceTickets, remainingTickets)
	fmt.Println("Welcome to the", conferenceName, "booking application.")
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("This application is written in Go!")
}

func getFirstNames(bookings []string) []string {
	firstNames := []string{}
	// this is how you iterate over an array in go
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		var firstName = names[0]
		firstNames = append(firstNames, firstName)
	}
	// fmt.Printf("Bookings are made by: %v\n", firstNames)
	return firstNames
}

func createBooking(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// bookings[0] = firstName + " " + lastName // this is how you add elements in an array
	bookings = append(bookings, firstName+" "+lastName) // this is how you add elements in a slice in go

	// fmt.Printf("The complete slice is : %v\n", bookings)
	// fmt.Printf("The first element of slice is : %v\n", bookings[0])
	// fmt.Printf("The size of slice is : %T\n", bookings)
	// fmt.Printf("The length of slice is : %v\n", len(bookings))

	fmt.Printf("Thank you %v %v for booking %v tickets, a confirmation email will be sent to %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
	// return remainingTickets, bookings
}

// 2:27:40
