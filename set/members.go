package set

func (s *set[T]) Members() []T {
	if !s.modified {
		return s.members
	}

	s.members = make([]T, 0, len(s.data))
	for member := range s.data {
		s.members = append(s.members, member)
	}
	s.modified = false
	return s.members
}
