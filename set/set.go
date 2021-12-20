package set

type Set[T comparable] interface {
	Size() int
	Add(T)
	Remove(T)
	Members() []T
	Each(func(T) bool)
}

type set[T comparable] struct {
	data     map[T]struct{}
	modified bool
	members  []T
}

func NewSet[T comparable]() Set[T] {
	return &set[T]{
		data: make(map[T]struct{}),
	}
}
