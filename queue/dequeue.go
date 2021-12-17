package queue

func (q *queue[T]) Dequeue() (val T, ok bool) {
	if q == nil || len(q.data) == 0 {
		return
	}

	val = q.data[0]
	q.data = q.data[1:]
	ok = true
	return
}
