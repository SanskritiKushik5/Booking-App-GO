package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

const confName = "Go Conf"
const confTkts int = 50
var remainTkts int = 50
// var bookings = make([]map[string]string, 0)
var bookings = make([]UserData, 0)
 
type UserData struct {
	fname string
	lname string
	email string
	tkts int
}

var wg = sync.WaitGroup{}

func main() {
	greetUsers()

	for {
		fname, lname, email, tkts := getInput()
		isValidName, isValidEmail, isValidTkt := helper.ValidateInput(fname, lname, email, tkts, remainTkts)
		
		if isValidName && isValidEmail && isValidTkt {
			bookTkt(tkts, fname, lname, email)

			wg.Add(1)
			// sendTkt(tkts, fname)      // concurrency
			go sendTkt(tkts, fname) // same thread go routines
			fmt.Printf("Bookings: %v\n", getFNames())

			if remainTkts == 0 {
				fmt.Println("Tickets sold out!")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("Name Issue")
			}
			if !isValidEmail {
				fmt.Println("Email Issue")
			}
			if !isValidTkt {
				fmt.Println("Ticket Issue")
			}
			fmt.Println("Wrong Input data")
		}
	} 
	wg.Wait()
}

func greetUsers() {
	fmt.Println("Welcome to", confName)
	fmt.Printf("Total tickets: %v\t", confTkts)
	fmt.Printf("Remaining tickets: %v\n", remainTkts)
	fmt.Println("Book Now!")
}

func getFNames() []string {
	fnames := []string{}
	for _, booking := range bookings {
		fnames = append(fnames, booking.fname)
	}
	return fnames
}

func getInput() (string, string, string, int) {
	var fname string 
	var lname string
	var email string
	var tkts int
	fmt.Printf("Enter first name: ")
	fmt.Scan(&fname)
	fmt.Printf("Enter last name: ")
	fmt.Scan(&lname)
	fmt.Printf("Enter email id: ")
	fmt.Scan(&email)
	fmt.Printf("Enter tickets: ")
	fmt.Scan(&tkts)
	return fname, lname, email, tkts
}

func sendTkt(tkts int, fname string) {
	time.Sleep(5 * time.Second) // for concurrency
	var tkt = fmt.Sprintf("%v tickets for %v\n", tkts, fname)
	fmt.Println("----------------------")
	fmt.Printf("%v\n", tkt)
	wg.Done()
}