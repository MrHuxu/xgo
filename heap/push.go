package heap

func (h *heap[T]) Push(value T) {
	h.data = append(h.data, value)
	h.buildHeap()
}

func (h *heap[T]) buildHeap() {
	var left, right int
	length := len(h.data)
	for mid := length/2 - 1; mid >= 0; mid-- {
		left = mid*2 + 1
		right = (mid + 1) * 2

		if right < length && h.compare(h.data[right], h.data[mid]) {
			h.data[mid], h.data[right] = h.data[right], h.data[mid]
		}

		if h.compare(h.data[left], h.data[mid]) {
			h.data[mid], h.data[left] = h.data[left], h.data[mid]
		}
	}
}
