package birdwatcher

func TotalBirdCount(birdsPerDay []int) int {
	totalBirdCount := 0

	for i := 0; i < len(birdsPerDay); i++ {
		totalBirdCount += birdsPerDay[i]
	}

	return totalBirdCount
}

func BirdsInWeek(birdsPerDay []int, week int) int {
	birdsInWeekCount := 0
	startIndex := (week - 1) * 7
	endIndex := startIndex + 7

	for i := startIndex; i < endIndex; i++ {
		birdsInWeekCount += birdsPerDay[i]
	}

	return birdsInWeekCount
}

func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i++ {
		if i%2 == 0 {
			birdsPerDay[i]++
		}
	}

	return birdsPerDay
}
