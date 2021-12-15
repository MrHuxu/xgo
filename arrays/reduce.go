package arrays

// Reduce reduces array to a value which is the accumulated result of running each element in array thru iteratee
func Reduce[T any](array []T, iteratee func(idx int, value, accumulator T) T, initAccumulator T) (result T) {
	result = initAccumulator

	for idx, value := range array {
		result = iteratee(idx, value, result)
	}

	return
}
