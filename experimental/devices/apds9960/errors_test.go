package apds9960

import (
	"errors"
	"testing"
)

func Test_ioError_Error(t *testing.T) {
	var errTestErr = errors.New("test")
	tests := []struct {
		name string
		Op   string
		Err  error
		want string
	}{
		{name: "nil", want: "ioerror while "},
		{name: "errTestErr", Err: errTestErr, want: "ioerror while : test"},
	}
	for _, tt := range tests {
		e := &ioError{
			Op:  tt.Op,
			Err: tt.Err,
		}
		if got := e.Error(); got != tt.want {
			t.Errorf("ioError.Error() = %v, want %v", got, tt.want)
		}
	}
}

func Test_settingError_Error(t *testing.T) {
	var errTestErr = errors.New("test")
	tests := []struct {
		name string
		Op   string
		Err  error
		want string
	}{
		{name: "nil", want: "settingError while "},
		{name: "errTestErr", Err: errTestErr, want: "settingError while : test"},
	}
	for _, tt := range tests {
		e := &settingError{
			Op:  tt.Op,
			Err: tt.Err,
		}
		if got := e.Error(); got != tt.want {
			t.Errorf("settingError.Error() = %v, want %v", got, tt.want)
		}
	}
}
