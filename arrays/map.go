package arrays

// Map creates an array of values by running each element of array thru iteratee
func Map[T any](array []T, iteratee func(idx int, value T) T) (result []T) {
	result = make([]T, 0, len(array))

	length := len(array)
	for idx := 0; idx < length; idx++ {
		result = append(result, iteratee(idx, array[idx]))
	}

	return
}
