// Package techplace generates the text for the new display
package techpalace

import (
	"fmt"
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	var border, msg string

	for ctr := 0; ctr < numStarsPerLine; ctr++ {
		border = border + "*"
	}

	msg = fmt.Sprintf("%v\n%v\n%v", border, welcomeMsg, border)

	return msg
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	msg := strings.ReplaceAll(oldMsg, "*", "")
	msg = strings.TrimSpace(msg)

	return msg
}
