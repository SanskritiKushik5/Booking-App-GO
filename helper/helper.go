package helper

import "strings"


func ValidateInput(fname string, lname string, email string, tkts int, remainTkts int) (bool, bool, bool) {
	isValidName := len(fname) >= 2 && len(lname) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTkt := tkts > 0 && tkts <= remainTkts
	return isValidName, isValidEmail, isValidTkt
}