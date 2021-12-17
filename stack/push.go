package stack

func (s *stack[T]) Push(val T) (ok bool) {
	if s == nil {
		return
	}

	s.data = append(s.data, val)
	ok = true
	return
}
