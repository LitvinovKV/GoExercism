package lasagna

const defaultTimeForOneLayerPreparation = 2
const amountOfNoodlesForOneLayer = 50
const amountOfSauceForOneLayer = 0.2
const defaultNumberOfPortions = 2.

func PreparationTime(layers []string, timeForOneLayerPreparation int) int {
	if timeForOneLayerPreparation == 0 {
		timeForOneLayerPreparation = defaultTimeForOneLayerPreparation
	}
	return len(layers) * timeForOneLayerPreparation
}

func Quantities(layers []string) (amountOfNoodles int, amountOfSauce float64) {
	amountOfNoodles = 0
	amountOfSauce = 0.

	for i := 0; i < len(layers); i++ {
		if layers[i] == "sauce" {
			amountOfSauce += amountOfSauceForOneLayer
		}
		if layers[i] == "noodles" {
			amountOfNoodles += amountOfNoodlesForOneLayer
		}
	}

	return
}

func AddSecretIngredient(friendsIngredients, myIngredients []string) {
	myIngredients[len(myIngredients)-1] = friendsIngredients[len(friendsIngredients)-1]
}

func ScaleRecipe(amountsForTwoPortions []float64, numberOfPortions int) []float64 {
	scaledAmounts := make([]float64, len(amountsForTwoPortions))
	scalingFactor := float64(numberOfPortions) / defaultNumberOfPortions

	for i := 0; i < len(amountsForTwoPortions); i++ {
		scaledAmounts[i] = amountsForTwoPortions[i] * scalingFactor
	}

	return scaledAmounts
}
