package cars

const CountOfCarsInGroup = 10
const PriceOfIndividuallyCar = 10000
const PriceOfCarGroup = 95000

func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * (successRate / float64(100))
}

func CalcualteWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(CalculateWorkingCarsPerHour(productionRate, successRate)) / 60
}

func CalculateCost(carsCount int) uint {
	return uint(((carsCount / CountOfCarsInGroup) * PriceOfCarGroup) + ((carsCount % CountOfCarsInGroup) * PriceOfIndividuallyCar))
}
