// Copyright 2019 The Periph Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

package halt

import (
	"errors"
	"testing"
	"time"
)

func TestWithCancel(t *testing.T) {
	block, cancel := WithCancel(Background())

	if err := block.Err(); err != nil {
		t.Errorf("should be no error yet, got %v", err)
	}
	cancel()
	if err := block.Err(); err != Canceled {
		t.Errorf("err should be: %v,  but got: %v ", Canceled, err)
	}
	// Calling cancel more than once should not be a problem.
	cancel()
	cancel()
	cancel()
	cancel()
	if err := block.Err(); err != Canceled {
		t.Errorf("err should be: %v,  but got: %v ", Canceled, err)
	}
}

var failingError = errors.New("a failing error")

func TestWithCancel_Fail(t *testing.T) {
	block, cancel := WithCancel(Background())

	if err := block.Err(); err != nil {
		t.Errorf("should be no error yet, got %v", err)
	}
	block.Fail(failingError)
	if err := block.Err(); err != failingError {
		t.Errorf("err should be: %v,  but got: %v ", failingError, err)
	}
	// Calling cancel after Fail should not be a problem.
	cancel()
	if err := block.Err(); err != failingError {
		t.Errorf("err should be: %v,  but got: %v ", failingError, err)
	}
}

func TestWithCancelChild(t *testing.T) {
	parent, cancelParent := WithCancel(Background())
	child, cancelChild := WithCancel(parent)
	defer cancelParent()
	cancelChild()

	if err := parent.Err(); err != nil {
		t.Errorf("parent err should be: %v,  but got: %v ", nil, err)
	}
	if err := child.Err(); err != Canceled {
		t.Errorf("child err should be: %v,  but got: %v ", Canceled, err)
	}

}

func TestWithCancelParent(t *testing.T) {
	parent, cancelParent := WithCancel(Background())
	child, cancelChild := WithCancel(parent)
	defer cancelChild()
	cancelParent()

	if err := parent.Err(); err != Canceled {
		t.Errorf("parent err should be: %v,  but got: %v ", Canceled, err)
	}
	// TODO: fix propagation delay
	time.Sleep(time.Millisecond)
	if err := child.Err(); err != Canceled {
		t.Errorf("child err should be: %v,  but got: %v ", Canceled, err)
	}

}
