package stack

// Your implementation of Stack goes here!

type Stack[T any] []T

func (s *Stack[T]) Push(t ...T) {
	*s = append(*s, t...)
}

func (s *Stack[T]) Pop() (T, bool) {
	var elt T
	var ok bool
	if len(*s) != 0 {
		elt = (*s)[len(*s)-1]
		*s = (*s)[:len(*s)-1]
		ok = true
	}
	return elt, ok
}

func (s *Stack[T]) Len() int {
	return len(*s)
}
