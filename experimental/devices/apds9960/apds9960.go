package apds9960

import (
	"encoding/binary"
	"sync"
	"time"

	"periph.io/x/periph/conn/i2c"
	"periph.io/x/periph/conn/mmr"
)

// Dev is an apds9960 and contain a mutex guarding the i2c device bus
type Dev struct {
	sync.Mutex
	c           mmr.Dev8
	interrupts  bool
	gestureData rawGestureData
}

// Opts holds the configuration options.
type Opts struct {
	Address uint16
}

// DefaultOpts are the recommended default options.
var DefaultOpts = Opts{}

// New initialises the apds9960
func New(b i2c.Bus, opts *Opts) (*Dev, error) {
	// {Conn: c, Order: binary.BigEndian}
	// c := &i2c.Dev{Bus: b, Addr: 0xD0}

	d := &Dev{
		c: mmr.Dev8{
			Conn: &i2c.Dev{
				Bus:  b,
				Addr: 0xD0,
			},
			Order: binary.LittleEndian},
	}
	// v, err := d.ReadUint8(0xD0)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// d := &Dev{c: &i2c.mm{Bus: bus, Addr: opts.Address}}

	// err := d.init()
	// TODO Configure
	time.Sleep(time.Millisecond * 500)
	return d, nil
}

type mode uint8

const (
	modePower           = 0
	modeAmbientLight    = 1
	modeProximity       = 2
	modeWait            = 3
	modeAmbientLightInt = 4
	modeProximityInt    = 5
	modeGesture         = 6
	modeAll             = 7
)

func (d *Dev) setMode(m mode, enable bool) error {
	d.Lock()
	defer d.Unlock()
	if m < 0 || m > 7 {
		return &settingError{"setMode()", errInvalidMode}
	}
	currentMode, err := d.c.ReadUint8(enableReg)
	if err != nil {
		return &ioError{"reading enable register", err}
	}
	var val = currentMode
	if m == modeAll {
		switch enable {
		case true:
			val = 0x7f
		case false:
			val = 0x00
		}
	} else if m != modeAll {
		switch enable {
		case true:
			val = setBit(currentMode, uint8(m))
		case false:
			val = clearBit(currentMode, uint8(m))
		}
	}
	if err := d.c.WriteUint8(enableReg, val); err != nil {
		return &ioError{"writing enable register", err}
	}
	return nil
}

func (d *Dev) init() error {

	err := d.setMode(modeAll, false)
	if err != nil {
		return err
	}

	defaults := make(map[byte][]byte)

	// defaults[apds9960Reg_ATIME] = []byte{0}
	defaults[aTimeReg] = []byte{0xFF}
	// defaults[apds9960Reg_WTIME] = []byte{171}
	defaults[wTimeReg] = []byte{0xFF}
	defaults[apds9960Reg_PPULSE] = []byte{0x87}
	defaults[apds9960Reg_POFFSET_UR] = []byte{0}
	defaults[apds9960Reg_POFFSET_DL] = []byte{0}
	defaults[apds9960Reg_CONFIG1] = []byte{0x60}
	defaults[apds9960Reg_PERS] = []byte{0x11}
	defaults[apds9960Reg_CONFIG2] = []byte{0x01}
	defaults[apds9960Reg_CONFIG3] = []byte{0x00}
	defaults[apds9960Reg_GOFFSET_U] = []byte{0x00}
	defaults[apds9960Reg_GOFFSET_D] = []byte{0x00}
	defaults[apds9960Reg_GOFFSET_L] = []byte{0x00}
	defaults[apds9960Reg_GOFFSET_R] = []byte{0x00}
	defaults[apds9960Reg_GPULSE] = []byte{0xc9}
	defaults[apds9960Reg_GCONF3] = []byte{0x00}

	d.Lock()
	for register, data := range defaults {
		if err := d.writeReg(register, data); err != nil {
			return err
		}
	}
	d.Unlock()

	if err = d.setLEDDrive(0x00); err != nil {
		return err
	}

	if err = d.setProximityGain(0x02); err != nil {
		return err
	}

	if err = d.setAmbientLightGain(0x01); err != nil {
		return err
	}

	if err = d.setProxIntLowThresh(0x00); err != nil {
		return err
	}

	if err = d.setProxIntHighThresh(50); err != nil {
		return err
	}

	if err = d.setLightIntLowThreshold(0xFFFF); err != nil {
		return err
	}

	if err = d.setLightIntHighThreshold(0x00); err != nil {
		return err
	}

	if err = d.setGestureEnterThresh(40); err != nil {
		return err
	}

	if err = d.setGestureExitThresh(30); err != nil {
		return err
	}

	if err = d.setGestureGain(0x02); err != nil {
		return err
	}

	if err = d.setGestureLEDDrive(0x00); err != nil {
		return err
	}

	if err = d.setGestureWaitTime(0x01); err != nil {
		return err
	}

	if err = d.setGestureIntEnable(0x00); err != nil {
		return err
	}

	return nil
}

// type LedDrive uint8

// const (
// 	LEDDrive100mA LedDrive = 0
// 	LEDDrive50mA           = 1
// 	LEDDrive25mA           = 2
// 	LEDDrive12mA           = 3
// )

func (d *Dev) setLEDDrive(drive uint8) error {
	d.Lock()
	defer d.Unlock()
	rx := make([]byte, 1)
	err := d.readReg(apds9960Reg_CONTROL, rx)
	if err != nil {
		return err
	}
	drive &= 0x03
	drive = drive << 6
	val := rx[0] & 0x3f
	val |= drive

	err = d.writeReg(apds9960Reg_CONTROL, []byte{val})
	return err
}

func (d *Dev) setProximityGain(drive uint8) error {
	d.Lock()
	defer d.Unlock()
	rx := make([]byte, 1)
	err := d.readReg(apds9960Reg_CONTROL, rx)
	if err != nil {
		return err
	}
	drive &= 0x03
	drive = drive << 2
	val := rx[0] & 0xf3
	val |= drive

	err = d.writeReg(apds9960Reg_CONTROL, []byte{val})
	return err
}

func (d *Dev) setAmbientLightGain(drive uint8) error {
	d.Lock()
	defer d.Unlock()
	rx := make([]byte, 1)
	err := d.readReg(apds9960Reg_CONTROL, rx)
	if err != nil {
		return err
	}
	drive &= 0x03
	val := rx[0] & 0xfc
	val |= drive

	err = d.writeReg(apds9960Reg_CONTROL, []byte{val})
	return err
}

func (d *Dev) setProxIntLowThresh(threshold uint8) error {
	d.Lock()
	defer d.Unlock()
	if err := d.c.WriteUint8(proxIntLowThresholdReg, threshold); err != nil {
		return &ioError{"writing proximity interupt low threshold", err}
	}
	return nil
}

func (d *Dev) setProxIntHighThresh(threshold uint8) error {
	d.Lock()
	defer d.Unlock()
	if err := d.c.WriteUint8(proxIntHighThresholdReg, threshold); err != nil {
		return &ioError{"writing proximity interupt high threshold", err}
	}
	return nil
}

func (d *Dev) setLightIntLowThreshold(threshold uint16) error {
	d.Lock()
	defer d.Unlock()
	if err := d.c.WriteUint16(lightIntLowThresholdReg, threshold); err != nil {
		return &ioError{"writing light interupt low threshold", err}
	}
	return nil
}

func (d *Dev) setLightIntHighThreshold(threshold uint16) error {
	d.Lock()
	defer d.Unlock()
	if err := d.c.WriteUint16(lightIntHighThresholdReg, threshold); err != nil {
		return &ioError{"writing light interupt high threshold", err}
	}
	return nil
}

func (d *Dev) setGestureEnterThresh(threshold uint8) error {
	d.Lock()
	defer d.Unlock()
	if err := d.c.WriteUint8(gestureEnterThresholdReg, threshold); err != nil {
		return &ioError{"writing gesture enter threshold", err}
	}
	return nil
}

func (d *Dev) setGestureExitThresh(threshold uint8) error {
	d.Lock()
	defer d.Unlock()
	if err := d.c.WriteUint8(gestureExitThresholdReg, threshold); err != nil {
		return &ioError{"writing gesture exit threshold", err}
	}
	return nil
}

func (d *Dev) setGestureGain(gain uint8) error {
	d.Lock()
	defer d.Unlock()
	rx := make([]byte, 1)
	err := d.readReg(apds9960Reg_GCONF2, rx)
	if err != nil {
		return err
	}

	/* Set bits in register to given value */
	gain &= 0x03
	gain = gain << 5
	val := rx[0] & 0x9F
	val |= gain

	err = d.readReg(apds9960Reg_GCONF2, []byte{val})
	return err
}

func (d *Dev) setGestureLEDDrive(drive uint8) error {
	d.Lock()
	defer d.Unlock()
	rx := make([]byte, 1)
	err := d.readReg(apds9960Reg_GCONF2, rx)
	if err != nil {
		return err
	}

	/* Set bits in register to given value */
	drive &= 0x03
	drive = drive << 3
	val := rx[0] & 0xe7
	val |= drive

	err = d.readReg(apds9960Reg_GCONF2, []byte{val})
	return err
}

func (d *Dev) setGestureWaitTime(time uint8) error {
	d.Lock()
	defer d.Unlock()
	rx := make([]byte, 1)
	err := d.readReg(apds9960Reg_GCONF2, rx)
	if err != nil {
		return err
	}

	/* Set bits in register to given value */
	time &= 0x07
	val := rx[0] & 0xf8
	val |= time

	err = d.readReg(apds9960Reg_GCONF2, []byte{val})
	return err
}

func (d *Dev) setGestureIntEnable(enable uint8) error {
	d.Lock()
	defer d.Unlock()
	rx := make([]byte, 1)
	err := d.readReg(apds9960Reg_GCONF4, rx)
	if err != nil {
		return err
	}

	/* Set bits in register to given value */
	enable &= 0x01
	enable = enable << 1
	val := rx[0] & 0xfd
	val |= enable

	err = d.readReg(apds9960Reg_GCONF4, []byte{val})
	return err
}

func (d *Dev) setAmbientLightIntEnable(enable bool) error {
	d.Lock()
	defer d.Unlock()
	rx := make([]byte, 1)
	err := d.readReg(enableReg, rx)
	if err != nil {
		return err
	}
	var enablebit uint8 = 1
	if !enable {
		enablebit = 0
	}
	/* Set bits in register to given value */
	enablebit &= 0x01
	enablebit = enablebit << 4
	val := rx[0] & 0xef
	val |= enablebit

	err = d.readReg(enableReg, []byte{val})
	return err
}

func (d *Dev) enablePower() error {
	return d.setMode(modePower, true)
}

func (d *Dev) setGestureMode(mode bool) error {
	d.Lock()
	defer d.Unlock()
	rx := make([]byte, 1)

	if err := d.readReg(apds9960Reg_GCONF4, rx); err != nil {
		return err
	}
	if mode {
		rx[0] = setBit(rx[0], 0)
	} else {
		rx[0] = clearBit(rx[0], 0)
	}

	if err := d.writeReg(apds9960Reg_GCONF4, rx); err != nil {
		return err
	}

	return nil
}

func (d *Dev) setLEDBoost(boost uint8) error {
	d.Lock()
	defer d.Unlock()
	rx := make([]byte, 1)
	err := d.readReg(apds9960Reg_GCONF2, rx)
	if err != nil {
		return err
	}

	/* Set bits in register to given value */
	boost &= 0x03
	boost = boost << 4
	val := rx[0] & 0xCF
	val |= boost

	err = d.readReg(apds9960Reg_GCONF2, []byte{boost})
	return err
}
