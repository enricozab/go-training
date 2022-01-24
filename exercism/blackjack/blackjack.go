// Package blackjack simulates the first turn of a Blackjack game
package blackjack

var cards = map[string]int{
	"ace":   11,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"ten":   10,
	"jack":  10,
	"queen": 10,
	"king":  10,
	"other": 0,
}

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten":
		return 10
	case "jack":
		return 10
	case "queen":
		return 10
	case "king":
		return 10
	default:
		return 0
	}
}

// IsBlackjack returns true if the player has a blackjack, false otherwise.
func IsBlackjack(card1, card2 string) bool {
	return cards[card1]+cards[card2] == 21
}

// LargeHand implements the decision tree for hand scores larger than 20 points.
func LargeHand(isBlackjack bool, dealerScore int) string {
	if isBlackjack && dealerScore != 11 && dealerScore != 10 {
		return "W"
	} else if !isBlackjack && dealerScore == 11 {
		return "P"
	} else if dealerScore == 11 || dealerScore == 10 {
		return "S"
	}
	return "H"
}

// SmallHand implements the decision tree for hand scores with less than 21 points.
func SmallHand(handScore, dealerScore int) string {
	if (handScore >= 12 && handScore <= 16 && dealerScore < 7) || handScore >= 17 {
		return "S"
	}
	return "H"
}
