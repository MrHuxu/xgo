package heap

func (h *heap[T]) Pop() (value T, ok bool) {
	if len(h.data) == 0 {
		return
	}

	value = h.data[0]
	ok = true

	h.data = h.data[1:]
	h.buildHeap()
	return
}
