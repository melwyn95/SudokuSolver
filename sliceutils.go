package main

func RemoveFromSlice(slice []uint8, index int) []uint8 {
	lastIndex := len(slice) - 1
	slice[index] = slice[lastIndex]
	slice[lastIndex] = 0
	return slice[:lastIndex]
}

func findSlice(slice []uint8, predicate func(uint8) bool) (int, uint8) {
	for index, value := range slice {
		if predicate(value) {
			return index, value
		}
	}
	return -1, 0
}

func FindIndexSlice(slice []uint8, value uint8) int {
	index, _ := findSlice(slice, func(item uint8) bool {
		return item == value
	})
	return index
}

func ExistsSlice(slice []uint8, value uint8) bool {
	return FindIndexSlice(slice, value) != -1
}
