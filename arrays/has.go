package arrays

// Has checks if any element of the array that predicate returns truthy for
func Has[T any](array []T, predicate func(idx int, value T) bool) bool {
	for idx, value := range array {
		if predicate(idx, value) {
			return true
		}
	}
	return false
}

// HasElem checks if any element of the array equals to the value
func HasElem[T comparable](array []T, value T) bool {
	for _, elem := range array {
		if elem == value {
			return true
		}
	}
	return false
}
