package queue

func (q *queue[T]) Enqueue(val T) (ok bool) {
	if q == nil {
		return
	}

	q.data = append(q.data, val)
	ok = true
	return
}
