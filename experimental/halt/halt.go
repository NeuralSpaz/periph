// Copyright 2019 The Periph Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

package halt

import (
	"errors"
	"sync"
)

type Blocking interface {
	Done() <-chan struct{}
	Err() error
	Fail(error)
}

type empty int

var background = new(empty)

func (empty) Done() <-chan struct{} { return nil }
func (empty) Err() error            { return nil }
func (empty) Fail(error)            {}

func Background() Blocking { return background }

type CancelFn func()

type cancelable struct {
	Blocking
	done chan struct{}

	sync.Mutex
	err error
}

func (c *cancelable) Done() <-chan struct{} { return c.done }
func (c *cancelable) Err() error {
	c.Lock()
	defer c.Unlock()
	return c.err
}
func (c *cancelable) Fail(err error) {
	c.Lock()
	defer c.Unlock()
	if c.err != nil {
		return
	}
	c.err = err
	close(c.done)
}

func WithCancel(parent Blocking) (Blocking, CancelFn) {
	c := &cancelable{
		Blocking: parent,
		done:     make(chan struct{}),
	}

	cFn := func() {
		c.Fail(Canceled)
	}

	go func() {
		select {
		case <-parent.Done():
			c.Fail(parent.Err())
		case <-c.Done():
		}
	}()

	return c, cFn
}

var (
	Canceled = errors.New("Canceled")
)
