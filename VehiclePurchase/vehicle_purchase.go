package purchase

func NeedsLicense(kind string) bool {
	return kind == "car" || kind == "truck"
}

func ChooseVehicle(option1, option2 string) string {
	choice := ""

	if option1 < option2 {
		choice = option1
	} else {
		choice = option2
	}

	return choice + " is clearly the better choice."
}

func CalculateResellPrice(originalPrice, age float64) float64 {
	if age < 3 {
		return originalPrice * 0.8
	}

	if age >= 10 {
		return originalPrice * 0.5
	}

	return originalPrice * 0.7
}
