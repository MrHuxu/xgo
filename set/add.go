package set

func (s *set[T]) Add(member T) {
	s.data[member] = struct{}{}
	s.modified = true
}
