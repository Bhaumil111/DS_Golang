package main

import (
	"fmt"
)

const maxNoOfSeats = 300

// creating a struct so i can extended it later and create method which attach type
type FlightData struct {
	Data [maxNoOfSeats]bool // Taking bool is beneficial it take only 1 byte
}

func (flightData *FlightData) display() {

	fmt.Printf("List of available seats\n")
	count := 0
	for id, val := range flightData.Data {
		if !val { // means false

			count++
			fmt.Printf("%v ", id+1)
		}
	}
	fmt.Printf("\n")

	fmt.Println("Total no of available seats ", count)

}

func (flightData *FlightData) bookSeat(seat int) error {
	if seat < 1 || seat > 300 {
		fmt.Errorf("Please enter valid seat number\n")

	} else if flightData.Data[seat-1] {
		fmt.Errorf("This seat is already booked . Enter new seat number\n")
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

		fmt.Printf("1: For list available seats\n 2: Book Seat\n 3: Exit Program\n Enter your command: ")

		//fmt.Scan(&input)

		if err, _ := fmt.Scanln(&input); err != nil {

		}
		switch input {
		case 1:
			flightData.display() // Display all the flight data
		case 2:
			var SeatNumber int
			fmt.Println("Enter the Seat Number ")
			fmt.Scan(&SeatNumber)
			flightData.bookSeat(SeatNumber) // Book Seat
		case 3:
			fmt.Println("Thank you for visiting")
			return
		}
	}

	return

}
