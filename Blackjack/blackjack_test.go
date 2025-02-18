package blackjack

import "testing"

func TestParseCard_WhenCardIsTwo(t *testing.T) {
	card := "two"
	expectedCardValue := 2

	testParseCard(t, card, expectedCardValue)
}

func TestParseCard_WhenCardIsThree(t *testing.T) {
	card := "three"
	expectedCardValue := 3

	testParseCard(t, card, expectedCardValue)
}

func TestParseCard_WhenCardIsFour(t *testing.T) {
	card := "four"
	expectedCardValue := 4

	testParseCard(t, card, expectedCardValue)
}

func TestParseCard_WhenCardIsFive(t *testing.T) {
	card := "five"
	expectedCardValue := 5

	testParseCard(t, card, expectedCardValue)
}

func TestParseCard_WhenCardIsSix(t *testing.T) {
	card := "six"
	expectedCardValue := 6

	testParseCard(t, card, expectedCardValue)
}

func TestParseCard_WhenCardIsSeven(t *testing.T) {
	card := "seven"
	expectedCardValue := 7

	testParseCard(t, card, expectedCardValue)
}

func TestParseCard_WhenCardIsEight(t *testing.T) {
	card := "eight"
	expectedCardValue := 8

	testParseCard(t, card, expectedCardValue)
}

func TestParseCard_WhenCardIsNine(t *testing.T) {
	card := "nine"
	expectedCardValue := 9

	testParseCard(t, card, expectedCardValue)
}

func TestParseCard_WhenCardIsTen(t *testing.T) {
	card := "ten"
	expectedCardValue := 10

	testParseCard(t, card, expectedCardValue)
}

func TestParseCard_WhenCardIsJack(t *testing.T) {
	card := "jack"
	expectedCardValue := 10

	testParseCard(t, card, expectedCardValue)
}

func TestParseCard_WhenCardIsQueen(t *testing.T) {
	card := "queen"
	expectedCardValue := 10

	testParseCard(t, card, expectedCardValue)
}

func TestParseCard_WhenCardIsKing(t *testing.T) {
	card := "king"
	expectedCardValue := 10

	testParseCard(t, card, expectedCardValue)
}

func TestParseCard_WhenCardIsAce(t *testing.T) {
	card := "ace"
	expectedCardValue := 11

	testParseCard(t, card, expectedCardValue)
}

func TestParseCard_WhenCardIsOther(t *testing.T) {
	card := "1"
	expectedCardValue := 0

	testParseCard(t, card, expectedCardValue)
}

func testParseCard(t *testing.T, card string, expectedCardValue int) {
	actualCardValue := ParseCard(card)

	if expectedCardValue != actualCardValue {
		t.Errorf(
			"ParseCard(%s) = %d. Expected %d",
			card,
			actualCardValue,
			expectedCardValue)
	}
}
