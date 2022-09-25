package compose

// Your implementation of Compose goes here!

func Compose[I any, O1 any, O2 any](outer func(i O1) O2, inner func(i I) O1, i I) O2 {
	return outer(inner(i))
}
