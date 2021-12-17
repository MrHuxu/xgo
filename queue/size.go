package queue

func (q *queue[T]) Size() int {
	if q == nil {
		return 0
	}

	return len(q.data)
}
