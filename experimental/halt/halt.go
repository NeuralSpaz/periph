// Copyright 2019 The Periph Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

package halt

import (
	"errors"
	"sync"
	"time"
)

type Blocking interface {
	Deadline() (deadline time.Time, ok bool)
	Done() <-chan struct{}
	Err() error
	Fail(error)
}

type empty int

var background = new(empty)

func (empty) Deadline() (deadline time.Time, ok bool) { return }
func (empty) Done() <-chan struct{}                   { return nil }
func (empty) Err() error                              { return nil }
func (empty) Fail(error)                              {}

func Background() Blocking { return background }

type CancelFn func()

type withCancel struct {
	Blocking
	done chan struct{}

	sync.Mutex
	err error
}

func (wc *withCancel) Done() <-chan struct{} { return wc.done }
func (wc *withCancel) Err() error {
	wc.Lock()
	defer wc.Unlock()
	return wc.err
}
func (wc *withCancel) Fail(err error) {
	wc.Lock()
	defer wc.Unlock()
	if wc.err != nil {
		return
	}
	wc.err = err
	close(wc.done)
}

func WithCancel(parent Blocking) (Blocking, CancelFn) {
	wc := &withCancel{
		Blocking: parent,
		done:     make(chan struct{}),
	}

	c := func() {
		wc.Fail(Canceled)
	}

	go func() {
		select {
		case <-parent.Done():
			wc.Fail(parent.Err())
		case <-wc.Done():
		}
	}()

	return wc, c
}

var (
	Canceled = errors.New("Canceled")
)
