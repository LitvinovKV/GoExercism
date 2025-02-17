package cards

func FavoriteCards() []int {
	return []int{2, 6, 9}
}

func GetItem(slice []int, index int) int {
	if index >= 0 && index < len(slice) {
		return slice[index]
	}

	return -1
}

func SetItem(slice []int, index, value int) []int {
	if GetItem(slice, index) == -1 {
		return append(slice, value)
	}

	slice[index] = value
	return slice
}

func PrependItems(slice []int, values ...int) []int {
	return append(values, slice...)
}

func RemoveItem(slice []int, index int) []int {
	if GetItem(slice, index) == -1 {
		return slice
	}

	return append(slice[:index], slice[index+1:]...)
}
