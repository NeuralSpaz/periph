// Copyright 2018 The Periph Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

package physic

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

// Angle is the measurement of the difference in orientation between two vectors
// stored as an int64 nano radian.
//
// A negative angle is valid.
//
// The highest representable value is a bit over 500,000,000,000°.
type Angle int64

// String returns the angle formatted as a string in degree.
func (a Angle) String() string {
	// Angle is not a S.I. unit, so it must not be prefixed by S.I. prefixes.
	if a == 0 {
		return "0°"
	}
	// Round.
	prefix := ""
	if a < 0 {
		a = -a
		prefix = "-"
	}
	switch {
	case a < Degree:
		v := ((a * 1000) + Degree/2) / Degree
		return prefix + "0." + prefixZeros(3, int(v)) + "°"
	case a < 10*Degree:
		v := ((a * 1000) + Degree/2) / Degree
		i := v / 1000
		v = v - i*1000
		return prefix + strconv.FormatInt(int64(i), 10) + "." + prefixZeros(3, int(v)) + "°"
	case a < 100*Degree:
		v := ((a * 1000) + Degree/2) / Degree
		i := v / 1000
		v = v - i*1000
		return prefix + strconv.FormatInt(int64(i), 10) + "." + prefixZeros(2, int(v)) + "°"
	case a < 1000*Degree:
		v := ((a * 1000) + Degree/2) / Degree
		i := v / 1000
		v = v - i*1000
		return prefix + strconv.FormatInt(int64(i), 10) + "." + prefixZeros(1, int(v)) + "°"
	default:
		v := (a + Degree/2) / Degree
		return prefix + strconv.FormatInt(int64(v), 10) + "°"
	}
}

const (
	NanoRadian  Angle = 1
	MicroRadian Angle = 1000 * NanoRadian
	MilliRadian Angle = 1000 * MicroRadian
	Radian      Angle = 1000 * MilliRadian

	// Theta is 2π. This is equivalent to 360°.
	Theta  Angle = 6283185307 * NanoRadian
	Pi     Angle = 3141592653 * NanoRadian
	Degree Angle = 17453293 * NanoRadian
)

// Set takes a string and tries to return a valid Angle.
func (a *Angle) Set(s string) error {
	v, n, err := setInt(s, nano)
	if err != nil {
		return err
	}
	switch s[n:] {
	case "°", "Degrees", "degrees":
		*a = (Angle)(v * int64(Degree) / int64(Radian))
	case "Pi", "pi", "π":
		*a = (Angle)(v * int64(Pi) / int64(Radian))
	case "Radians", "radians", "Radian", "radian":
		*a = (Angle)(v)
	default:
		return noUnits("Radian")
	}
	return nil
}

// Distance is a measurement of length stored as an int64 nano metre.
//
// This is one of the base unit in the International System of Units.
//
// The highest representable value is 9.2Gm.
type Distance int64

// String returns the distance formatted as a string in metre.
func (d Distance) String() string {
	return nanoAsString(int64(d)) + "m"
}

const (
	NanoMetre  Distance = 1
	MicroMetre Distance = 1000 * NanoMetre
	MilliMetre Distance = 1000 * MicroMetre
	Metre      Distance = 1000 * MilliMetre
	KiloMetre  Distance = 1000 * Metre
	MegaMetre  Distance = 1000 * KiloMetre
	GigaMetre  Distance = 1000 * MegaMetre

	// Conversion between Metre and imperial units.
	Thou Distance = 25400 * NanoMetre
	Inch Distance = 1000 * Thou
	Foot Distance = 12 * Inch
	Yard Distance = 3 * Foot
	Mile Distance = 1760 * Yard
)

// Set takes a string and tries to return a valid Distance.
func (d *Distance) Set(s string) error {
	dec, n, err := atod(s)
	if err != nil {
		return err
	}
	prefix := prefix(none)
	if !(n == len(s)) {
		var n1 int
		prefix, n1 = parseSIPrefix([]rune(s[n:])[0])
		if prefix == milli && !(s[n:] == "mm") {
			prefix = none
		}
		if prefix == mega && !(s[n:] == "Mm") {
			prefix = none
		}
		if prefix != none {
			n += n1
		}
	}
	v, err := dec.dtoi(int(prefix - nano))
	if err != nil {
		return err
	}
	switch s[n:] {
	case "mile", "Mile", "miles", "Miles":
		*d = (Distance)((v * 1609344) / 1000)
	case "yard", "Yard", "yards", "Yards":
		*d = (Distance)((v * 9144) / 10000)
	case "foot", "Foot", "Feet", "feet", "ft", "Ft":
		*d = (Distance)((v * 3048) / 10000)
	case "in", "In", "inch", "Inch", "inches", "Inches":
		*d = (Distance)((v * 254) / 10000)
	case "m", "metre", "metres", "Metre", "Metres":
		*d = (Distance)(v)
	default:
		return noUnits("m")
	}
	return nil
}

// ElectricCurrent is a measurement of a flow of electric charge stored as an
// int64 nano Ampere.
//
// This is one of the base unit in the International System of Units.
//
// The highest representable value is 9.2GA.
type ElectricCurrent int64

// String returns the current formatted as a string in Ampere.
func (e ElectricCurrent) String() string {
	return nanoAsString(int64(e)) + "A"
}

// Set takes a string and tries to return a valid ElectricCurrent.
func (e *ElectricCurrent) Set(s string) error {
	v, n, err := setInt(s, nano)
	if err != nil {
		return err
	}

	switch s[n:] {
	case "A", "a", "amp", "amps", "Amp", "Amps":
		*e = (ElectricCurrent)(v)
	default:
		return noUnits("A")
	}
	return nil
}

const (
	NanoAmpere  ElectricCurrent = 1
	MicroAmpere ElectricCurrent = 1000 * NanoAmpere
	MilliAmpere ElectricCurrent = 1000 * MicroAmpere
	Ampere      ElectricCurrent = 1000 * MilliAmpere
	KiloAmpere  ElectricCurrent = 1000 * Ampere
	MegaAmpere  ElectricCurrent = 1000 * KiloAmpere
	GigaAmpere  ElectricCurrent = 1000 * MegaAmpere
)

// ElectricPotential is a measurement of electric potential stored as an int64
// nano Volt.
//
// The highest representable value is 9.2GV.
type ElectricPotential int64

// String returns the tension formatted as a string in Volt.
func (e ElectricPotential) String() string {
	return nanoAsString(int64(e)) + "V"
}

// Set takes a string and tries to return a valid ElectricPotential.
func (e *ElectricPotential) Set(s string) error {
	v, n, err := setInt(s, nano)
	if err != nil {
		return err
	}

	switch s[n:] {
	case "V", "v", "volt", "Volt", "volts", "Volts":
		*e = (ElectricPotential)(v)
	default:
		return noUnits("V")
	}
	return nil
}

const (
	// Volt is W/A, kg⋅m²/s³/A.
	NanoVolt  ElectricPotential = 1
	MicroVolt ElectricPotential = 1000 * NanoVolt
	MilliVolt ElectricPotential = 1000 * MicroVolt
	Volt      ElectricPotential = 1000 * MilliVolt
	KiloVolt  ElectricPotential = 1000 * Volt
	MegaVolt  ElectricPotential = 1000 * KiloVolt
	GigaVolt  ElectricPotential = 1000 * MegaVolt
)

// ElectricResistance is a measurement of the difficulty to pass an electric
// current through a conductor stored as an int64 nano Ohm.
//
// The highest representable value is 9.2GΩ.
type ElectricResistance int64

// String returns the resistance formatted as a string in Ohm.
func (e ElectricResistance) String() string {
	return nanoAsString(int64(e)) + "Ω"
}

// Set takes a string and tries to return a valid ElectricResistance.
func (e *ElectricResistance) Set(s string) error {
	v, n, err := setInt(s, nano)
	if err != nil {
		return err
	}

	switch s[n:] {
	case "Ohm", "Ohms", "ohm", "ohms", "Ω":
		*e = (ElectricResistance)(v)
	default:
		return noUnits("Ohm")
	}
	return nil
}

const (
	// Ohm is V/A, kg⋅m²/s³/A².
	NanoOhm  ElectricResistance = 1
	MicroOhm ElectricResistance = 1000 * NanoOhm
	MilliOhm ElectricResistance = 1000 * MicroOhm
	Ohm      ElectricResistance = 1000 * MilliOhm
	KiloOhm  ElectricResistance = 1000 * Ohm
	MegaOhm  ElectricResistance = 1000 * KiloOhm
	GigaOhm  ElectricResistance = 1000 * MegaOhm
)

// Force is a measurement of interaction that will change the motion of an
// object stored as an int64 nano Newton.
//
// A measurement of Force is a vector and has a direction but this unit only
// represents the magnitude. The orientation needs to be stored as a Quaternion
// independently.
//
// The highest representable value is 9.2TN.
type Force int64

// String returns the force formatted as a string in Newton.
func (f Force) String() string {
	return nanoAsString(int64(f)) + "N"
}

// Set takes a string and tries to return a valid Force.
func (f *Force) Set(s string) error {
	d, n, err := atod(s)
	if err != nil {
		return err
	}
	prefix := prefix(none)
	if !(n == len(s)) {
		var n1 int
		prefix, n1 = parseSIPrefix([]rune(s[n:])[0])
		if prefix == nano && !(s[n:] == "nN") {
			prefix = none
		}
		if prefix != none {
			n += n1
		}
	}
	v, err := d.dtoi(int(prefix - nano))
	if err != nil {
		return err
	}

	switch s[n:] {
	case "N", "Newton", "newton", "Newtons", "newtons":
		*f = (Force)(v)
	default:
		return noUnits("N")
	}
	return nil
}

const (
	// Newton is kg⋅m/s².
	NanoNewton  Force = 1
	MicroNewton Force = 1000 * NanoNewton
	MilliNewton Force = 1000 * MicroNewton
	Newton      Force = 1000 * MilliNewton
	KiloNewton  Force = 1000 * Newton
	MegaNewton  Force = 1000 * KiloNewton
	GigaNewton  Force = 1000 * MegaNewton

	EarthGravity Force = 9806650 * MicroNewton

	// Conversion between Newton and imperial units.
	// Pound is both a unit of mass and weight (force). The suffix Force is added
	// to disambiguate the measurement it represents.
	PoundForce Force = 4448221615261 * NanoNewton
)

// Frequency is a measurement of cycle per second, stored as an int32 micro
// Hertz.
//
// The highest representable value is 9.2THz.
type Frequency int64

// String returns the frequency formatted as a string in Hertz.
func (f Frequency) String() string {
	return microAsString(int64(f)) + "Hz"
}

// Set takes a string and tries to return a valid Frequency.

func (f *Frequency) Set(s string) error {
	v, n, err := setInt(s, micro)
	if err != nil {
		return err
	}

	switch s[n:] {
	case "Hz", "hz":
		*f = (Frequency)(v)
	default:
		return noUnits("Hz")
	}
	return nil
}

// Duration returns the duration of one cycle at this frequency.
func (f Frequency) Duration() time.Duration {
	// Note: Duration() should have been named Period().
	// TODO(maruel): Rounding should be fine-tuned.
	return time.Second * time.Duration(Hertz) / time.Duration(f)
}

//ParseFrequency takes a string and returns a frequency. Formating of string is
// 10MHz, 100µHz etc.
func ParseFrequency(s string) (Frequency, error) {
	var f Frequency
	err := f.Set(s)
	return f, err
}

// PeriodToFrequency returns the frequency for a period of this interval.
func PeriodToFrequency(t time.Duration) Frequency {
	return Frequency(time.Second) * Hertz / Frequency(t)
}

const (
	// Hertz is 1/s.
	MicroHertz Frequency = 1
	MilliHertz Frequency = 1000 * MicroHertz
	Hertz      Frequency = 1000 * MilliHertz
	KiloHertz  Frequency = 1000 * Hertz
	MegaHertz  Frequency = 1000 * KiloHertz
	GigaHertz  Frequency = 1000 * MegaHertz
	TeraHertz  Frequency = 1000 * GigaHertz
)

// Mass is a measurement of mass stored as an int64 nano gram.
//
// This is one of the base unit in the International System of Units.
//
// The highest representable value is 9.2Gg.
type Mass int64

// String returns the mass formatted as a string in gram.
func (m Mass) String() string {
	return nanoAsString(int64(m)) + "g"
}

// Set takes a string and tries to return a valid Mass.
func (m *Mass) Set(s string) error {
	d, n, err := atod(s)
	if err != nil {
		return err
	}
	prefix := prefix(none)
	if !(n == len(s)) {
		var n1 int
		prefix, n1 = parseSIPrefix([]rune(s[n:])[0])
		if prefix == giga && !(s[n:] == "Gg") {
			prefix = none
		}
		if prefix == tera && !(s[n:] == "Tg") {
			prefix = none
		}
		if prefix != none {
			n += n1
		}
	}
	v, err := d.dtoi(int(prefix - nano))
	if err != nil {
		return err
	}
	// fmt.Println(s, n, v)
	// fmt.Println("into switch ", s[n:])
	switch s[n:] {
	case "tonne", "Tonne", "tonnes", "Tonnes":
		*m = (Mass)(v * 1000000)
	case "pound", "Pound", "pounds", "Pounds", "lb":
		*m = (Mass)((v * 45359237) / 100000)
	case "ounce", "Ounce", "ounces", "Ounces", "oz", "Oz":
		*m = (Mass)(v / 1000000000 * 28349523125)
	case "g", "gram", "grams", "Gram", "Grams":
		*m = (Mass)(v)
	default:
		return noUnits("g")
	}
	return nil
}

const (
	NanoGram  Mass = 1
	MicroGram Mass = 1000 * NanoGram
	MilliGram Mass = 1000 * MicroGram
	Gram      Mass = 1000 * MilliGram
	KiloGram  Mass = 1000 * Gram
	MegaGram  Mass = 1000 * KiloGram
	GigaGram  Mass = 1000 * MegaGram
	Tonne     Mass = MegaGram

	// Conversion between Gram and imperial units.
	// Ounce is both a unit of mass, weight (force) or volume depending on
	// context. The suffix Mass is added to disambiguate the measurement it
	// represents.
	OunceMass Mass = 28349523125 * NanoGram
	// Pound is both a unit of mass and weight (force). The suffix Mass is added
	// to disambiguate the measurement it represents.
	PoundMass Mass = 16 * OunceMass
	Slug      Mass = 14593903 * MilliGram
)

// Pressure is a measurement of force applied to a surface per unit
// area (stress) stored as an int64 nano Pascal.
//
// The highest representable value is 9.2GPa.
type Pressure int64

// String returns the pressure formatted as a string in Pascal.
func (p Pressure) String() string {
	return nanoAsString(int64(p)) + "Pa"
}

// Set takes a string and tries to return a valid Pressure.
func (p *Pressure) Set(s string) error {
	v, n, err := setInt(s, nano)
	if err != nil {
		return err
	}

	switch s[n:] {
	case "Pa", "pa", "Pascal", "pascal", "Pascals", "pascals":
		*p = (Pressure)(v)
	default:
		return noUnits("Pa")
	}
	return nil
}

const (
	// Pascal is N/m², kg/m/s².
	NanoPascal  Pressure = 1
	MicroPascal Pressure = 1000 * NanoPascal
	MilliPascal Pressure = 1000 * MicroPascal
	Pascal      Pressure = 1000 * MilliPascal
	KiloPascal  Pressure = 1000 * Pascal
	MegaPascal  Pressure = 1000 * KiloPascal
	GigaPascal  Pressure = 1000 * MegaPascal
)

// RelativeHumidity is a humidity level measurement stored as an int32 fixed
// point integer at a precision of 0.00001%rH.
//
// Valid values are between 0% and 100%.
type RelativeHumidity int32

// String returns the humidity formatted as a string.
func (r RelativeHumidity) String() string {
	r /= MilliRH
	frac := int(r % 10)
	if frac == 0 {
		return strconv.Itoa(int(r)/10) + "%rH"
	}
	if frac < 0 {
		frac = -frac
	}
	return strconv.Itoa(int(r)/10) + "." + strconv.Itoa(frac) + "%rH"
}

// Set takes a string and tries to return a valid RelativeHumidity.
func (r *RelativeHumidity) Set(s string) error {
	v, n, err := setInt(s, prefix(-7))
	if err != nil {
		return err
	}
	switch s[n:] {
	case "rH", "rh":
		*r = (RelativeHumidity)(v)
	case "%rH", "%rh":
		*r = (RelativeHumidity)(v / 100)
	default:
		return noUnits("%rH")
	}
	return nil
}

const (
	TenthMicroRH RelativeHumidity = 1                 // 0.00001%rH
	MicroRH      RelativeHumidity = 10 * TenthMicroRH // 0.0001%rH
	MilliRH      RelativeHumidity = 1000 * MicroRH    // 0.1%rH
	PercentRH    RelativeHumidity = 10 * MilliRH      // 1%rH
)

// Speed is a measurement of magnitude of velocity stored as an int64 nano
// Metre per Second.
//
// The highest representable value is 9.2Gm/s.
type Speed int64

// String returns the speed formatted as a string in m/s.
func (s Speed) String() string {
	return nanoAsString(int64(s)) + "m/s"
}

// Set takes a string and tries to return a valid Speed.
func (s *Speed) Set(str string) error {
	d, n, err := atod(str)
	if err != nil {
		return err
	}
	prefix := prefix(none)
	if !(n == len(str)) {
		var n1 int
		prefix, n1 = parseSIPrefix([]rune(str[n:])[0])
		if prefix == milli && !(str[n:] == "mm/s") {
			prefix = none
		}
		if prefix == kilo && !(str[n:] == "km/s") {
			prefix = none
		}
		if prefix != none {
			n += n1
		}
	}
	v, err := d.dtoi(int(prefix - nano))
	if err != nil {
		return err
	}

	switch str[n:] {
	case "fps":
		*s = (Speed)((v / 1000000) * 304800)
	case "mph":
		*s = (Speed)((v / 1000000) * 447040)
	case "km/h":
		*s = (Speed)(v * 10 / 36)
	case "m/s":
		*s = (Speed)(v)
	default:
		return noUnits("m/s")
	}
	return nil
}

const (
	// MetrePerSecond is m/s.
	NanoMetrePerSecond  Speed = 1
	MicroMetrePerSecond Speed = 1000 * NanoMetrePerSecond
	MilliMetrePerSecond Speed = 1000 * MicroMetrePerSecond
	MetrePerSecond      Speed = 1000 * MilliMetrePerSecond
	KiloMetrePerSecond  Speed = 1000 * MetrePerSecond
	MegaMetrePerSecond  Speed = 1000 * KiloMetrePerSecond
	GigaMetrePerSecond  Speed = 1000 * MegaMetrePerSecond

	LightSpeed Speed = 299792458 * MetrePerSecond

	KilometrePerHour Speed = 277777778 * NanoMetrePerSecond
	MilePerHour      Speed = 447040 * MicroMetrePerSecond
	FootPerSecond    Speed = 304800 * MicroMetrePerSecond
)

// Temperature is a measurement of hotness stored as a nano kelvin.
//
// Negative values are invalid.
//
// The highest representable value is 9.2GK.
type Temperature int64

// String returns the temperature formatted as a string in °Celsius.
func (t Temperature) String() string {
	return nanoAsString(int64(t-ZeroCelsius)) + "°C"
}

// Set takes a string and tries to return a valid Temperature.
func (t *Temperature) Set(s string) error {
	v, n, err := setInt(s, nano)
	if err != nil {
		return err
	}
	switch s[n:] {
	//TODO(neuralspaz) Fahrenheit
	case "K":
		*t = (Temperature)(v)
	case "C", "°C":
		*t = (Temperature)(v + int64(ZeroCelsius))
	default:
		return noUnits("K or C")
	}
	return nil
}

const (
	NanoKelvin  Temperature = 1
	MicroKelvin Temperature = 1000 * NanoKelvin
	MilliKelvin Temperature = 1000 * MicroKelvin
	Kelvin      Temperature = 1000 * MilliKelvin
	KiloKelvin  Temperature = 1000 * Kelvin
	MegaKelvin  Temperature = 1000 * KiloKelvin
	GigaKelvin  Temperature = 1000 * MegaKelvin

	// Conversion between Kelvin and Celsius.
	ZeroCelsius  Temperature = 273150 * MilliKelvin
	MilliCelsius Temperature = MilliKelvin
	Celsius      Temperature = Kelvin

	// Conversion between Kelvin and Fahrenheit.
	ZeroFahrenheit  Temperature = 255372 * MilliKelvin
	MilliFahrenheit Temperature = 555555 * NanoKelvin
	Fahrenheit      Temperature = 555555555 * NanoKelvin
)

// Power is a measurement of power stored as a nano watts.
//
// The highest representable value is 9.2GW.
type Power int64

// String returns the power formatted as a string in watts.
func (p Power) String() string {
	return nanoAsString(int64(p)) + "W"
}

// Set takes a string and tries to return a valid Power.
func (p *Power) Set(s string) error {
	v, n, err := setInt(s, nano)
	if err != nil {
		return err
	}

	switch s[n:] {
	case "watt", "watts", "Watt", "Watts", "W", "w":
		*p = (Power)(v)
	default:
		return noUnits("W")
	}
	return nil
}

const (
	// Watt is unit of power J/s, kg⋅m²⋅s⁻³
	NanoWatt  Power = 1
	MicroWatt Power = 1000 * NanoWatt
	MilliWatt Power = 1000 * MicroWatt
	Watt      Power = 1000 * MilliWatt
	KiloWatt  Power = 1000 * Watt
	MegaWatt  Power = 1000 * KiloWatt
	GigaWatt  Power = 1000 * MegaWatt
)

// Energy is a measurement of work stored as a nano joules.
//
// The highest representable value is 9.2GJ.
type Energy int64

// String returns the energy formatted as a string in Joules.
func (e Energy) String() string {
	return nanoAsString(int64(e)) + "J"
}

// Set takes a string and tries to return a valid Energy.
func (e *Energy) Set(s string) error {
	v, n, err := setInt(s, nano)
	if err != nil {
		return err
	}

	switch s[n:] {
	case "Joule", "Joules", "joule", "joules", "J", "j":
		*e = (Energy)(v)
	default:
		return noUnits("J")
	}
	return nil
}

const (
	// Joule is a unit of work. kg⋅m²⋅s⁻²
	NanoJoule  Energy = 1
	MicroJoule Energy = 1000 * NanoJoule
	MilliJoule Energy = 1000 * MicroJoule
	Joule      Energy = 1000 * MilliJoule
	KiloJoule  Energy = 1000 * Joule
	MegaJoule  Energy = 1000 * KiloJoule
	GigaJoule  Energy = 1000 * MegaJoule
)

// ElectricalCapacitance is a measurement of capacitance stored as a pico farad.
//
// The highest representable value is 9.2MF.
type ElectricalCapacitance int64

// String returns the energy formatted as a string in Farad.
func (c ElectricalCapacitance) String() string {
	return picoAsString(int64(c)) + "F"
}

// Set takes a string and tries to return a valid ElectricalCapacitance.
func (e *ElectricalCapacitance) Set(s string) error {
	v, n, err := setInt(s, pico)
	if err != nil {
		return err
	}

	switch s[n:] {
	case "f", "farad", "farads", "F", "Farad", "Farads":
		*e = (ElectricalCapacitance)(v)
	default:
		return noUnits("F")
	}
	return nil
}

const (
	// Farad is a unit of capacitance. kg⁻¹⋅m⁻²⋅s⁴A²
	PicoFarad  ElectricalCapacitance = 1
	NanoFarad  ElectricalCapacitance = 1000 * PicoFarad
	MicroFarad ElectricalCapacitance = 1000 * NanoFarad
	MilliFarad ElectricalCapacitance = 1000 * MicroFarad
	Farad      ElectricalCapacitance = 1000 * MilliFarad
	KiloFarad  ElectricalCapacitance = 1000 * Farad
	MegaFarad  ElectricalCapacitance = 1000 * KiloFarad
)

// LuminousIntensity is a measurement of the quantity of visible light energy
// emitted per unit solid angle with wavelength power weighted by a luminosity
// function which represents the human eye's response to different wavelengths.
// The CIE 1931 luminosity function is the SI standard for candela.
//
// LuminousIntensity is stored as nano candela.
//
// This is one of the base unit in the International System of Units.
//
// The highest representable value is 9.2Gcd.
type LuminousIntensity int64

// String returns the energy formatted as a string in Candela.
func (l LuminousIntensity) String() string {
	return nanoAsString(int64(l)) + "cd"
}

// Set takes a string and tries to return a valid LuminousIntensity.
func (l *LuminousIntensity) Set(s string) error {
	v, n, err := setInt(s, nano)
	if err != nil {
		return err
	}

	switch s[n:] {
	case "cd", "Candela", "candela", "Candelas", "candelas":
		*l = (LuminousIntensity)(v)
	default:
		return noUnits("cd")
	}
	return nil
}

const (
	// Candela is a unit of luminous intensity. cd
	NanoCandela  LuminousIntensity = 1
	MicroCandela LuminousIntensity = 1000 * NanoCandela
	MilliCandela LuminousIntensity = 1000 * MicroCandela
	Candela      LuminousIntensity = 1000 * MilliCandela
	KiloCandela  LuminousIntensity = 1000 * Candela
	MegaCandela  LuminousIntensity = 1000 * KiloCandela
	GigaCandela  LuminousIntensity = 1000 * MegaCandela
)

// LuminousFlux is a measurement of total quantity of visible light energy
// emitted with wavelength power weighted by a luminosity function which
// represents a model of the human eye's response to different wavelengths.
// The CIE 1931 luminosity function is the standard for lumens.
//
// LuminousFlux is stored as nano lumens.
//
// The highest representable value is 9.2Glm.
type LuminousFlux int64

// String returns the energy formatted as a string in Lumens.
func (f LuminousFlux) String() string {
	return nanoAsString(int64(f)) + "lm"
}

// Set takes a string and tries to return a valid LuminousFlux.
func (l *LuminousFlux) Set(s string) error {
	v, n, err := setInt(s, nano)
	if err != nil {
		return err
	}

	switch s[n:] {
	case "lm", "Lumen", "lumen", "Lumens", "lumens":
		*l = (LuminousFlux)(v)
	default:
		return noUnits("lm")
	}
	return nil
}

const (
	// Lumen is a unit of luminous flux. cd⋅sr
	NanoLumen  LuminousFlux = 1
	MicroLumen LuminousFlux = 1000 * NanoLumen
	MilliLumen LuminousFlux = 1000 * MicroLumen
	Lumen      LuminousFlux = 1000 * MilliLumen
	KiloLumen  LuminousFlux = 1000 * Lumen
	MegaLumen  LuminousFlux = 1000 * KiloLumen
	GigaLumen  LuminousFlux = 1000 * MegaLumen
)

//

func prefixZeros(digits, v int) string {
	// digits is expected to be around 2~3.
	s := strconv.Itoa(v)
	for len(s) < digits {
		// O(n²) but since digits is expected to run 2~3 times at most, it doesn't
		// matter.
		s = "0" + s
	}
	return s
}

// nanoAsString converts a value in S.I. unit in a string with the predefined
// prefix.
func nanoAsString(v int64) string {
	sign := ""
	if v < 0 {
		if v == -9223372036854775808 {
			v++
		}
		sign = "-"
		v = -v
	}
	var frac int
	var base int
	var precision int64
	unit := ""
	switch {
	case v >= 999999500000000001:
		precision = v % 1000000000000000
		base = int(v / 1000000000000000)
		if precision > 500000000000000 {
			base++
		}
		frac = (base % 1000)
		base = base / 1000
		unit = "G"
	case v >= 999999500000001:
		precision = v % 1000000000000
		base = int(v / 1000000000000)
		if precision > 500000000000 {
			base++
		}
		frac = (base % 1000)
		base = base / 1000
		unit = "M"
	case v >= 999999500001:
		precision = v % 1000000000
		base = int(v / 1000000000)
		if precision > 500000000 {
			base++
		}
		frac = (base % 1000)
		base = base / 1000
		unit = "k"
	case v >= 999999501:
		precision = v % 1000000
		base = int(v / 1000000)
		if precision > 500000 {
			base++
		}
		frac = (base % 1000)
		base = base / 1000
		unit = ""
	case v >= 1000000:
		precision = v % 1000
		base = int(v / 1000)
		if precision > 500 {
			base++
		}
		frac = (base % 1000)
		base = base / 1000
		unit = "m"
	case v >= 1000:
		frac = int(v) % 1000
		base = int(v) / 1000
		unit = "µ"
	default:
		if v == 0 {
			return "0"
		}
		base = int(v)
		unit = "n"
	}
	if frac == 0 {
		return sign + strconv.Itoa(base) + unit
	}
	return sign + strconv.Itoa(base) + "." + prefixZeros(3, frac) + unit
}

// microAsString converts a value in S.I. unit in a string with the predefined
// prefix.
func microAsString(v int64) string {
	sign := ""
	if v < 0 {
		if v == -9223372036854775808 {
			v++
		}
		sign = "-"
		v = -v
	}
	var frac int
	var base int
	var precision int64
	unit := ""
	switch {
	case v >= 999999500000000001:
		precision = v % 1000000000000000
		base = int(v / 1000000000000000)
		if precision > 500000000000000 {
			base++
		}
		frac = (base % 1000)
		base = base / 1000
		unit = "T"
	case v >= 999999500000001:
		precision = v % 1000000000000
		base = int(v / 1000000000000)
		if precision > 500000000000 {
			base++
		}
		frac = (base % 1000)
		base = base / 1000
		unit = "G"
	case v >= 999999500001:
		precision = v % 1000000000
		base = int(v / 1000000000)
		if precision > 500000000 {
			base++
		}
		frac = (base % 1000)
		base = base / 1000
		unit = "M"
	case v >= 999999501:
		precision = v % 1000000
		base = int(v / 1000000)
		if precision > 500000 {
			base++
		}
		frac = (base % 1000)
		base = base / 1000
		unit = "k"
	case v >= 1000000:
		precision = v % 1000
		base = int(v / 1000)
		if precision > 500 {
			base++
		}
		frac = (base % 1000)
		base = base / 1000
		unit = ""
	case v >= 1000:
		frac = int(v) % 1000
		base = int(v) / 1000
		unit = "m"
	default:
		if v == 0 {
			return "0"
		}
		base = int(v)
		unit = "µ"
	}
	if frac == 0 {
		return sign + strconv.Itoa(base) + unit
	}
	return sign + strconv.Itoa(base) + "." + prefixZeros(3, frac) + unit
}

// picoAsString converts a value in S.I. unit in a string with the predefined
// prefix.
func picoAsString(v int64) string {
	sign := ""
	if v < 0 {
		if v == -9223372036854775808 {
			v++
		}
		sign = "-"
		v = -v
	}
	var frac int
	var base int
	var precision int64
	unit := ""
	switch {
	case v >= 999999500000000001:
		precision = v % 1000000000000000
		base = int(v / 1000000000000000)
		if precision > 500000000000000 {
			base++
		}
		frac = (base % 1000)
		base = base / 1000
		unit = "M"
	case v >= 999999500000001:
		precision = v % 1000000000000
		base = int(v / 1000000000000)
		if precision > 500000000000 {
			base++
		}
		frac = (base % 1000)
		base = base / 1000
		unit = "k"
	case v >= 999999500001:
		precision = v % 1000000000
		base = int(v / 1000000000)
		if precision > 500000000 {
			base++
		}
		frac = (base % 1000)
		base = base / 1000
		unit = ""
	case v >= 999999501:
		precision = v % 1000000
		base = int(v / 1000000)
		if precision > 500000 {
			base += 1
		}
		frac = (base % 1000)
		base = base / 1000
		unit = "m"
	case v >= 1000000:
		precision = v % 1000
		base = int(v / 1000)
		if precision > 500 {
			base++
		}
		frac = (base % 1000)
		base = base / 1000
		unit = "µ"
	case v >= 1000:
		frac = int(v) % 1000
		base = int(v) / 1000
		unit = "n"
	default:
		if v == 0 {
			return "0"
		}
		base = int(v)
		unit = "p"
	}
	if frac == 0 {
		return sign + strconv.Itoa(base) + unit
	}
	return sign + strconv.Itoa(base) + "." + prefixZeros(3, frac) + unit
}

// decimal representation of number.
type decimal struct {
	digits string
	exp    int
	neg    bool
}

// Positive powers of 10 in the form such that powerOF10[index] = 10^index.
var powerOf10 = [19]uint64{
	1,
	10,
	100,
	1000,
	10000,
	100000,
	1000000,
	10000000,
	100000000,
	1000000000,
	10000000000,
	100000000000,
	1000000000000,
	10000000000000,
	100000000000000,
	1000000000000000,
	10000000000000000,
	100000000000000000,
	1000000000000000000,
}

func (d *decimal) dtoi(scale int) (int64, error) {
	const max = (1<<63 - 1)
	// Use uint till the last as it allows checks for overflows.
	var u uint64
	for _, c := range []byte(d.digits) {
		// Check that is is a digit.
		if c >= '0' && c <= '9' {
			digit := c - '0'
			u *= 10
			check := u + uint64(digit)
			if check < u || check > max {
				return int64(u), &parseError{err: errors.New("overflows")}
			}
			u = check
		} else {
			// Should really not get here and just for safety.
			return 0, &parseError{err: errors.New("not a number")}
		}
	}

	// Get the total magnitude of the number
	mag := d.exp + scale
	if mag > 18 {
		return 0, errors.New("exceeds maximum exponent")
	}
	if mag < 0 {
		mag *= -1
	}
	// time.Sleep(time.Second)
	if d.exp+scale < 0 {
		u /= powerOf10[mag]
	} else {
		check := u * powerOf10[mag]
		if check < u || check > max {
			return max, &parseError{err: errors.New("overflows")}
		}
		u *= powerOf10[mag]
	}

	n := int64(u)
	if d.neg {
		n *= -1
	}
	return n, nil
}

func atod(s string) (decimal, int, error) {
	var d decimal
	start := 0
	dp := 0
	end := 0
	seenDigit := false
	seenZero := false
	isPoint := false

	// Strip leading zeros, +/- and mark DP.
	for i := 0; i < len(s); i++ {
		switch {
		case s[i] == '-':
			d.neg = true
			start++
		case s[i] == '+':
			start++
		case s[i] == '.':
			if isPoint {
				return d, 0, &parseError{s, i, errors.New("multiple decimal points")}
			}
			isPoint = true
			dp = i
			if !seenDigit {
				start++
			}
		case s[i] == '0':
			if !seenDigit {
				start++
			}
			seenZero = true
		case s[i] >= '1' && s[i] <= '9':
			seenDigit = true
		default:
			if !seenDigit && !seenZero {
				return d, 0, &parseError{s: s, err: errors.New("is not a number")}
			}
			end = i
			break
		}
	}

	if end == 0 {
		end = len(s)
	}
	last := end
	seenDigit = false
	exp := 0
	// Strip non significant zeros.
	for i := end - 1; i > start-1; i-- {
		switch {
		case s[i] >= '1' && s[i] <= '9':
			seenDigit = true
		case s[i] == '.':
			if !seenDigit {
				end--
			}
		case s[i] == '0':
			if !seenDigit {
				if i > dp {
					end--
				}
				if i <= dp || dp == 0 {
					exp++
				}
			}
		default:
			last--
			end--
		}
	}

	if dp > start && dp < end {
		//Concatenate with out decimal point
		d.digits = s[start:dp] + s[dp+1:end]
	} else {
		d.digits = s[start:end]
	}
	if !isPoint {
		d.exp = exp
	} else {
		ttl := dp - start
		length := len(d.digits)
		if ttl > 0 {
			d.exp = ttl - length
		} else {
			d.exp = ttl - length + 1
		}
	}
	return d, last, nil
}

func setInt(s string, base prefix) (int64, int, error) {
	d, n, err := atod(s)
	if err != nil {
		return 0, n, err
	}
	si := prefix(none)
	if !(n == len(s)) {
		var n1 int
		si, n1 = parseSIPrefix([]rune(s[n:])[0])
		if si != none {
			n += n1
		}
	}
	v, err := d.dtoi(int(si - base))
	if err != nil {
		return v, 0, err
	}
	return v, n, nil
}

type parseError struct {
	s        string
	position int
	err      error
}

func (p *parseError) Error() string {
	return fmt.Sprintf("parse error: %v: \"%s\"", p.err, p.s)
}

func noUnits(s string) error {
	return &parseError{s: s, err: errors.New("no units provided, need")}
}

type prefix int

const (
	pico  prefix = -12
	nano  prefix = -9
	micro prefix = -6
	milli prefix = -3
	none  prefix = 0
	deca  prefix = 1
	hecto prefix = 2
	kilo  prefix = 3
	mega  prefix = 6
	giga  prefix = 9
	tera  prefix = 12
)

func parseSIPrefix(r rune) (prefix, int) {
	switch r {
	case 'p':
		return pico, len(string(r))
	case 'n':
		return nano, len(string(r))
	case 'u':
		return micro, len(string(r))
	case 'µ':
		return micro, len(string(r))
	case 'm':
		return milli, len(string(r))
	case 'k':
		return kilo, len(string(r))
	case 'M':
		return mega, len(string(r))
	case 'G':
		return giga, len(string(r))
	case 'T':
		return tera, len(string(r))
	default:
		return none, 0
	}
}
