package partyrobot

import "fmt"

// Welcome greets a person by name.
func Welcome(name string) string {
	welcomeMessage := "Welcome to my party, " + name + "!"
	return welcomeMessage
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	happyBirthdayMessage := "Happy birthday " + name + "! You are now " + fmt.Sprintf("%d", age) + " years old!"
	return happyBirthdayMessage
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	return "Welcome to my party, " + name + "!" + "\nYou have been assigned to table " + fmt.Sprintf("%.3d", table) + ". Your table is " + direction + ", exactly " + fmt.Sprintf("%.1f", distance) + " meters from here.\nYou will be sitting next to " + neighbor + "."
}
