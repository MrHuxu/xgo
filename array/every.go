package array

// Every checks if predicate returns truthy for **all** elements of array
func Every[T any](array []T, predicate func(value T) bool) bool {
	for i := range array {
		if !predicate(array[i]) {
			return false
		}
	}

	return true
}
