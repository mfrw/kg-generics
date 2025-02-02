package intish

// Your Intish interface goes here!
type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Intish interface {
	Signed | Unsigned
}

func IsPositive[T Intish](v T) bool {
	return v > 0
}
