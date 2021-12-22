package array

// Uniq remove duplicate elements from an array
func Uniq[T comparable](array []T) []T {
	tmp := make(map[T]struct{}, len(array))
	var uniqIdx int
	for _, elem := range array {
		if _, ok := tmp[elem]; !ok {
			array[uniqIdx] = elem
			uniqIdx++

			tmp[elem] = struct{}{}
		}
	}

	return array[:uniqIdx]
}

// UniqBy is like Uniq except that it accepts iteratee which is invoked for each element in array to generate the criterion by which uniqueness is computed
func UniqBy[T any, C comparable](array []T, iteratee func(T) C) []T {
	tmp := make(map[C]struct{}, len(array))
	var uniqIdx int
	for _, elem := range array {
		val := iteratee(elem)
		if _, ok := tmp[val]; !ok {
			array[uniqIdx] = elem
			uniqIdx++

			tmp[val] = struct{}{}
		}
	}

	return array[:uniqIdx]
}
