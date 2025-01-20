package main //initiliasation
import (
	"fmt"
	"gofolder/help"
	"sync"
	"time"
)

var conferenceName = "Go Conference"

const conferenceTickets int = 50 //total number of tickets
var remainingTickets uint = 50
var bookings = make([]UserData, 0) //50 is the array size string is the type of elements in array array declaration

//fmt-format package

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberofTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	userGreet()
	// fmt.Printf("conferenceTickets is %T, conferenceName is %T, remainingTickets is %T \n", conferenceTickets, conferenceName, remainingTickets)

	//loop for below function

	//checking for validation
	firstName, lastName, email, userTickets := getUserInput()
	isValidName, isValidEmail, isValidTicketNumber := help.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {
		bookTicket(userTickets, firstName, lastName, email)

		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)
		//call function for print
		firstNames := getprintFirstNames()
		fmt.Printf("The first names of bookings are: %v\n", firstNames)

		if remainingTickets == 0 {
			fmt.Println("Our conference is booked out. Come back next year.")
			//break
		}
	} else {
		if !isValidName {
			fmt.Println("First name or last name you entered is too short")
		}
		if !isValidEmail {
			fmt.Println("Email address you entered does not contain @ sign")
		}
		if !isValidTicketNumber {
			fmt.Println("Number of tickets you entered is invalid")

		}
		//!isValidName means not true
	}
	wg.Wait() // it blocks until waitgroup is complete
}

func userGreet() {

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here")
}

func getprintFirstNames() []string {
	firstNames := []string{} //for getting first name
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName) //adding names to firstNames variable
	}

	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name")
	fmt.Scan(&firstName) // & points to variable name firstName to store the input

	fmt.Println("Enter your last name")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email:")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}
func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	//created map for user
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberofTickets: userTickets,
	}
	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v \n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(50 * time.Second)                                                       // simulate time delay in sending email,s;eep func blocks current execution
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName) // Sprintf helps to put together a string
	fmt.Printf("##############")
	fmt.Printf("sending ticket : \n %v \n to email address %v\n", ticket, email)
	fmt.Printf("##############")
	wg.Done()
} //sending this ticket to users email
