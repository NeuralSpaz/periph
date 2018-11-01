package apds9960

//go:generate stringer -type=Gesture

import (
	"errors"
	"time"
)

type Gesture int

const (
	None Gesture = iota
	Left
	Right
	Up
	Down
	Near
	Far
	All
)

type rawGestureData struct {
	udata        []byte
	ddata        []byte
	ldata        []byte
	rdata        []byte
	index        uint8
	total        uint8
	inThreshold  uint8
	outThreshold uint8
}

func (d *Dev) StartGestureEngine(interval time.Duration) (<-chan Gesture, chan bool, <-chan error) {

	gestures := make(chan Gesture)
	pause := make(chan bool)
	errc := make(chan error, 3)
	if interval < time.Millisecond*5 {
		interval = time.Millisecond * 5
		errc <- errors.New("interval time too short setting to 5 milliseconds")
	}
	err := d.enableGesture()
	if err != nil {
		errc <- errors.New("unable to set gesture mode")
	}
	go func() {
		poll := time.NewTimer(interval)
		for {
			select {
			case <-poll.C:
				// fmt.Println("checking interputs")
				available, err := d.isGestureAvailable()
				if err != nil {
					errc <- err
				}
				if available {
					gesture, err := d.getGesture()
					if err != nil {
						errc <- err
					}
					for _, g := range gesture {
						gestures <- g
					}
				}
				poll.Reset(interval)
			case p := <-pause:
				if p {
					if !poll.Stop() {
						<-poll.C
					}
				}
				if !p {
					poll.Reset(interval)
				}
			}
		}
	}()
	return gestures, pause, errc
}

func (d *Dev) getGesture() ([]Gesture, error) {
	gestureData, n, err := d.readGestures()
	if err != nil {
		return []Gesture{}, nil
	}
	if n != 0 {
		g := d.processGestures(gestureData)
		return g, nil

	}
	return []Gesture{}, nil
}

func (d *Dev) readGestures() (rawGestureData, int, error) {
	rx := make([]byte, 1)
	d.Lock()
	err := d.readReg(apds9960Reg_GFLVL, rx)
	d.Unlock()
	if err != nil {
		return rawGestureData{}, 0, err
	}
	fifoLength := uint8(rx[0])
	if fifoLength > 0 {
		rx = make([]byte, fifoLength*4)
		d.Lock()
		err := d.readReg(apds9960Reg_GFIFO_U, rx)
		d.Unlock()
		if err != nil {
			return rawGestureData{}, 0, err
		}
		// fmt.Printf("RawGestureDATA: %x\n", rx)
		// uData := rx[:fifoLength]
		// dData := rx[fifoLength : fifoLength*2]
		// lData := rx[fifoLength*2 : fifoLength*3]
		// rData := rx[fifoLength*3 : fifoLength*4]
		uData := make([]byte, 0)
		dData := make([]byte, 0)
		lData := make([]byte, 0)
		rData := make([]byte, 0)
		for i := 0; i < int(fifoLength*4); i += 4 {
			uData = append(uData, rx[i])
			dData = append(dData, rx[i+1])
			lData = append(lData, rx[i+2])
			rData = append(rData, rx[i+3])
		}
		gestureData := rawGestureData{
			udata: uData,
			ddata: dData,
			ldata: lData,
			rdata: rData,
		}
		return gestureData, int(fifoLength), nil
	}

	return rawGestureData{}, 0, nil
}

func (d *Dev) processGestures(data rawGestureData) []Gesture {
	threshold := byte(0x01)
	incomming := 0
	for i := 0; i < len(data.udata); i++ {
		if (data.udata[i] > threshold) &&
			(data.ddata[i] > threshold) &&
			(data.ldata[i] > threshold) &&
			(data.rdata[i] > threshold) {
			d.gestureData.udata = append(d.gestureData.udata, data.udata[i])
			d.gestureData.ddata = append(d.gestureData.ddata, data.ddata[i])
			d.gestureData.ldata = append(d.gestureData.ldata, data.ldata[i])
			d.gestureData.rdata = append(d.gestureData.rdata, data.rdata[i])
			incomming++
		}
	}
	last := len(d.gestureData.udata)
	Δud := 0
	Δlr := 0
	if last != 0 {
		for k := 0; k < len(d.gestureData.udata)-1; k++ {
			δud := 0
			δlr := 0
			up := int(int8(d.gestureData.udata[k]))
			down := int(int8(d.gestureData.ddata[k]))
			left := int(int8(d.gestureData.ldata[k]))
			right := int(int8(d.gestureData.rdata[k]))

			u1 := int(int8(d.gestureData.udata[k+1]))
			d1 := int(int8(d.gestureData.ddata[k+1]))
			l1 := int(int8(d.gestureData.ldata[k+1]))
			r1 := int(int8(d.gestureData.rdata[k+1]))

			εud := (up + down)
			εlr := (left + right)
			εud1 := (u1 + d1)
			εlr1 := (l1 + r1)

			if εud != 0 && εlr != 0 && εud1 != 0 && εlr1 != 0 {
				ratioUD := (up - down) * 100 / εud
				ratioLR := (left - right) * 100 / εlr

				ratioUD1 := (u1 - d1) * 100 / εud1
				ratioLR1 := (l1 - r1) * 100 / εlr1

				δud += ratioUD1 - ratioUD
				δlr += ratioLR1 - ratioLR

				Δud += δud
				Δlr += δlr
			}
		}
	}
	// fmt.Printf("\n\n%6.6f, %6.6f\n\n", Δud, Δlr)
	// }

	if incomming != 0 && len(d.gestureData.udata) != 0 {
		g := decodeDirection(Δud, Δlr)
		// fmt.Println(g)
		d.gestureData = rawGestureData{}
		return []Gesture{g}
	}

	return []Gesture{}
}

func decodeDirection(ud, lr int) Gesture {
	var dir Gesture
	// fmt.Println(ud, lr)
	if ud*ud > lr*lr {
		if ud < 0 {
			dir = Down
		} else {
			dir = Up
		}
	} else {
		if lr < 0 {
			dir = Right
		} else {
			dir = Left
		}
	}
	return dir
}

func (d *Dev) isGestureAvailable() (bool, error) {
	rx := make([]byte, 1)
	d.Lock()
	err := d.readReg(apds9960Reg_GSTATUS, rx)
	d.Unlock()
	if err != nil {
		return false, err
	}
	// fmt.Println(rx[0])

	return hasBit(rx[0], 0), nil
}

func (d *Dev) enableGesture() error {
	/* Enable gesture mode
	   Set ENABLE to 0 (power off)
	   Set WTIME to 0xFF
	   Set AUX to LED_BOOST_300
	   Enable PON, WEN, PEN, GEN in ENABLE
	*/

	gDefaults := make(map[byte][]byte)
	gDefaults[wTimeReg] = []byte{0xFF}
	gDefaults[apds9960Reg_PPULSE] = []byte{0x89}

	d.setMode(modePower, false)

	d.Lock()
	for register, data := range gDefaults {
		if err := d.writeReg(register, data); err != nil {
			return err
		}
	}
	d.Unlock()

	if err := d.setLEDBoost(0x03); err != nil {
		return err
	}
	if err := d.setGestureMode(true); err != nil {
		return err
	}
	if err := d.setMode(modePower, true); err != nil {
		return err
	}
	if err := d.setMode(modeWait, true); err != nil {
		return err
	}
	if err := d.setMode(modeProximity, true); err != nil {
		return err
	}
	if err := d.setMode(modeGesture, true); err != nil {
		return err
	}

	return nil
}
