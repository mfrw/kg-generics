package group

type Group[E any] []E

// Your Len function goes here!

func Len[E any](s []E) int {
	return len(s)
}
