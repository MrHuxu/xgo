package arrays

// Filter iterates over elements of array, returning an array of all elements predicate returns truthy for
func Filter[T any](array []T, predicate func(idx int, value T) bool) (result []T) {
	result = make([]T, 0, len(array))

	for idx, value := range array {
		if predicate(idx, value) {
			result = append(result, value)
		}
	}

	return
}
