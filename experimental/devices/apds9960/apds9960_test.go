package apds9960

import (
	"encoding/binary"
	"testing"

	"periph.io/x/periph/conn/i2c"
	"periph.io/x/periph/conn/i2c/i2ctest"
	"periph.io/x/periph/conn/mmr"
)

func TestDev_setMode(t *testing.T) {
	const ()
	tests := []struct {
		name    string
		tx      []i2ctest.IO
		mode    mode
		enable  bool
		wantErr error
	}{
		{name: "modePower", mode: modePower, enable: true, wantErr: nil,
			tx: []i2ctest.IO{
				{Addr: 0x39, W: []byte{enableReg}, R: []byte{0x00}},
				{Addr: 0x39, W: []byte{enableReg, 0x01}},
			},
		},
		{name: "modeAmbientLight", mode: modeAmbientLight, enable: true, wantErr: nil,
			tx: []i2ctest.IO{
				{Addr: 0x39, W: []byte{enableReg}, R: []byte{0x00}},
				{Addr: 0x39, W: []byte{enableReg, 0x02}},
			}},
		{name: "modeProximity", mode: modeProximity, enable: true, wantErr: nil,
			tx: []i2ctest.IO{
				{Addr: 0x39, W: []byte{enableReg}, R: []byte{0x00}},
				{Addr: 0x39, W: []byte{enableReg, 0x04}},
			}},
		{name: "modeWait", mode: modeWait, enable: true, wantErr: nil,
			tx: []i2ctest.IO{
				{Addr: 0x39, W: []byte{enableReg}, R: []byte{0x00}},
				{Addr: 0x39, W: []byte{enableReg, 0x08}},
			}},
		{name: "modeAmbientLightInt", mode: modeAmbientLightInt, enable: true, wantErr: nil,
			tx: []i2ctest.IO{
				{Addr: 0x39, W: []byte{enableReg}, R: []byte{0x00}},
				{Addr: 0x39, W: []byte{enableReg, 0x010}},
			}},
		{name: "modeProximityInt", mode: modeProximityInt, enable: true, wantErr: nil,
			tx: []i2ctest.IO{
				{Addr: 0x39, W: []byte{enableReg}, R: []byte{0x00}},
				{Addr: 0x39, W: []byte{enableReg, 0x20}},
			}},
		{name: "modeGesture", mode: modeGesture, enable: true, wantErr: nil,
			tx: []i2ctest.IO{
				{Addr: 0x39, W: []byte{enableReg}, R: []byte{0x00}},
				{Addr: 0x39, W: []byte{enableReg, 0x40}},
			}},
		{name: "modeAll", mode: modeAll, enable: true, wantErr: nil,
			tx: []i2ctest.IO{
				{Addr: 0x39, W: []byte{enableReg}, R: []byte{0x00}},
				{Addr: 0x39, W: []byte{enableReg, 0x7f}},
			}},
		{name: "disableAll", mode: modeAll, enable: false, wantErr: nil,
			tx: []i2ctest.IO{
				{Addr: 0x39, W: []byte{enableReg}, R: []byte{0x00}},
				{Addr: 0x39, W: []byte{enableReg, 0x00}},
			}},
		{name: "disableGesture", mode: modeGesture, enable: false, wantErr: nil,
			tx: []i2ctest.IO{
				{Addr: 0x39, W: []byte{enableReg}, R: []byte{0x40}},
				{Addr: 0x39, W: []byte{enableReg, 0x00}},
			}},
		{name: "errInvalidMode", wantErr: errInvalidMode, mode: mode(255)},
		{name: "settingError", wantErr: &settingError{}, mode: mode(255)},
		{name: "read_ioError", wantErr: &ioError{}},
		{name: "write_ioError", wantErr: &ioError{},
			tx: []i2ctest.IO{
				{Addr: 0x39, W: []byte{enableReg}, R: []byte{0x00}},
			}},
	}
	for _, tt := range tests {

		b := &i2ctest.Playback{
			Ops:       tt.tx,
			DontPanic: true,
		}
		d := &Dev{
			c: mmr.Dev8{
				Conn: &i2c.Dev{
					Bus:  b,
					Addr: 0x39,
				},
				Order: binary.LittleEndian},
		}

		err := d.setMode(tt.mode, tt.enable)
		if err != nil && tt.wantErr == nil {
			t.Errorf("Dev.setLEDDrive() case: %s, unexpected error =%v", tt.name, err)

		}
		switch tt.wantErr.(type) {
		case *ioError:
			if _, ok := err.(*ioError); !ok {
				t.Errorf("Dev.setLEDDrive() case: %s, expected error of type *ioError but got %v of type %T", tt.name, err, err)
			}
		case *settingError:
			if _, ok := err.(*settingError); !ok {
				t.Errorf("Dev.setLEDDrive() case: %s, expected error of type *settingError but got %v of type %T", tt.name, err, err)
			}
		}
	}
}

func TestDev_setProxIntLowThresh(t *testing.T) {
	tests := []struct {
		name      string
		threshold uint8
		tx        []i2ctest.IO
		wantErr   error
	}{
		{name: "write_ioError", wantErr: &ioError{}},
		{name: "nil", wantErr: nil,
			tx: []i2ctest.IO{
				{Addr: 0x39, W: []byte{proxIntLowThresholdReg, 0x00}, R: []byte{}},
			},
		},
		{name: "0x7F", wantErr: nil, threshold: 127,
			tx: []i2ctest.IO{
				{Addr: 0x39, W: []byte{proxIntLowThresholdReg, 0x7F}, R: []byte{}},
			},
		},
	}
	for _, tt := range tests {

		b := &i2ctest.Playback{
			Ops:       tt.tx,
			DontPanic: true,
		}
		d := &Dev{
			c: mmr.Dev8{
				Conn: &i2c.Dev{
					Bus:  b,
					Addr: 0x39,
				},
				Order: binary.LittleEndian},
		}

		err := d.setProxIntLowThresh(tt.threshold)
		if err != nil && tt.wantErr == nil {
			t.Errorf("Dev.setProxIntLowThresh() case: %s, unexpected error =%v", tt.name, err)

		}
		switch tt.wantErr.(type) {
		case *ioError:
			if _, ok := err.(*ioError); !ok {
				t.Errorf("Dev.setProxIntLowThresh() case: %s, expected error of type *ioError but got %v of type %T", tt.name, err, err)
			}
		case *settingError:
			if _, ok := err.(*settingError); !ok {
				t.Errorf("Dev.setProxIntLowThresh() case: %s, expected error of type *settingError but got %v of type %T", tt.name, err, err)
			}
		}
	}
}

func TestDev_setProxIntHighThresh(t *testing.T) {
	tests := []struct {
		name      string
		threshold uint8
		tx        []i2ctest.IO
		wantErr   error
	}{
		{name: "write_ioError", wantErr: &ioError{}},
		{name: "nil", wantErr: nil,
			tx: []i2ctest.IO{
				{Addr: 0x39, W: []byte{proxIntHighThresholdReg, 0x00}, R: []byte{}},
			},
		},
		{name: "0x7F", wantErr: nil, threshold: 127,
			tx: []i2ctest.IO{
				{Addr: 0x39, W: []byte{proxIntHighThresholdReg, 0x7F}, R: []byte{}},
			},
		},
	}
	for _, tt := range tests {

		b := &i2ctest.Playback{
			Ops:       tt.tx,
			DontPanic: true,
		}
		d := &Dev{
			c: mmr.Dev8{
				Conn: &i2c.Dev{
					Bus:  b,
					Addr: 0x39,
				},
				Order: binary.LittleEndian},
		}

		err := d.setProxIntHighThresh(tt.threshold)
		if err != nil && tt.wantErr == nil {
			t.Errorf("Dev.setProxIntHighThresh() case: %s, unexpected error =%v", tt.name, err)

		}
		switch tt.wantErr.(type) {
		case *ioError:
			if _, ok := err.(*ioError); !ok {
				t.Errorf("Dev.setProxIntHighThresh() case: %s, expected error of type *ioError but got %v of type %T", tt.name, err, err)
			}
		case *settingError:
			if _, ok := err.(*settingError); !ok {
				t.Errorf("Dev.setProxIntHighThresh() case: %s, expected error of type *settingError but got %v of type %T", tt.name, err, err)
			}
		}
	}
}

func TestDev_setLightIntLowThreshold(t *testing.T) {
	tests := []struct {
		name      string
		threshold uint16
		tx        []i2ctest.IO
		wantErr   error
	}{
		{name: "write_ioError", wantErr: &ioError{}},
		{name: "nil", wantErr: nil,
			tx: []i2ctest.IO{
				{Addr: 0x39, W: []byte{lightIntLowThresholdReg, 0x00, 0x00}, R: []byte{}},
			},
		},
		{name: "FF00", wantErr: nil, threshold: 0xFF00,
			tx: []i2ctest.IO{
				{Addr: 0x39, W: []byte{lightIntLowThresholdReg, 0x00, 0xFF}, R: []byte{}},
			},
		},
	}
	for _, tt := range tests {

		b := &i2ctest.Playback{
			Ops:       tt.tx,
			DontPanic: true,
		}
		d := &Dev{
			c: mmr.Dev8{
				Conn: &i2c.Dev{
					Bus:  b,
					Addr: 0x39,
				},
				Order: binary.LittleEndian},
		}

		err := d.setLightIntLowThreshold(tt.threshold)
		if err != nil && tt.wantErr == nil {
			t.Errorf("Dev.setLightIntLowThreshold() case: %s, unexpected error =%v", tt.name, err)

		}
		switch tt.wantErr.(type) {
		case *ioError:
			if _, ok := err.(*ioError); !ok {
				t.Errorf("Dev.setLightIntLowThreshold() case: %s, expected error of type *ioError but got %v of type %T", tt.name, err, err)
			}
		case *settingError:
			if _, ok := err.(*settingError); !ok {
				t.Errorf("Dev.setLightIntLowThreshold() case: %s, expected error of type *settingError but got %v of type %T", tt.name, err, err)
			}
		}
	}
}

func TestDev_setLightIntHighThreshold(t *testing.T) {
	tests := []struct {
		name      string
		threshold uint16
		tx        []i2ctest.IO
		wantErr   error
	}{
		{name: "write_ioError", wantErr: &ioError{}},
		{name: "nil", wantErr: nil,
			tx: []i2ctest.IO{
				{Addr: 0x39, W: []byte{lightIntHighThresholdReg, 0x00, 0x00}, R: []byte{}},
			},
		},
		{name: "FF00", wantErr: nil, threshold: 0xFF00,
			tx: []i2ctest.IO{
				{Addr: 0x39, W: []byte{lightIntHighThresholdReg, 0x00, 0xFF}, R: []byte{}},
			},
		},
	}
	for _, tt := range tests {

		b := &i2ctest.Playback{
			Ops:       tt.tx,
			DontPanic: true,
		}
		d := &Dev{
			c: mmr.Dev8{
				Conn: &i2c.Dev{
					Bus:  b,
					Addr: 0x39,
				},
				Order: binary.LittleEndian},
		}

		err := d.setLightIntHighThreshold(tt.threshold)
		if err != nil && tt.wantErr == nil {
			t.Errorf("Dev.setLightIntHighThreshold() case: %s, unexpected error =%v", tt.name, err)

		}
		switch tt.wantErr.(type) {
		case *ioError:
			if _, ok := err.(*ioError); !ok {
				t.Errorf("Dev.setLightIntHighThreshold() case: %s, expected error of type *ioError but got %v of type %T", tt.name, err, err)
			}
		case *settingError:
			if _, ok := err.(*settingError); !ok {
				t.Errorf("Dev.setLightIntHighThreshold() case: %s, expected error of type *settingError but got %v of type %T", tt.name, err, err)
			}
		}
	}
}

func TestDev_setGestureEnterThresh(t *testing.T) {
	tests := []struct {
		name      string
		threshold uint8
		tx        []i2ctest.IO
		wantErr   error
	}{
		{name: "write_ioError", wantErr: &ioError{}},
		{name: "nil", wantErr: nil,
			tx: []i2ctest.IO{
				{Addr: 0x39, W: []byte{gestureEnterThresholdReg, 0x00}, R: []byte{}},
			},
		},
		{name: "0x7F", wantErr: nil, threshold: 127,
			tx: []i2ctest.IO{
				{Addr: 0x39, W: []byte{gestureEnterThresholdReg, 0x7F}, R: []byte{}},
			},
		},
	}
	for _, tt := range tests {

		b := &i2ctest.Playback{
			Ops:       tt.tx,
			DontPanic: true,
		}
		d := &Dev{
			c: mmr.Dev8{
				Conn: &i2c.Dev{
					Bus:  b,
					Addr: 0x39,
				},
				Order: binary.LittleEndian},
		}

		err := d.setGestureEnterThresh(tt.threshold)
		if err != nil && tt.wantErr == nil {
			t.Errorf("Dev.setGestureEnterThresh() case: %s, unexpected error =%v", tt.name, err)

		}
		switch tt.wantErr.(type) {
		case *ioError:
			if _, ok := err.(*ioError); !ok {
				t.Errorf("Dev.setGestureEnterThresh() case: %s, expected error of type *ioError but got %v of type %T", tt.name, err, err)
			}
		case *settingError:
			if _, ok := err.(*settingError); !ok {
				t.Errorf("Dev.setGestureEnterThresh() case: %s, expected error of type *settingError but got %v of type %T", tt.name, err, err)
			}
		}
	}
}

func TestDev_setGestureExitThresh(t *testing.T) {
	tests := []struct {
		name      string
		threshold uint8
		tx        []i2ctest.IO
		wantErr   error
	}{
		{name: "write_ioError", wantErr: &ioError{}},
		{name: "nil", wantErr: nil,
			tx: []i2ctest.IO{
				{Addr: 0x39, W: []byte{gestureExitThresholdReg, 0x00}, R: []byte{}},
			},
		},
		{name: "0x7F", wantErr: nil, threshold: 127,
			tx: []i2ctest.IO{
				{Addr: 0x39, W: []byte{gestureExitThresholdReg, 0x7F}, R: []byte{}},
			},
		},
	}
	for _, tt := range tests {

		b := &i2ctest.Playback{
			Ops:       tt.tx,
			DontPanic: true,
		}
		d := &Dev{
			c: mmr.Dev8{
				Conn: &i2c.Dev{
					Bus:  b,
					Addr: 0x39,
				},
				Order: binary.LittleEndian},
		}

		err := d.setGestureExitThresh(tt.threshold)
		if err != nil && tt.wantErr == nil {
			t.Errorf("Dev.setGestureExitThresh() case: %s, unexpected error =%v", tt.name, err)

		}
		switch tt.wantErr.(type) {
		case *ioError:
			if _, ok := err.(*ioError); !ok {
				t.Errorf("Dev.setGestureExitThresh() case: %s, expected error of type *ioError but got %v of type %T", tt.name, err, err)
			}
		case *settingError:
			if _, ok := err.(*settingError); !ok {
				t.Errorf("Dev.setGestureExitThresh() case: %s, expected error of type *settingError but got %v of type %T", tt.name, err, err)
			}
		}
	}
}
