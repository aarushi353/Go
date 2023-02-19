package main
import (
	"fmt"
	"time"
	"booking-app/helper"
)

// package level variables
const conferenceTickets = 50
var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint
}

func main(){
   
   greetUsers()

   // for len(bookings) < 50 && remainingTickets > 0 {if u wanna add a condition in for loop}
   for {
    
	firstName, lastName, email, userTickets := getUserInput()	    
	isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)
	
	if isValidName && isValidEmail && isValidTicketNumber{
		
		bookTicket(userTickets, firstName, lastName, email)
		go sendTicket(userTickets, firstName, lastName, email)

	    firstNames := getFirstNames()
    	fmt.Printf("Bookings are: %v\n", firstNames)
		if remainingTickets == 0 {
			// end the program
			fmt.Println("Our conference is booked out, Come back nest year.")
			break
		}
	} else {
		if !isValidName{
			fmt.Printf("first name you entered is too short, try again!\n")			
		}
		if !isValidEmail{
			fmt.Printf("email you entered is invalid, try again!\n")			
		}
		if !isValidTicketNumber{
			fmt.Printf("number of tickets you entered is invalid, try again!\n")			
		}

	}
   }
}

func greetUsers() {
	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets to attend")
}

func getFirstNames() []string{
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}

    return firstNames
}


func getUserInput()(string, string, string, uint){
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)
 
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)
 
	fmt.Println("Enter your email adress: ")
	fmt.Scan(&email)
 
	fmt.Println("Enter number of user tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
     	remainingTickets = remainingTickets - userTickets
	 
		// create a map for a user

		var userData = UserData {
			firstName: firstName,
			lastName: lastName,
			email: email,
			numberOfTickets: userTickets,
		}
		
		bookings = append(bookings, userData)
	    fmt.Printf("List of bookings is %v\n", bookings)

	 	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	 	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

}

func sendTicket(userTickets uint, firstName string, lastName string, email string){
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
    fmt.Println("#################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
    fmt.Println("#################")
}