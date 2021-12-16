package maps

// Merge copies all key value pairs of maps to the base map
func Merge[K comparable, V any](base map[K]V, maps ...map[K]V) (result map[K]V) {
	if len(maps) == 0 {
		return base
	}

	result = base
	length := len(maps)
	for idx := 0; idx < length; idx++ {
		for k, v := range maps[idx] {
			result[k] = v
		}
	}

	return
}
