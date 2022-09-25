package billable

import "sync/atomic"

// Your implementation of Channel goes here!!

type Channel[T any] struct {
	ch chan T
	tx atomic.Int32
	rx atomic.Int32
}

func NewChannel[T any](capacity int) *Channel[T] {
	return &Channel[T]{
		ch: make(chan T, capacity),
	}
}

func (c *Channel[T]) Send(e T) {
	defer func() { c.tx.Add(1) }()
	c.ch <- e
}

func (c *Channel[T]) Receive() T {
	defer func() { c.rx.Add(1) }()
	return <-c.ch
}

func (c *Channel[T]) Sends() int {
	return int(c.tx.Load())
}

func (c *Channel[T]) Receives() int {
	return int(c.rx.Load())
}
