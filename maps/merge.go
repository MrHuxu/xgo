package maps

// Merge copies all key value pairs of maps to the base map
func Merge[K comparable, V any](base map[K]V, maps ...map[K]V) (result map[K]V) {
	if len(maps) == 0 {
		return base
	}

	result = base
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}

	return
}
