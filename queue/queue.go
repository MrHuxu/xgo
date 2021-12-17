package queue

type Queue[T any] interface {
	Size() int
	Enqueue(T) bool
	Dequeue() (T, bool)
}

type queue[T any] struct {
	data []T
}

func NewQueue[T any]() Queue[T] {
	return &queue[T]{}
}
