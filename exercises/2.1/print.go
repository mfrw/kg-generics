package print

import (
	"fmt"
	"io"
)

// Your PrintAnythingTo function goes here!

func PrintAnythingTo[T any](buf io.Writer, val T) {
	fmt.Fprintln(buf, val)
}
