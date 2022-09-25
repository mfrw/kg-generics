package merge

// Your implementation of Merge goes here!

func Merge[M ~map[K]V, K comparable, V any](mps ...M) M {
	mp := map[K]V{}

	for _, m := range mps {
		for k, v := range m {
			mp[k] = v
		}
	}
	return mp
}
