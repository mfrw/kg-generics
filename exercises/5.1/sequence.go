package sequence

// Your implementation of Sequence goes here!

type Sequence[T any] []T

func (s *Sequence[T]) Empty() bool {
	return len(*s) == 0
}
