package main

import (
	"fmt"
)

func bookTkt(tkts int, fname string, lname string, email string) {
	remainTkts = remainTkts - tkts

	var userData = UserData {
		fname: fname,
		lname: lname,
		email: email,
		tkts: tkts,
	}
	// var userData = make(map[string]string)
	// userData["fname"] = fname
	// userData["lname"] = lname
	// userData["email"] = email
	// userData["tkts"] = strconv.FormatInt(int64(tkts), 10)
	bookings = append(bookings, userData)

	fmt.Printf("%v booked %v tickets.\n", fname, tkts)
	fmt.Printf("Remaining Tickets: %v\n", remainTkts)

}