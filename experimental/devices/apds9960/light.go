package apds9960

import (
	"errors"
	"fmt"
	"image/color"
	"log"
	"time"
)

func (d *Dev) StartLightSensor(interval time.Duration) (<-chan color.Color, chan bool, error) {
	if interval < time.Millisecond*10 {
		return nil, nil, errors.New("interval to short for reading colour sensor")
	}
	if err := d.enableLightSensor(); err != nil {
		return nil, nil, fmt.Errorf("error enabling light sensor: %v", err)
	}
	time.Sleep(time.Second * 5)
	c := make(chan color.Color)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-time.After(interval):
				fmt.Println("check color")
				color, err := d.readColorSensor()
				if err != nil {
					log.Println("error reading light sesnor ", err)
				}
				c <- color
			case <-done:
				fmt.Println("cancel reading colour")
				if err := d.disableLightSensor(); err != nil {
					log.Println("error enabling light sensor ", err)
				}
				close(done)
				close(c)
				return
			}

		}
	}()
	return c, done, nil
}

func (d *Dev) enableLightSensor() error {

	d.setAmbientLightGain(1)          //4x
	d.setAmbientLightIntEnable(false) //no interput
	d.enablePower()
	d.setMode(modeAmbientLight, true)
	d.setMode(modeWait, false)
	return nil
}

func (d *Dev) disableLightSensor() error {
	d.Lock()
	//DO STUFF
	d.Unlock()
	return nil
}

func (d *Dev) readColorSensor() (color.Color, error) {

	clear, err := d.readClearLight()
	if err != nil {
		return color.RGBA64{}, err
	}
	red, err := d.readRedLight()
	if err != nil {
		return color.RGBA64{}, err
	}
	green, err := d.readGreenLight()
	if err != nil {
		return color.RGBA64{}, err
	}
	blue, err := d.readBlueLight()
	if err != nil {
		return color.RGBA64{}, err
	}
	return color.RGBA64{red, green, blue, clear}, err
}

func (d *Dev) readRedLight() (uint16, error) {
	d.Lock()
	rx := make([]byte, 2)
	err := d.readReg(apds9960Reg_RDATAL, rx)
	d.Unlock()
	fmt.Println(rx)
	return (uint16(rx[0]) | (uint16(rx[1]) << 8)), err
}
func (d *Dev) readGreenLight() (uint16, error) {
	d.Lock()
	rx := make([]byte, 2)
	err := d.readReg(apds9960Reg_GDATAL, rx)
	d.Unlock()
	fmt.Println(rx)
	return (uint16(rx[0]) | (uint16(rx[1]) << 8)), err
}
func (d *Dev) readBlueLight() (uint16, error) {
	d.Lock()
	rx := make([]byte, 2)
	err := d.readReg(apds9960Reg_BDATAL, rx)
	d.Unlock()
	fmt.Println(rx)
	return (uint16(rx[0]) | (uint16(rx[1]) << 8)), err
}
func (d *Dev) readClearLight() (uint16, error) {
	d.Lock()
	rx := make([]byte, 2)
	err := d.readReg(apds9960Reg_CDATAL, rx)
	d.Unlock()
	fmt.Println(rx)
	return (uint16(rx[0]) | (uint16(rx[1]) << 8)), err
}
