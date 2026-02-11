package main

import (
	"fmt"
)

const maxNoOfSeats = 300

// creating a struct so i can be easily extensible in future 
type FlightData struct {
	Data [maxNoOfSeats]bool // Taking bool is beneficial it take only 1 byte
}
// creating display functions which display total available seats and also print all available seats
func (flightData *FlightData) listAvailableSeats() {
	fmt.Printf("List of available seats\n")
	for id, val := range flightData.Data {
		if !val { // means false
			fmt.Printf("%v ", id+1)
		}
	}
	fmt.Printf("\n")
}


// following function to calculate total no of seats and using pointer to save memory as  i need to do modification based on updated flight data
func(flightData * FlightData) totalAvailableSeats(){

	total:=0

	for i:=0;i<300;i++{

		if !flightData.Data[i]{
			total++
		}
	}

	fmt.Println("Total no of available Seats are",total)
}
func (flightData *FlightData) bookSeat(seat int) error {
	if seat < 1 || seat > 300 {
		return fmt.Errorf("Please enter valid seat number\n")

	} else if flightData.Data[seat-1] {
		return fmt.Errorf("This seat is already booked . Enter new seat number\n")
	} else {
		flightData.Data[seat-1] = true
	}
	return nil
}

func main() {

	fmt.Println("Flight Booking System")

	flightData := FlightData{} // intialize flightData

	fmt.Printf("Welcome to Flight Booking System..\n\n")

	var input int

	for {

		fmt.Printf("1: List all available seats \n2: Book Seat\n3: Total no of available seats\n4: Exit Program\nEnter your command: ")

		if _ , err := fmt.Scanln(&input) ; err != nil{
			fmt.Printf("Pleasee enter valid input\n")
			continue
		}
		switch input {

		case 1:
			flightData.listAvailableSeats() // Display all the flight data
		case 2:
			var SeatNumber int
			fmt.Printf("Enter Seat Number: ")
			if _, err := fmt.Scanln(&SeatNumber); err !=nil{ // err case 
				fmt.Printf("Please enter valid seat number")
				continue

			}
			if err:= flightData.bookSeat(SeatNumber); err !=nil{

				fmt.Printf("Failed to book ticket %v\n", err)
				continue

			}

		case 3:
			flightData.totalAvailableSeats() // total no of available seats
		case 4:
			fmt.Println("Thank you for visiting") //exit
			return
		}
	}


}
