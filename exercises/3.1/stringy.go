package stringy

import (
	"fmt"
	"io"
)

// Your StringifyTo function goes here!
func StringifyTo[T fmt.Stringer](buf io.Writer, v T) {
	fmt.Fprintln(buf, v)
}
