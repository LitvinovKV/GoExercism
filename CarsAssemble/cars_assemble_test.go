package cars

import "testing"

func TestCalculateWorkingCarsPerHour_WhenProuctionRateIs0(t *testing.T) {
	productionRate := 0
	successRate := 100.
	expectedWorkingCarsPerHour := 0.

	testCalculateWorkingCarsPerHour(t, productionRate, successRate, expectedWorkingCarsPerHour)
}

func TestCalculateWorkingCarsPerHour_WhenSuccessRateIs0(t *testing.T) {
	productionRate := 100
	successRate := 0.
	expectedWorkingCarsPerHour := 0.

	testCalculateWorkingCarsPerHour(t, productionRate, successRate, expectedWorkingCarsPerHour)
}

func TestCalculateWorkingCarsPerHour_WhenSuccessRateIs100(t *testing.T) {
	productionRate := 100
	successRate := 100.
	expectedWorkingCarsPerHour := 100.

	testCalculateWorkingCarsPerHour(t, productionRate, successRate, expectedWorkingCarsPerHour)
}

func TestCalculateWorkingCarsPerHour_WhenSuccessRateWithRemainder(t *testing.T) {
	productionRate := 100
	successRate := 17.11
	expectedWorkingCarsPerHour := 17.11

	testCalculateWorkingCarsPerHour(t, productionRate, successRate, expectedWorkingCarsPerHour)
}

func TestCalcualteWorkingCarsPerMinute_WhenProductionRateIs0(t *testing.T) {
	productionRate := 0
	successRate := 100.
	expectedWorkingCarsPerMinute := 0

	testCalcualteWorkingCarsPerMinute(t, productionRate, successRate, expectedWorkingCarsPerMinute)
}

func TestCalcualteWorkingCarsPerMinute_WhenSuccessRateIs100(t *testing.T) {
	productionRate := 100
	successRate := 100.
	expectedWorkingCarsPerMinute := 1

	testCalcualteWorkingCarsPerMinute(t, productionRate, successRate, expectedWorkingCarsPerMinute)
}

func TestCalcualteWorkingCarsPerMinute_WhenSuccessRateIs0(t *testing.T) {
	productionRate := 100
	successRate := 0.
	expectedWorkingCarsPerMinute := 0

	testCalcualteWorkingCarsPerMinute(t, productionRate, successRate, expectedWorkingCarsPerMinute)
}

func TestCalcualteWorkingCarsPerMinute_WhenSuccessRateWithRemainder(t *testing.T) {
	productionRate := 100
	successRate := 17.11
	expectedWorkingCarsPerMinute := 0

	testCalcualteWorkingCarsPerMinute(t, productionRate, successRate, expectedWorkingCarsPerMinute)
}

func TestCalculateCost_WhenCarsCountIs0(t *testing.T) {
	carsCount := 0
	var expectedCost uint = 0

	testCalculateCost(t, carsCount, expectedCost)
}

func TestCalculateCost_WhenCarsCountIs9(t *testing.T) {
	carsCount := 9
	var expectedCost uint = 90_000

	testCalculateCost(t, carsCount, expectedCost)
}

func TestCalculateCost_WhenCarsCountIs10(t *testing.T) {
	carsCount := 10
	var expectedCost uint = 95_000

	testCalculateCost(t, carsCount, expectedCost)
}

func TestCalculateCost_WhenCarsCountIs11(t *testing.T) {
	carsCount := 11
	var expectedCost uint = 105_000

	testCalculateCost(t, carsCount, expectedCost)
}

func TestCalculateCost_WhenCarsCountIs777(t *testing.T) {
	carsCount := 777
	var expectedCost uint = 7_385_000

	testCalculateCost(t, carsCount, expectedCost)
}

func testCalculateWorkingCarsPerHour(
	t *testing.T,
	productionRate int,
	successRate float64,
	expectedWorkingCarsPerHour float64) {

	actualWorkingCarsPerHour := CalculateWorkingCarsPerHour(productionRate, successRate)

	if expectedWorkingCarsPerHour != actualWorkingCarsPerHour {
		t.Errorf(
			"CalculateWorkingCarsPerHour(%d, %f) = %f. Expected %f",
			productionRate,
			successRate,
			actualWorkingCarsPerHour,
			expectedWorkingCarsPerHour)
	}
}

func testCalcualteWorkingCarsPerMinute(
	t *testing.T,
	productionRate int,
	successRate float64,
	expectedWorkingCarsPerMinute int) {

	actualWorkingCarsPerMinute := CalcualteWorkingCarsPerMinute(productionRate, successRate)

	if expectedWorkingCarsPerMinute != actualWorkingCarsPerMinute {
		t.Errorf(
			"CalcualteWorkingCarsPerMinute(%d, %f) = %d. Expected %d",
			productionRate,
			successRate,
			actualWorkingCarsPerMinute,
			expectedWorkingCarsPerMinute)
	}
}

func testCalculateCost(
	t *testing.T,
	carsCount int,
	expectedCost uint) {

	actualCost := CalculateCost(carsCount)

	if expectedCost != actualCost {
		t.Errorf(
			"CalculateCost(%d) = %d. Expected %d",
			carsCount,
			actualCost,
			expectedCost)
	}
}
