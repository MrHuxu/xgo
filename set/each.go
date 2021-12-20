package set

func (s *set[T]) Each(iteratee func(T) bool) {
	for member := range s.data {
		if !iteratee(member) {
			break
		}
	}
}
