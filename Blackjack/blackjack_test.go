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

func TestFirstTurn_WhenPlayerHasTwoAces(t *testing.T) {
	card1 := "ace"
	card2 := "ace"
	expectedFirstTurn := "P"

	testFirstTurn(t, card1, card2, "queen", expectedFirstTurn)
}

func TestFirstTurn_WhenPlayerHasWeakHand(t *testing.T) {
	card1 := "2"
	card2 := "3"
	expectedFirstTurn := "H"

	testFirstTurn(t, card1, card2, "ace", expectedFirstTurn)
}

func TestFirstTurn_WhenPlayerHitWithLowerMediumHand(t *testing.T) {
	card1 := "queen"
	card2 := "two"
	dealerCard := "ace"
	expectedFirstTurn := "H"

	testFirstTurn(t, card1, card2, dealerCard, expectedFirstTurn)
}

func TestFirstTurn_WhenPlayerStandWithLowerMediumHand(t *testing.T) {
	card1 := "seven"
	card2 := "five"
	dealerCard := "three"
	expectedFirstTurn := "S"

	testFirstTurn(t, card1, card2, dealerCard, expectedFirstTurn)
}

func TestFirstTurn_WhenPlayerStandWithUpperMediumHand(t *testing.T) {
	card1 := "king"
	card2 := "six"
	dealerCard := "three"
	expectedFirstTurn := "S"

	testFirstTurn(t, card1, card2, dealerCard, expectedFirstTurn)
}

func TestFirstTurn_WhenPlayerHitWithUpperMediumHand(t *testing.T) {
	card1 := "eight"
	card2 := "eight"
	expectedFirstTurn := "H"

	testFirstTurn(t, card1, card2, "nine", expectedFirstTurn)
}

func TestFirstTurn_WhenPlayerHitWithMediumHand(t *testing.T) {
	card1 := "seven"
	card2 := "six"
	dealerCard := "nine"
	expectedFirstTurn := "H"

	testFirstTurn(t, card1, card2, dealerCard, expectedFirstTurn)
}

func TestFirstTurn_WhenPlayerStandWithMediumHand(t *testing.T) {
	card1 := "ace"
	card2 := "four"
	dealerCard := "three"
	expectedFirstTurn := "S"

	testFirstTurn(t, card1, card2, dealerCard, expectedFirstTurn)
}

func TestFirstTurn_WhenPlayerHasLowerStrongHand(t *testing.T) {
	card1 := "ace"
	card2 := "six"
	expectedFirstTurn := "S"

	testFirstTurn(t, card1, card2, "ace", expectedFirstTurn)
}

func TestFirstTurn_WhenPlayerHasUpperStrongHand(t *testing.T) {
	card1 := "queen"
	card2 := "king"
	expectedFirstTurn := "S"

	testFirstTurn(t, card1, card2, "ace", expectedFirstTurn)
}

func TestFirstTurn_WhenPlayerHasStrongHand(t *testing.T) {
	card1 := "nine"
	card2 := "eight"
	expectedFirstTurn := "S"

	testFirstTurn(t, card1, card2, "ace", expectedFirstTurn)
}

func TestFirstTurn_WhenPlayerWinWithBlackJack(t *testing.T) {
	card1 := "ace"
	card2 := "queen"
	dealerCard := "nine"
	expectedFirstTurn := "W"

	testFirstTurn(t, card1, card2, dealerCard, expectedFirstTurn)
}

func TestFirstTurn_WhenPlayerStandWithBlackJack(t *testing.T) {
	card1 := "ace"
	card2 := "queen"
	dealerCard := "ten"
	expectedFirstTurn := "S"

	testFirstTurn(t, card1, card2, dealerCard, expectedFirstTurn)
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

func testFirstTurn(t *testing.T, card1, card2, dealerCard, expectedFirstTurn string) {
	actualFirstTurn := FirstTurn(card1, card2, dealerCard)

	if expectedFirstTurn != actualFirstTurn {
		t.Errorf(
			"FirstTurn(%s, %s) = %s. Expected %s",
			card1,
			card2,
			actualFirstTurn,
			expectedFirstTurn)
	}
}
