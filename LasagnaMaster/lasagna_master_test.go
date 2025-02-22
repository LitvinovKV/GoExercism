package lasagna

import (
	"math"
	"testing"
)

func TestPreparationTime_WhenLayersArrayIsEmpty(t *testing.T) {
	layers := []string{}
	timeForOneLayerPreparation := 7
	expectedPreparationTime := 0

	testPreparationTime(t, layers, timeForOneLayerPreparation, expectedPreparationTime)
}

func TestPreparationTime_WhenLayersArrayIsNil(t *testing.T) {
	var layers []string
	timeForOneLayerPreparation := 7
	expectedPreparationTime := 0

	testPreparationTime(t, layers, timeForOneLayerPreparation, expectedPreparationTime)
}

func TestPreparationTime_WhenLayersContainsOneElement(t *testing.T) {
	layers := []string{"some layer"}
	timeForOneLayerPreparation := 7
	expectedPreparationTime := 7

	testPreparationTime(t, layers, timeForOneLayerPreparation, expectedPreparationTime)
}

func TestPreparationTime_WhenLayersContainsSeveralElements(t *testing.T) {
	layers := []string{"some layer", "some layer", "some layer", "some layer", "some layer"}
	timeForOneLayerPreparation := 3
	expectedPreparationTime := 15

	testPreparationTime(t, layers, timeForOneLayerPreparation, expectedPreparationTime)
}

func TestPreparationTime_WhenLayersContainsManyElements(t *testing.T) {
	layers := make([]string, 0, 100)

	for i := 0; i < 100; i++ {
		layers = append(layers, "some layer")
	}

	timeForOneLayerPreparation := 15
	expectedPreparationTime := 1500

	testPreparationTime(t, layers, timeForOneLayerPreparation, expectedPreparationTime)
}

func TestPreparationTime_WhenLayersTimeForOnePreparationIsDefault(t *testing.T) {
	layers := []string{"some layer", "some layer", "some layer"}
	expectedPreparationTime := 6

	testPreparationTime(t, layers, 0, expectedPreparationTime)
}

func TestQuantities_WhenLayersArrayIsEmpty(t *testing.T) {
	layers := []string{}
	expectedAmountOfNoodles := 0
	expectedAmountOfSauce := 0.

	testQuantities(t, layers, expectedAmountOfNoodles, expectedAmountOfSauce)
}

func TestQuantities_WhenLayersArrayIsNil(t *testing.T) {
	var layers []string
	expectedAmountOfNoodles := 0
	expectedAmountOfSauce := 0.

	testQuantities(t, layers, expectedAmountOfNoodles, expectedAmountOfSauce)
}

func TestQuantities_WithOneNoodles(t *testing.T) {
	layers := []string{"noodles"}
	expectedAmountOfNoodles := 50
	expectedAmountOfSauce := 0.

	testQuantities(t, layers, expectedAmountOfNoodles, expectedAmountOfSauce)
}

func TestQuantities_WithOneSauce(t *testing.T) {
	layers := []string{"sauce"}
	expectedAmountOfNoodles := 0
	expectedAmountOfSauce := 0.2

	testQuantities(t, layers, expectedAmountOfNoodles, expectedAmountOfSauce)
}

func TestQuantities_WithSeveralIngredients(t *testing.T) {
	layers := []string{"sauce", "noodles", "sauce", "noodles", "sauce", "noodles"}
	expectedAmountOfNoodles := 150
	expectedAmountOfSauce := 0.6

	testQuantities(t, layers, expectedAmountOfNoodles, expectedAmountOfSauce)
}

func TestQuantities_WithManyIngredients(t *testing.T) {

	layers := make([]string, 0, 100)

	for i := 0; i < 50; i++ {
		layers = append(layers, "sauce")
		layers = append(layers, "noodles")
	}

	expectedAmountOfNoodles := 2500
	expectedAmountOfSauce := 10.

	testQuantities(t, layers, expectedAmountOfNoodles, expectedAmountOfSauce)
}

func TestAddSecretIngredient_WithOneElementsInBothCollections(t *testing.T) {
	friendsIngredients := []string{"secret ingredient"}
	myIngredients := []string{"my ingredient"}
	expectedMyIngredients := []string{"secret ingredient"}

	testAddSecretIngredient(t, friendsIngredients, myIngredients, expectedMyIngredients)
}

func TestAddSecretIngredient_WhenFriendHasMoreIngredients(t *testing.T) {
	friendsIngredients := []string{"ingredient1", "ingredient1", "secret ingredient"}
	myIngredients := []string{"my ingredient"}
	expectedMyIngredients := []string{"secret ingredient"}

	testAddSecretIngredient(t, friendsIngredients, myIngredients, expectedMyIngredients)
}

func TestAddSecretIngredient_WhenIHaveMoreIngredients(t *testing.T) {
	friendsIngredients := []string{"secret ingredient"}
	myIngredients := []string{"my ingredient1", "my ingredient2", "my ingredient3", "my ingredient4"}
	expectedMyIngredients := []string{"my ingredient1", "my ingredient2", "my ingredient3", "secret ingredient"}

	testAddSecretIngredient(t, friendsIngredients, myIngredients, expectedMyIngredients)
}

func TestAddSecretIngredient_WhenBothHaveManyIngredients(t *testing.T) {
	friendsIngredients := make([]string, 0, 101)
	myIngredients := make([]string, 0, 100)

	for i := 0; i < 100; i++ {
		friendsIngredients = append(friendsIngredients, "friend ingredient")
		myIngredients = append(myIngredients, "my ingredient")
	}

	friendsIngredients = append(friendsIngredients, "secret ingredient")

	expectedMyIngredients := make([]string, len(myIngredients))
	copy(expectedMyIngredients, myIngredients)
	expectedMyIngredients[len(expectedMyIngredients)-1] = "secret ingredient"

	testAddSecretIngredient(t, friendsIngredients, myIngredients, expectedMyIngredients)
}

func TestScaleRecipe_WhenAmountsArrayIsEmpty(t *testing.T) {
	amountsForTwoPortions := []float64{}
	expectedAmountForPortions := []float64{}

	testScaleRecipe(t, amountsForTwoPortions, 2, expectedAmountForPortions)
}

func TestScaleRecipe_WhenAmountsArrayIsNil(t *testing.T) {
	var amountsForTwoPortions []float64
	expectedAmountForPortions := []float64{}

	testScaleRecipe(t, amountsForTwoPortions, 2, expectedAmountForPortions)
}

func TestScaleRecipe_WhenForOnePerson(t *testing.T) {
	amountsForTwoPortions := []float64{1, 2, 3, 4, 5}
	numberOfPortions := 1
	expectedAmountForPortions := []float64{0.5, 1, 1.5, 2, 2.5}

	testScaleRecipe(t, amountsForTwoPortions, numberOfPortions, expectedAmountForPortions)
}

func TestScaleRecipe_WhenForEvenNumberOfPersons(t *testing.T) {
	amountsForTwoPortions := []float64{1, 2, 3, 4, 5}
	numberOfPortions := 6
	expectedAmountForPortions := []float64{3, 6, 9, 12, 15}

	testScaleRecipe(t, amountsForTwoPortions, numberOfPortions, expectedAmountForPortions)
}

func TestScaleRecipe_WhenForOddNumberOfPersons(t *testing.T) {
	amountsForTwoPortions := []float64{1, 2, 3, 4, 5}
	numberOfPortions := 5
	expectedAmountForPortions := []float64{2.5, 5, 7.5, 10, 12.5}

	testScaleRecipe(t, amountsForTwoPortions, numberOfPortions, expectedAmountForPortions)
}

func testPreparationTime(t *testing.T, layers []string, timeForOneLayerPreparation, expectedPreparationTime int) {
	actualPreparationTime := PreparationTime(layers, timeForOneLayerPreparation)

	if expectedPreparationTime != actualPreparationTime {
		t.Errorf(
			"PreparationTime(%v, %d) = %d. Expected %d.",
			layers,
			timeForOneLayerPreparation,
			actualPreparationTime,
			expectedPreparationTime)
	}
}

func testQuantities(t *testing.T, layers []string, expectedAmountOfNoodles int, expectedAmountOfSauce float64) {
	actualAmountOfNoodles, actualAmountOfSauce := Quantities(layers)

	roundedActualAmountOfSauce := math.Round(actualAmountOfSauce*100) / 100
	roundedExpectedAmountOfSauce := math.Round(expectedAmountOfSauce*100) / 100

	if expectedAmountOfNoodles != actualAmountOfNoodles ||
		roundedExpectedAmountOfSauce != roundedActualAmountOfSauce {
		t.Errorf(
			"Quantities(%v) = (%d, %f). Expected (%d, %f).",
			layers,
			actualAmountOfNoodles,
			actualAmountOfSauce,
			expectedAmountOfNoodles,
			expectedAmountOfSauce)
	}
}

func testAddSecretIngredient(t *testing.T, friendsIngredients, myIngredients, expectedIngredients []string) {
	baseMyIngredients := make([]string, len(myIngredients))
	copy(baseMyIngredients, myIngredients)

	AddSecretIngredient(friendsIngredients, myIngredients)

	if !areEquals(expectedIngredients, myIngredients) {
		t.Errorf(
			"AddSecretIngredient(%v, %v) = %v. Expected %v.",
			friendsIngredients,
			baseMyIngredients,
			myIngredients,
			expectedIngredients)
	}
}

func testScaleRecipe(t *testing.T, amountsForTwoPortions []float64, numberOfPortions int, expectedAmountForPortions []float64) {
	actualAmountForPortions := ScaleRecipe(amountsForTwoPortions, numberOfPortions)

	if !areEquals(expectedAmountForPortions, actualAmountForPortions) {
		t.Errorf(
			"ScaleRecipe(%v, %d) = %v. Expected %v.",
			amountsForTwoPortions,
			numberOfPortions,
			actualAmountForPortions,
			expectedAmountForPortions)
	}
}

func areEquals[T int | float64 | string](slice1, slice2 []T) bool {
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
