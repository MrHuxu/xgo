package arrays

// Filter iterates over elements of array, returning an array of all elements predicate returns truthy for
func Filter[T any](array []T, predicate func(idx int, value T) bool) (result []T) {
	result = make([]T, 0, len(array))

	length := len(array)
	for idx := 0; idx < length; idx++ {
		if predicate(idx, array[idx]) {
			result = append(result, array[idx])
		}
	}

	return
}
