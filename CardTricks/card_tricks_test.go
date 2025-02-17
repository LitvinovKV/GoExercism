package cards

import "testing"

func TestFavoriteCards(t *testing.T) {
	expectedFavoriteCards := []int{2, 6, 9}
	actualFavoriteCards := FavoriteCards()

	if !areEquals(expectedFavoriteCards, actualFavoriteCards) {
		t.Errorf(
			"TestFavoriteCards() = %v. Expected %v",
			actualFavoriteCards,
			expectedFavoriteCards)
	}
}

func TestGetItem_WhenIndexIsNegative(t *testing.T) {
	slice := []int{1, 2, 3}
	index := -1
	expectedValue := -1

	testGetItem(t, slice, index, expectedValue)
}

func TestGetItem_WhenIndexIsOutOfRange(t *testing.T) {
	slice := []int{1, 2, 3}
	index := 4
	expectedValue := -1

	testGetItem(t, slice, index, expectedValue)
}

func TestGetItem_GetFirstItem(t *testing.T) {
	slice := []int{1, 2, 3}
	index := 0
	expectedValue := 1

	testGetItem(t, slice, index, expectedValue)
}

func TestGetItem_GetLastItem(t *testing.T) {
	slice := []int{1, 2, 3}
	index := 2
	expectedValue := 3

	testGetItem(t, slice, index, expectedValue)
}

func TestGetItem_GetSomeItem(t *testing.T) {
	slice := []int{1, 2, 3}
	index := 1
	expectedValue := 2

	testGetItem(t, slice, index, expectedValue)
}

func TestSetItem_WhenIndexIsNegative(t *testing.T) {
	slice := []int{1, 2, 3}
	index := -1
	value := 5
	expectedSlice := []int{1, 2, 3, 5}

	testSetItem(t, slice, index, value, expectedSlice)
}

func TestSetItem_WhenIndexIsOutOfRange(t *testing.T) {
	slice := []int{1, 2, 3}
	index := 4
	value := 5
	expectedSlice := []int{1, 2, 3, 5}

	testSetItem(t, slice, index, value, expectedSlice)
}

func TestSetItem_OverwriteFirstElement(t *testing.T) {
	slice := []int{1, 2, 3}
	index := 0
	value := 5
	expectedSlice := []int{5, 2, 3}

	testSetItem(t, slice, index, value, expectedSlice)
}

func TestSetItem_OverwriteLastElement(t *testing.T) {
	slice := []int{1, 2, 3}
	index := 2
	value := 5
	expectedSlice := []int{1, 2, 5}

	testSetItem(t, slice, index, value, expectedSlice)
}

func TestSetItem_OverwriteSomeElement(t *testing.T) {
	slice := []int{1, 2, 3}
	index := 1
	value := 5
	expectedSlice := []int{1, 5, 3}

	testSetItem(t, slice, index, value, expectedSlice)
}

func TestPrependItems_WhenFirstSliceIsEmpty(t *testing.T) {
	slice1 := []int{}
	slice2 := []int{1, 2, 3}
	expectedSlice := []int{1, 2, 3}

	testPrependItems(t, slice1, slice2, expectedSlice)
}

func TestPrependItems_WhenSecondSliceIsEmpty(t *testing.T) {
	slice1 := []int{1, 2, 3}
	slice2 := []int{}
	expectedSlice := []int{1, 2, 3}

	testPrependItems(t, slice1, slice2, expectedSlice)
}

func TestPrependItems_WhenFirstSliceIsNil(t *testing.T) {
	var slice1 []int = nil
	slice2 := []int{1, 2, 3}
	expectedSlice := []int{1, 2, 3}

	testPrependItems(t, slice1, slice2, expectedSlice)
}

func TestPrependItems_WhenSecondSliceIsNil(t *testing.T) {
	slice1 := []int{1, 2, 3}
	var slice2 []int = nil
	expectedSlice := []int{1, 2, 3}

	testPrependItems(t, slice1, slice2, expectedSlice)
}

func TestPrependItems_WhenBothSlicesAreEmpty(t *testing.T) {
	slice1 := []int{}
	slice2 := []int{}
	expectedSlice := []int{}

	testPrependItems(t, slice1, slice2, expectedSlice)
}

func TestPrependItems_WhenBothSlicesAreNil(t *testing.T) {
	var slice1 []int = nil
	var slice2 []int = nil
	expectedSlice := []int{}

	testPrependItems(t, slice1, slice2, expectedSlice)
}

func TestPrependItems_WhenBothSlicesWithSomeValues(t *testing.T) {
	slice1 := []int{1, 2, 3}
	slice2 := []int{4, 5, 6}
	expectedSlice := []int{4, 5, 6, 1, 2, 3}

	testPrependItems(t, slice1, slice2, expectedSlice)
}

func TestRemoveItem_WhenIndexIsNegative(t *testing.T) {
	slice := []int{1, 2, 3}
	index := -1
	expectedSlice := []int{1, 2, 3}

	testRemoveItem(t, slice, index, expectedSlice)
}

func TestRemoveItem_WhenIndexIsOutOfRange(t *testing.T) {
	slice := []int{1, 2, 3}
	index := 4
	expectedSlice := []int{1, 2, 3}

	testRemoveItem(t, slice, index, expectedSlice)
}

func TestRemoveItem_RemoveFirstElement(t *testing.T) {
	slice := []int{1, 2, 3}
	index := 0
	expectedSlice := []int{2, 3}

	testRemoveItem(t, slice, index, expectedSlice)
}

func TestRemoveItem_RemoveLastElement(t *testing.T) {
	slice := []int{1, 2, 3}
	index := 2
	expectedSlice := []int{1, 2}

	testRemoveItem(t, slice, index, expectedSlice)
}

func TestRemoveItem_RemoveSomeElement(t *testing.T) {
	slice := []int{1, 2, 3}
	index := 1
	expectedSlice := []int{1, 3}

	testRemoveItem(t, slice, index, expectedSlice)
}

func testGetItem(t *testing.T, slice []int, index, expectedValue int) {
	actualValue := GetItem(slice, index)

	if expectedValue != actualValue {
		t.Errorf(
			"GetItem(%v, %d) = %d. Expected %d",
			slice,
			index,
			actualValue,
			expectedValue)
	}
}

func testSetItem(t *testing.T, slice []int, index, value int, expectedSlice []int) {
	actualSlice := SetItem(slice, index, value)

	if !areEquals(expectedSlice, actualSlice) {
		t.Errorf(
			"SetItem(%v, %d, %d) = %v. Expected %v",
			slice,
			index,
			value,
			actualSlice,
			expectedSlice)
	}
}

func testPrependItems(t *testing.T, slice1, slice2, expectedSlice []int) {
	actualSlice := PrependItems(slice1, slice2...)

	if !areEquals(expectedSlice, actualSlice) {
		t.Errorf(
			"PrependItems(%v, %v) = %v. Expected %v",
			slice1,
			slice2,
			actualSlice,
			expectedSlice)
	}
}

func testRemoveItem(t *testing.T, slice []int, index int, expectedSlice []int) {
	actualSlice := RemoveItem(slice, index)

	if !areEquals(expectedSlice, actualSlice) {
		t.Errorf(
			"RemoveItem(%v, %d) = %v. Expected = %v",
			slice,
			index,
			actualSlice,
			expectedSlice)
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
