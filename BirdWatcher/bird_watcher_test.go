package birdwatcher

import "testing"

func TestTotalBirdCount_WhenBirdsArrayIsEmpty(t *testing.T) {
	birds := []int{}
	expectedBirdsCount := 0

	testTotalBirdCount(t, birds, expectedBirdsCount)
}

func TestTotalBirdCount_WhenBirdsArrayIsNil(t *testing.T) {
	var birds []int = nil
	expectedBirdsCount := 0

	testTotalBirdCount(t, birds, expectedBirdsCount)
}

func TestTotalBirdCount_WithSomeBirds(t *testing.T) {
	birds := []int{1, 2, 3}
	expectedBirdsCount := 6

	testTotalBirdCount(t, birds, expectedBirdsCount)
}

func TestTotalBirdCount_WithLargeCountOfBirds(t *testing.T) {
	birds := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expectedBirdsCount := 550

	testTotalBirdCount(t, birds, expectedBirdsCount)
}

func TestBirdsInWeek_WhenOneWeek(t *testing.T) {
	birds := []int{1, 2, 3, 4, 5, 6, 7}
	week := 1
	expectedCountOfBirds := 28

	testBirdsInWeek(t, birds, week, expectedCountOfBirds)
}

func TestBirdsInWeek_WhenFirstWeek(t *testing.T) {
	birds := []int{1, 2, 3, 4, 5, 6, 7, 11, 12, 13, 14, 15, 16, 17}
	week := 1
	expectedCountOfBirds := 28

	testBirdsInWeek(t, birds, week, expectedCountOfBirds)
}

func TestBirdsInWeek_WhenLastWeek(t *testing.T) {
	birds := []int{11, 12, 13, 14, 15, 16, 17, 1, 2, 3, 4, 5, 6, 7}
	week := 2
	expectedCountOfBirds := 28

	testBirdsInWeek(t, birds, week, expectedCountOfBirds)
}

func TestBirdsInWeek_WhenSomeWeek(t *testing.T) {
	birds := []int{11, 12, 13, 14, 15, 16, 17, 1, 2, 3, 4, 5, 6, 7, 21, 22, 23, 24, 25, 26, 27}
	week := 2
	expectedCountOfBirds := 28

	testBirdsInWeek(t, birds, week, expectedCountOfBirds)
}

func TestFixBirdCountLog_WhenBirdsArrayIsEmpty(t *testing.T) {
	birds := []int{}
	expectedBirdsPerDay := []int{}

	testFixBirdCountLog(t, birds, expectedBirdsPerDay)
}

func TestFixBirdCountLog_WithOneBird(t *testing.T) {
	birds := []int{1}
	expectedBirdsPerDay := []int{2}

	testFixBirdCountLog(t, birds, expectedBirdsPerDay)
}

func TestFixBirdCountLog_WithSeveralBirds(t *testing.T) {
	birds := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expectedBirdsPerDay := []int{2, 2, 4, 4, 6, 6, 8, 8, 10, 10}

	testFixBirdCountLog(t, birds, expectedBirdsPerDay)
}

func TestFixBirdCountLog_WithALotOfBirds(t *testing.T) {
	birds := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	expectedBirdsPerDay := []int{
		2, 2, 4, 4, 6, 6, 8, 8, 10, 10,
		2, 2, 4, 4, 6, 6, 8, 8, 10, 10,
		2, 2, 4, 4, 6, 6, 8, 8, 10, 10,
		2, 2, 4, 4, 6, 6, 8, 8, 10, 10,
		2, 2, 4, 4, 6, 6, 8, 8, 10, 10,
		2, 2, 4, 4, 6, 6, 8, 8, 10, 10,
		2, 2, 4, 4, 6, 6, 8, 8, 10, 10,
		2, 2, 4, 4, 6, 6, 8, 8, 10, 10,
		2, 2, 4, 4, 6, 6, 8, 8, 10, 10,
		2, 2, 4, 4, 6, 6, 8, 8, 10, 10}

	testFixBirdCountLog(t, birds, expectedBirdsPerDay)
}

func testTotalBirdCount(t *testing.T, birds []int, expectedBirdsCount int) {
	actualBirdsCount := TotalBirdCount(birds)

	if expectedBirdsCount != actualBirdsCount {
		t.Errorf(
			"TotalBirdCount(%v) = %d. Expected %v.",
			birds,
			actualBirdsCount,
			expectedBirdsCount)
	}
}

func testBirdsInWeek(t *testing.T, birdsPerDay []int, week, expectedBirdsCount int) {
	actualBirdsCount := BirdsInWeek(birdsPerDay, week)

	if expectedBirdsCount != actualBirdsCount {
		t.Errorf(
			"BirdsInWeek(%v, %d) = %d. Expected %d.",
			birdsPerDay,
			week,
			actualBirdsCount,
			expectedBirdsCount)
	}
}

func testFixBirdCountLog(t *testing.T, birdsPerDay, expectedBirdsPerDay []int) {
	actualBirdsPerDay := FixBirdCountLog(birdsPerDay)

	if !areEquals(expectedBirdsPerDay, actualBirdsPerDay) {
		t.Errorf(
			"FixBirdCountLog(%v) = %v. Expected %v.",
			birdsPerDay,
			actualBirdsPerDay,
			expectedBirdsPerDay)
	}
}

func areEquals(slice1, slice2 []int) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}

	return true
}
