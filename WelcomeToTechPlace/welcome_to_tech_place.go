package techplace

import "strings"

func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	lineOfStars := strings.Repeat("*", numStarsPerLine)
	return lineOfStars + "\n" + welcomeMsg + "\n" + lineOfStars
}

func CleanupMessage(oldMsg string) string {
	return strings.Trim(oldMsg, "* \n")
}
