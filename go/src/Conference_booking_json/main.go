package main

import (
	"fmt"
	"strings"
)

var NumberOfTickets uint = 50

type Client struct{
	Name string
	Surname string
	Email string
	PhoneNumber string
	NumberOfTicketsbooking uint
}




func main(){

	var answer string
	listOfClients := getClient()
	
	fmt.Println("\nThe number of Tickets available:",NumberOfTickets)

	fmt.Println("\nDo you want to see client's information (y/n)? ")
	fmt.Scanln(&answer)

	if string(answer[0]) == "y"{
		fmt.Println(listOfClients)
	}
}




func welcomeUser() bool {
	fmt.Println("\nWelcome, here you can buy tickets for conference\n\nConference name: \"One step closer to Future\"\n\nSpeaker: Steve Jobs\n")

	
	var res string
	var response bool

	fmt.Printf("There are %v tickets available\n",NumberOfTickets)
	fmt.Println("\nDo you want to book any ticket(s) (y/n)? ")
	 
    fmt.Scanln(&res)
	if string(res[0]) == "y"{
		response = true
	}
	return response
}

func getClient()interface{}{
	clientAnswer := welcomeUser()

	clients := make([]Client,0)
	var client Client

	if clientAnswer{
		
        fmt.Print("\nEnter your firstname: \n")
		fmt.Scanln(&client.Name)

		fmt.Print("\nEnter your secondName: \n")
		fmt.Scanln(&client.Surname)

		fmt.Print("\nEnter your email address: \n")
		fmt.Scanln(&client.Email)

		fmt.Print("\nEnter your phone number: \n",)
		fmt.Scanln(&client.PhoneNumber)

		fmt.Print("\nEnter the number of tickets you want to buy:\n")
		fmt.Scanln(&client.NumberOfTicketsbooking)


		isValidFirstName,isValidEmail,isValidPhoneNumber, numberOfTicketsBooking := Validate(client.Name,client.Surname,client.Email,client.PhoneNumber,client.NumberOfTicketsbooking)

		if isValidFirstName && isValidEmail && isValidPhoneNumber && numberOfTicketsBooking <= NumberOfTickets{
			fmt.Println("\nWe will sent you ID number to your email soon. With it, you can take part in Conference.")
			clients = append(clients, client)
		}else if !isValidFirstName{
			fmt.Println("\nYou made a mistake when entering your firstname or lastname!\nYou try again via refreshing")
			return ":)"
		}else if !isValidEmail{
			fmt.Println("\nYou made a mistake when entering your Email!\nYou try again via refreshing")
			return ":)"
		}else if !isValidPhoneNumber{
			fmt.Println("\nYou made a mistake when entering your phone number!\nYou try again via refreshing")
			return ":)"
		}else if numberOfTicketsBooking > NumberOfTickets{ 
			fmt.Println("\nYou entered more than available tickets\nYou try again via refreshing")
			return ":)"
		}
		fmt.Println("\nThanks for booking tickets, See you soon on conference\n")
		NumberOfTickets = NumberOfTickets - client.NumberOfTicketsbooking


		

    } else{
		fmt.Println("\nOk, you are welcome anytime to book, but the number of tickets is limited so if you want to book, do it as soon as possible!")
		return ":)"
	}
	return clients
}



func Validate(firstName, secondName,email,phoneNumber string, numberOfTicketsBooking uint)(bool,bool,bool,uint){
	isValidName := len(firstName) >= 3 && len(secondName) >= 3
	isValidEmail := strings.Contains(email, "@")
	isValidPhoneNumber := phoneNumberChecking(phoneNumber)
	
	return isValidName,isValidEmail,isValidPhoneNumber,numberOfTicketsBooking
}



func phoneNumberChecking(phoneNumber string)bool{
	var istrue bool
	if len(phoneNumber) == 13 && phoneNumber[:4] == "+998" {
		for i := 4; i < 13; i++{
			if phoneNumber[i] >= 48 && phoneNumber[i] <= 57{
				istrue = true
			}else{
				return false
			}
		}
	}
	return istrue
}





