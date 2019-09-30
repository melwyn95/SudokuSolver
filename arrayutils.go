package main

func findArray(array *[9]uint8, predicate func(uint8) bool) (int, uint8) {
	for index, value := range array {
		if predicate(value) {
			return index, value
		}
	}
	return -1, 0
}

func FindIndexArray(array *[9]uint8, value uint8) int {
	index, _ := findArray(array, func(item uint8) bool {
		return item == value
	})
	return index
}

func ExistsArray(array *[9]uint8, value uint8) bool {
	return FindIndexArray(array, value) != -1
}
