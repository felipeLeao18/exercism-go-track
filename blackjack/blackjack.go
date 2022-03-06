package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	flag := map[string]int{
		"ace":   11,
		"other": 0,
		"eight": 8,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"nine":  9,
		"ten":   10,
		"jack":  10,
		"queen": 10,
		"king":  10,
	}

	return flag[card]
}

// IsBlackjack returns true if the player has a blackjack, false otherwise.
func IsBlackjack(card1, card2 string) bool {
	var card1Value int = ParseCard(card1)
	var card2Value int = ParseCard(card2)
	blackjack := card1Value+card2Value == 21

	return blackjack

}

// LargeHand implements the decision tree for hand scores larger than 20 points.
func LargeHand(isBlackjack bool, dealerScore int) string {
	if !isBlackjack {
		return "P"
	}
	if dealerScore < 10 {
		return "W"
	}

	return "S"
}

// SmallHand implements the decision tree for hand scores with less than 21 points.
func SmallHand(handScore, dealerScore int) string {
	if handScore > 17 {
		return "S"
	}
	if handScore <= 11 {
		return "H"
	}
	if handScore >= 12 && handScore <= 16 && dealerScore <= 6 {
		return "S"
	}
	if handScore >= 12 && handScore <= 16 && dealerScore >= 7 {
		return "H"
	}

	return "S"
}
