package arrays

// First iterates over elements of array, returning the first element predicate returns truthy for
func First[T any](array []T, predicate func(idx int, value T) bool) (result T) {
	for idx := range array {
		if predicate(idx, array[idx]) {
			return array[idx]
		}
	}

	return
}
