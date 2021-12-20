package set

func (s *set[T]) Size() int {
	return len(s.data)
}
