package set

func (s *set[T]) Remove(member T) {
	delete(s.data, member)
	s.modified = true
}
