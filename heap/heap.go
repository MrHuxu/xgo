package heap

type Heap[T any] interface {
	Size() int
	Push(T)
	Pop() (T, bool)
}

type heap[T any] struct {
	data    []T
	compare compareFunc[T]
}

type compareFunc[T any] func(T, T) bool

func NewHeap[T any](compareFunc compareFunc[T]) Heap[T] {
	return &heap[T]{
		compare: compareFunc,
	}
}
