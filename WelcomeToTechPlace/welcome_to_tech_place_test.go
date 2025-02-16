package techplace

import (
	"testing"
)

func TestWelcomeMessage_WhenCustomerNameHasOnlyLowercaseLetters(t *testing.T) {
	customer := "luffy"
	expectedWelcomeMessage := "Welcome to the Tech Palace, LUFFY"

	testWelcomeMessage(t, customer, expectedWelcomeMessage)
}

func TestWelcomeMessage_WhenCustomerNameHasOnlyUppercaseLetters(t *testing.T) {
	customer := "LUFFY"
	expectedWelcomeMessage := "Welcome to the Tech Palace, LUFFY"

	testWelcomeMessage(t, customer, expectedWelcomeMessage)
}

func TestWelcomeMessage_WhenCustomerNameHasSeveralUppercaseLetters(t *testing.T) {
	customer := "LuffY"
	expectedWelcomeMessage := "Welcome to the Tech Palace, LUFFY"

	testWelcomeMessage(t, customer, expectedWelcomeMessage)
}

func TestAddBorder_WithoutStars(t *testing.T) {
	welcomeMsg := "welcome message"
	numStarsPerLine := 0
	expectedBorderedMsg := "\nwelcome message\n"

	testAddBorder(t, welcomeMsg, numStarsPerLine, expectedBorderedMsg)
}

func TestAddBorder_WithSeveralStars(t *testing.T) {
	welcomeMsg := "welcome message"
	numStarsPerLine := 10
	expectedBorderedMsg := "**********\nwelcome message\n**********"

	testAddBorder(t, welcomeMsg, numStarsPerLine, expectedBorderedMsg)
}

func TestCleanUpMessage_WhenOldMessageWithoutSpecialChars(t *testing.T) {
	oldMsg := "some message"
	expectedMsg := "some message"

	testCleanupMessage(t, oldMsg, expectedMsg)
}

func TestCleanUpMessage_WhenOldMessageHasOnlySpecialChars(t *testing.T) {
	oldMsg := "\n *"
	expectedMsg := ""

	testCleanupMessage(t, oldMsg, expectedMsg)
}

func TestCleanupMessage_WhenOldMessageContainsTextWithSpecialChars(t *testing.T) {
	oldMsg := "\n* some message*\n "
	expectedMsg := "some message"

	testCleanupMessage(t, oldMsg, expectedMsg)
}

func testWelcomeMessage(t *testing.T, customer, expectedWelcomeMessage string) {
	actualWelcomeMessage := WelcomeMessage(customer)

	if expectedWelcomeMessage != actualWelcomeMessage {
		t.Errorf(
			"WelcomeMessage(%s) = \"%s\". Expected: \"%s\"",
			customer,
			actualWelcomeMessage,
			expectedWelcomeMessage)
	}
}

func testAddBorder(
	t *testing.T,
	welcomeMsg string,
	numStarsPerLine int,
	expectedBorderedMessage string) {

	actualBorderedMessage := AddBorder(welcomeMsg, numStarsPerLine)

	if expectedBorderedMessage != actualBorderedMessage {
		t.Errorf(
			"AddBorder(%s, %d) = \"%s\". Expected: \"%s\"",
			welcomeMsg,
			numStarsPerLine,
			actualBorderedMessage,
			expectedBorderedMessage)
	}
}

func testCleanupMessage(t *testing.T, oldMsg, expectedMsg string) {
	actualMsg := CleanupMessage(oldMsg)

	if expectedMsg != actualMsg {
		t.Errorf(
			"CleanupMessage(%s) = \"%s\". Expected: \"%s\"",
			oldMsg,
			actualMsg,
			expectedMsg)
	}
}
