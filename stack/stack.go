package stack

type Stack[T any] interface {
	Size() int
	Push(T) bool
	Pop() (T, bool)
}

type stack[T any] struct {
	data []T
}

func NewStack[T any]() Stack[T] {
	return &stack[T]{}
}
