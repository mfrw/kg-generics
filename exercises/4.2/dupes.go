package dupes

// Your Dupes function goes here!

func Dupes[T comparable](s []T) bool {
	mp := make(map[T]struct{})
	for _, v := range s {
		if _, ok := mp[v]; ok {
			return true
		}
		mp[v] = struct{}{}
	}
	return false
}
