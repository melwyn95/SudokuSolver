package arrayutils

func find(array *[9]uint8, predicate func(uint8) bool) (int, uint8) {
	for index, value := range array {
		if predicate(value) {
			return index, value
		}
	}
	return -1, 0
}

func FindIndex(array *[9]uint8, value uint8) int {
	index, _ := find(array, func(item uint8) bool {
		return item == value
	})
	return index
}

func Exists(array *[9]uint8, value uint8) bool {
	return FindIndex(array, value) != -1
}
