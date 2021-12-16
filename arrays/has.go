package arrays

// Has checks if any element of the array that predicate returns truthy for
func Has[T any](array []T, predicate func(idx int, value T) bool) bool {
	length := len(array)
	for idx := 0; idx < length; idx++ {
		if predicate(idx, array[idx]) {
			return true
		}
	}
	return false
}

// HasElem checks if any element of the array equals to the value
func HasElem[T comparable](array []T, value T) bool {
	length := len(array)
	for idx := 0; idx < length; idx++ {
		if array[idx] == value {
			return true
		}
	}
	return false
}
