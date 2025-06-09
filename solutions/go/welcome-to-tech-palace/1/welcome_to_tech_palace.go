package techpalace

import (
    "strings"
)

var b strings.Builder

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
    b.Reset()
	b.WriteString("Welcome to the Tech Palace, ")
    b.WriteString(strings.ToUpper(customer))
    return b.String()
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	b.Reset()
    b.WriteString(strings.Repeat("*", numStarsPerLine))
    b.WriteString("\n")
    b.WriteString(welcomeMsg)
    b.WriteString("\n")
    b.WriteString(strings.Repeat("*", numStarsPerLine))
    return b.String()
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	b.Reset()
    b.WriteString(strings.TrimSpace(strings.ReplaceAll(oldMsg, "*", "")))
    return b.String()
}
