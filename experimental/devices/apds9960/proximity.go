package apds9960

import (
	"errors"
	"log"
	"time"
)

func (d *Dev) StartProxSensor(interval time.Duration) (<-chan float64, chan bool, error) {
	if interval < time.Millisecond*1 {
		return nil, nil, errors.New("interval to short for reading colour sensor")
	}
	// if err := s.enableLightSensor(); err != nil {
	// 	return nil, nil, fmt.Errorf("error enabling light sensor: %v", err)
	// }
	// time.Sleep(time.Second * 5)
	d.enableProximitySensor()

	c := make(chan float64)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-time.After(interval):
				distance, err := d.readProximitySensor()
				if err != nil {
					log.Println("error reading light sesnor ", err)
				}
				c <- distance
			case <-done:
				// fmt.Println("cancel reading colour")
				// if err := d.disableLightSensor(); err != nil {
				// 	log.Println("error enabling light sensor ", err)
				// }
				close(done)
				close(c)
				return
			}

		}
	}()
	return c, done, nil
}

func (d *Dev) readProximitySensor() (float64, error) {
	d.Lock()
	rx := make([]byte, 1)
	err := d.readReg(apds9960Reg_PDATA, rx)
	d.Unlock()
	// fmt.Println(rx)
	return float64(uint8(rx[0])), err
}

func (d *Dev) enableProximitySensor() error {
	if err := d.setLEDDrive(0x00); err != nil {
		return err
	}

	if err := d.setProximityGain(0x02); err != nil {
		return err
	}
	// if d.interrupts {
	// 	if err := d.setProximityIntEnable(true); err != nil {
	// 		return err
	// 	}
	// } else {
	// 	if err := d.setProximityIntEnable(false); err != nil {
	// 		return err
	// 	}
	// }

	if err := d.enablePower(); err != nil {
		return err
	}

	if err := d.setMode(modeProximity, true); err != nil {
		return err
	}
	return nil
}

// func (d *Dev) setProximityIntEnable(enable bool) error { return nil }
