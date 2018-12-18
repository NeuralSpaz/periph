package tcs3272

// https://ams.com/documents/20143/36005/TCS3472_DS000390_2-00.pdf
import (
	"fmt"
	"time"

	"periph.io/x/periph/conn/physic"
)

type Opts struct{}

var Defaults Opts

func New() error { return nil }

type Dev struct{}

func (d *Dev) Sense() (Spectrum, error) { return Spectrum{}, nil }

type Spectrum struct {
	Bands             []Band
	SensorTemperature physic.Temperature
	Gain              string
	LedDrive          physic.ElectricCurrent
	Integration       time.Duration
}

func (s Spectrum) String() string {
	str := fmt.Sprintf("Spectrum: Gain:%s, Led Drive %s, Sense Time: %s", s.Gain, s.LedDrive, s.Integration)
	for _, band := range s.Bands {
		str += "\n" + band.String()
	}
	return str
}

type Band struct {
	Wavelength physic.Distance
	Value      float64
	Counts     uint16
	Name       string
}

func (b Band) String() string {
	return fmt.Sprintf("%s Band(%s) %7.1f counts", b.Name, b.Wavelength, b.Value)
}
