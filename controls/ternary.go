package controls

// Ternary returns valueIfTrue if the result of condition is truthy, else returns the valueIfFalse
func Ternary[T any](condition bool, valueIfTrue, valueIfFalse T) T {
	if condition {
		return valueIfTrue
	}
	return valueIfFalse
}
