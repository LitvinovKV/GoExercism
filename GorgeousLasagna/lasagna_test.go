package lasagna

import "testing"

func TestRemainingOvenTime(t *testing.T) {
	actualMinutesInOven := 15
	expectedRemainingOvenTime := 25

	actualRemainingOvenTime := RemainingOvenTime(actualMinutesInOven)

	if actualRemainingOvenTime != expectedRemainingOvenTime {
		t.Errorf(
			"RemainingOvenTime(%d) = %d. Expected %d",
			actualMinutesInOven,
			actualRemainingOvenTime,
			expectedRemainingOvenTime)
	}
}

func TestPreparationTimeForOneLayer(t *testing.T) {
	numberOfLayers := 1
	expectedPreparationTime := 2

	testPreparationTime(t, numberOfLayers, expectedPreparationTime)
}

func TestPreparationTimeForMultipleLayers(t *testing.T) {
	numberOfLayers := 5
	expectedPreparationTime := 10

	testPreparationTime(t, numberOfLayers, expectedPreparationTime)
}

func TestElapsedTimeForOneLayer(t *testing.T) {
	numberOfLayers := 1
	actualMinutesInOven := 30
	expectedElapsedTime := 32

	testElapsedTime(
		t,
		numberOfLayers,
		actualMinutesInOven,
		expectedElapsedTime)
}

func TestElapsedTimeForMultipleLayers(t *testing.T) {
	numberOfLayers := 5
	actualMinutesInOven := 30
	expectedElapsedTime := 40

	testElapsedTime(
		t,
		numberOfLayers,
		actualMinutesInOven,
		expectedElapsedTime)
}

func testPreparationTime(
	t *testing.T,
	numberOfLayers, expectedPreparationTime int) {

	actualPreparationTime := PreparationTime(numberOfLayers)

	if actualPreparationTime != expectedPreparationTime {
		t.Errorf(
			"PreparationTime(%d) = %d. Expected %d",
			numberOfLayers,
			actualPreparationTime,
			expectedPreparationTime)
	}
}

func testElapsedTime(
	t *testing.T,
	numberOfLayers, actualMinutesInOven, expectedElapsedTime int) {

	actualElapsedTime := ElapsedTime(numberOfLayers, actualMinutesInOven)

	if actualElapsedTime != expectedElapsedTime {
		t.Errorf(
			"ElapsedTime(%d, %d) = %d. Expected %d",
			numberOfLayers,
			actualMinutesInOven,
			actualElapsedTime,
			expectedElapsedTime)
	}
}
