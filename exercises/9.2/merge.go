package merge

import "golang.org/x/exp/maps"

// Your implementation of Merge goes here!

func Merge[M ~map[K]V, K comparable, V any](mps ...M) M {
	mp := M{}
	for _, m := range mps {
		maps.Copy[map[K]V](mp, m)
	}
	return mp
}
