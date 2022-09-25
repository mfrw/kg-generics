package print

import (
	"fmt"
	"io"
)

// Your PrintAnythingTo function goes here!

func PrintAnythingTo[T any](buf io.Writer, val T) {
	buf.Write([]byte(fmt.Sprintln(val)))
}
