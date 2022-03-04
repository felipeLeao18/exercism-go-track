package techpalace

import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	message := "Welcome to the Tech Palace, " + strings.ToUpper(customer)
	return message
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	border := strings.Repeat("*", numStarsPerLine)
	message := border + "\n" + welcomeMsg + "\n" + border
	return message
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	messageWithoutBorders := strings.ReplaceAll(oldMsg, "*", "")
	removeSpaces := strings.TrimSpace(messageWithoutBorders)
	return removeSpaces
}
