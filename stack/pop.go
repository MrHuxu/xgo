package stack

func (s *stack[T]) Pop() (val T, ok bool) {
	if s == nil || len(s.data) == 0 {
		return
	}

	length := len(s.data)
	val = s.data[length-1]
	s.data = s.data[0 : length-1]
	ok = true
	return
}
