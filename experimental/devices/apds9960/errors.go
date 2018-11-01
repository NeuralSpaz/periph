package apds9960

import "errors"

// IOError is a I/O specific error.
type ioError struct {
	Op  string
	Err error
}

func (e *ioError) Error() string {
	if e.Err != nil {
		return "ioerror while " + e.Op + ": " + e.Err.Error()
	}
	return "ioerror while " + e.Op
}

type settingError struct {
	Op  string
	Err error
}

func (e *settingError) Error() string {
	if e.Err != nil {
		return "settingError while " + e.Op + ": " + e.Err.Error()
	}
	return "settingError while " + e.Op
}

var (
	errInvalidMode = errors.New("invalid mode")
)
