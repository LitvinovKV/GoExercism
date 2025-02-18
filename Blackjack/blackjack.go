package blackjack

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

func FirstTurn(card1, card2, dealerCard string) string {
	sum := ParseCard(card1) + ParseCard(card2)

	switch {
	case sum == 22:
		return "P"

	case sum == 21:
		if ParseCard(dealerCard) < 10 {
			return "W"
		} else {
			return "S"
		}

	case sum >= 17 && sum <= 20:
		return "S"

	case sum >= 12 && sum <= 16:
		if ParseCard(dealerCard) >= 7 {
			return "H"
		} else {
			return "S"
		}

	default:
		return "H"
	}
}
