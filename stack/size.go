package stack

func (s *stack[T]) Size() int {
	if s == nil {
		return 0
	}

	return len(s.data)
}
