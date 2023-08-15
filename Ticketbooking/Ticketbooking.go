package main

import(
	"fmt"
	"strings"
	"time"

)
	var conferenceName = "Go conference"
	const conferenceTickets int = 50
	var reaminingTickets uint = 50
	var bookings = make([]userdata,0)

	type userdata struct{
		firstname string
		lastname string
		email string
		numberofticket uint
	}
func main(){
	greetuser()

	for{
	
	firstname,lastname,email,userTicket:= userinput()

	isvalidname, isvalidemail,isvaliduserticket := validuserinput(firstname, lastname, email,  userTicket)

	if isvalidname &&  isvalidemail && isvaliduserticket {

		bookTicket(userTicket,firstname,lastname,email)

		go sendticket(userTicket,firstname, lastname,email)

		fmt.Printf("The first name of ticket is %v\n\n",printfirstname())

	 
	if reaminingTickets==0{
		fmt.Println("Our conference is booked out. please comeback next year")
		break
		}
	}else{
			if !isvalidname{
				fmt.Println(" The character of First name or last name you entered is less than 3")
			}
			if !isvalidemail{
				fmt.Println(" The email you entered doesn't contain @ sign")
			}
			if !isvaliduserticket{
				fmt.Println(" The number of ticket you entered is invalid")
			}
		}
	}
}

func greetuser(){

	fmt.Printf("Welcome to %v booking Application\n",conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n",conferenceTickets,reaminingTickets)
	fmt.Printf("Get your ticket to attend\n")
}

func printfirstname() []string{
	firstnames := []string{}

	for _,booking := range bookings {
		firstnames = append(firstnames,booking.firstname) 
	}
 return firstnames
}

func validuserinput(firstname string , lastname string , email string , userTicket uint)(bool,bool,bool){
	isvalidname := len(firstname)>=3 && len(lastname)>=3
	isvalidemail := strings.Contains(email,"@")
	isvaliduserticket := userTicket >0 && userTicket <= reaminingTickets
	return isvalidname,isvalidemail,isvaliduserticket
}

func userinput()(string , string , string, uint){
	var firstname string
	var lastname string
	var email string
	var userTicket uint 
    
	fmt.Println("Enter your first name :")
	fmt.Scan(&firstname)

	fmt.Println("Enter your last name :")
	fmt.Scan(&lastname)

	fmt.Println("Enter your Email address :")
	fmt.Scan(&email)

	fmt.Println("Enter the no. of ticket you want to book :")
	fmt.Scan(&userTicket)

	return firstname, lastname, email,userTicket
}

func bookTicket( userTicket uint,firstname string , lastname string ,email string ){
		reaminingTickets = reaminingTickets-userTicket

		 var userData = userdata {
			firstname: firstname,
			lastname: lastname,
			email: email,
			numberofticket: userTicket,
		}
		bookings = append(bookings,userData)
		
		fmt.Printf("The list of booking is %v\n",bookings)

		fmt.Printf("Thankyou %v %v for booking %v Tickets. You will receive confirmation email at %v\n",firstname,lastname,userTicket,email)
		fmt.Printf("%v tickets are reamining for %v\n\n",reaminingTickets,conferenceName)
}

func sendticket(usertickets uint , firstname string , lastname string,email string){
	time.Sleep(5 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v ",usertickets,firstname,lastname)
	fmt.Printf("***************************\n")
	fmt.Printf("Sending ticket:\n%v \nEmail address: %v\n\n",ticket, email)
	fmt.Printf("***************************\n\n")
	
}