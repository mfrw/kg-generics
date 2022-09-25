package contains

import "golang.org/x/exp/slices"

// Your implementation of ContainsFunc goes here!

func ContainsFunc[T any](s []T, fn func(T) bool) bool {
	return slices.IndexFunc(s, fn) >= 0
}
