package purchase

import "testing"

func TestNeedsLicense_WhenKindIsCar(t *testing.T) {
	kind := "car"
	expectedIsLicenseNeeded := true

	testNeedsLicense(t, kind, expectedIsLicenseNeeded)
}

func TestNeedsLicense_WhenKindIsTruck(t *testing.T) {
	kind := "truck"
	expectedIsLicenseNeeded := true

	testNeedsLicense(t, kind, expectedIsLicenseNeeded)
}

func TestNeedsLicense_WhenLicenseIsNoNeeded(t *testing.T) {
	kind := "bicycle"
	expectedIsLicenseNeeded := false

	testNeedsLicense(t, kind, expectedIsLicenseNeeded)
}

func TestChooseVehicle_WhenOption1IsGreaterThenOption2(t *testing.T) {
	option1 := "vehicle B"
	option2 := "vehicle A"
	expectedOption := "vehicle A is clearly the better choice."

	testChooseVehicle(t, option1, option2, expectedOption)
}

func TestChooseVehicle_WhenOption2IsGreaterThenOption1(t *testing.T) {
	option1 := "vehicle A"
	option2 := "vehicle B"
	expectedOption := "vehicle A is clearly the better choice."

	testChooseVehicle(t, option1, option2, expectedOption)
}

func TestChooseVehicle_WhenOption1AndOption2AreEquals(t *testing.T) {
	option1 := "vehicle A"
	option2 := "vehicle A"
	expectedOption := "vehicle A is clearly the better choice."

	testChooseVehicle(t, option1, option2, expectedOption)
}

func TestCalculateResellPrice_WhenVehicleIs2Years(t *testing.T) {
	originalPrice := 1000.
	age := 2.
	expectedPrice := 800.

	testCalculateResellPrice(t, originalPrice, age, expectedPrice)
}

func TestCalculateResellPrice_WhenVehicleIs3Years(t *testing.T) {
	originalPrice := 1000.
	age := 3.
	expectedPrice := 700.

	testCalculateResellPrice(t, originalPrice, age, expectedPrice)
}

func TestCalculateResellPrice_WhenVehicleIs9Years(t *testing.T) {
	originalPrice := 1000.
	age := 9.
	expectedPrice := 700.

	testCalculateResellPrice(t, originalPrice, age, expectedPrice)
}

func TestCalculateResellPrice_WhenVehicleIs10Years(t *testing.T) {
	originalPrice := 1000.
	age := 10.
	expectedPrice := 500.

	testCalculateResellPrice(t, originalPrice, age, expectedPrice)
}

func TestCalculateResellPrice_WhenVehicleIs17Years(t *testing.T) {
	originalPrice := 1000.
	age := 17.
	expectedPrice := 500.

	testCalculateResellPrice(t, originalPrice, age, expectedPrice)
}

func testNeedsLicense(t *testing.T, kind string, expectedIsLicenseNeeded bool) {
	actualIsLicenseNeeded := NeedsLicense(kind)

	if expectedIsLicenseNeeded != actualIsLicenseNeeded {
		t.Errorf(
			"NeedsLicense(%s) = %t. Expected %t",
			kind,
			actualIsLicenseNeeded,
			expectedIsLicenseNeeded)
	}
}

func testChooseVehicle(t *testing.T, option1, option2, expectedOption string) {
	actualOption := ChooseVehicle(option1, option2)

	if expectedOption != actualOption {
		t.Errorf(
			"ChooseVehicle(%s, %s) = \"%s\". Expected \"%s\"",
			option1,
			option2,
			actualOption,
			expectedOption)
	}
}

func testCalculateResellPrice(t *testing.T, originalPrice, age, expectedPrice float64) {
	actualPrice := CalculateResellPrice(originalPrice, age)

	if expectedPrice != actualPrice {
		t.Errorf(
			"CalculateResellPrice(%f, %f) = %f. Expected %f",
			originalPrice,
			age,
			actualPrice,
			expectedPrice)
	}
}
