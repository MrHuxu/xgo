package heap

func (h *heap[T]) Size() int {
	return len(h.data)
}
