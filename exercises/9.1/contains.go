package contains

// Your implementation of ContainsFunc goes here!

func ContainsFunc[T any](s []T, fn func(T) bool) bool {
	for _, v := range s {
		if fn(v) {
			return true
		}
	}
	return false
}
