// Copyright 2018 The Periph Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

package physic

import (
	"bytes"
	"flag"
	"fmt"
	"testing"
	"time"
)

func TestAngle_String(t *testing.T) {
	data := []struct {
		in       Angle
		expected string
	}{
		{0, "0°"},
		{Degree/10000 + Degree/2000, "0.001°"},
		{-Degree/10000 - Degree/2000, "-0.001°"},
		{Degree / 1000, "0.001°"},
		{-Degree / 1000, "-0.001°"},
		{Degree / 2, "0.500°"},
		{-Degree / 2, "-0.500°"},
		{Degree, "1.000°"},
		{-Degree, "-1.000°"},
		{10 * Degree, "10.00°"},
		{-10 * Degree, "-10.00°"},
		{100 * Degree, "100.0°"},
		{-100 * Degree, "-100.0°"},
		{1000 * Degree, "1000°"},
		{-1000 * Degree, "-1000°"},
		{100000000000 * Degree, "100000000000°"},
		{-100000000000 * Degree, "-100000000000°"},
		{(9223372036854775807 - 17453293) * NanoRadian, "528460276054°"},
		{(-9223372036854775807 + 17453293) * NanoRadian, "-528460276054°"},
		{Pi, "180.0°"},
		{Theta, "360.0°"},
		{Radian, "57.296°"},
	}
	for i, line := range data {
		if s := line.in.String(); s != line.expected {
			t.Fatalf("%d: Degree(%d).String() = %s != %s", i, int64(line.in), s, line.expected)
		}
	}
}

func TestDistance_String(t *testing.T) {
	if s := Mile.String(); s != "1.609km" {
		t.Fatalf("%#v", s)
	}
}

func TestElectricCurrent_String(t *testing.T) {
	if s := Ampere.String(); s != "1A" {
		t.Fatalf("%#v", s)
	}
}

func TestElectricPotential_String(t *testing.T) {
	if s := Volt.String(); s != "1V" {
		t.Fatalf("%#v", s)
	}
}

func TestElectricResistance_String(t *testing.T) {
	if s := Ohm.String(); s != "1Ω" {
		t.Fatalf("%#v", s)
	}
}

func TestForce_String(t *testing.T) {
	if s := Newton.String(); s != "1N" {
		t.Fatalf("%#v", s)
	}
}

func TestFrequency_String(t *testing.T) {
	if s := Hertz.String(); s != "1Hz" {
		t.Fatalf("%#v", s)
	}
}

func TestFrequency_Duration(t *testing.T) {
	if v := MegaHertz.Duration(); v != time.Microsecond {
		t.Fatalf("%#v", v)
	}
}

func TestFrequency_PeriodToFrequency(t *testing.T) {
	if v := PeriodToFrequency(time.Millisecond); v != KiloHertz {
		t.Fatalf("%#v", v)
	}
}

func TestMass_String(t *testing.T) {
	if s := PoundMass.String(); s != "453.592g" {
		t.Fatalf("%#v", s)
	}
}

func TestPressure_String(t *testing.T) {
	if s := NanoPascal.String(); s != "1nPa" {
		t.Fatalf("%v", s)
	}
	if s := MicroPascal.String(); s != "1µPa" {
		t.Fatalf("%v", s)
	}
	if s := MilliPascal.String(); s != "1mPa" {
		t.Fatalf("%v", s)
	}
	if s := Pascal.String(); s != "1Pa" {
		t.Fatalf("%v", s)
	}
	if s := KiloPascal.String(); s != "1kPa" {
		t.Fatalf("%v", s)
	}
	if s := MegaPascal.String(); s != "1MPa" {
		t.Fatalf("%v", s)
	}
	if s := GigaPascal.String(); s != "1GPa" {
		t.Fatalf("%v", s)
	}

}

func TestRelativeHumidity_String(t *testing.T) {
	data := []struct {
		in       RelativeHumidity
		expected string
	}{
		{TenthMicroRH, "0%rH"},
		{MicroRH, "0%rH"},
		{10 * MicroRH, "0%rH"},
		{100 * MicroRH, "0%rH"},
		{1000 * MicroRH, "0.1%rH"},
		{506000 * MicroRH, "50.6%rH"},
		{90 * PercentRH, "90%rH"},
		{100 * PercentRH, "100%rH"},
		// That's a lot of humidity. This is to test the value doesn't overflow
		// int32 too quickly.
		{1000 * PercentRH, "1000%rH"},
		// That's really dry.
		{-501000 * MicroRH, "-50.1%rH"},
	}
	for i, line := range data {
		if s := line.in.String(); s != line.expected {
			t.Fatalf("%d: RelativeHumidity(%d).String() = %s != %s", i, int64(line.in), s, line.expected)
		}
	}
}

func TestSpeed_String(t *testing.T) {
	if s := MilePerHour.String(); s != "447.040mm/s" {
		t.Fatalf("%#v", s)
	}
}

func TestTemperature_String(t *testing.T) {
	if s := ZeroCelsius.String(); s != "0°C" {
		t.Fatalf("%#v", s)
	}
	if s := Temperature(0).String(); s != "-273.150°C" {
		t.Fatalf("%#v", s)
	}
}

func TestPower_String(t *testing.T) {
	if s := NanoWatt.String(); s != "1nW" {
		t.Fatalf("%v", s)
	}
	if s := MicroWatt.String(); s != "1µW" {
		t.Fatalf("%v", s)
	}
	if s := MilliWatt.String(); s != "1mW" {
		t.Fatalf("%v", s)
	}
	if s := Watt.String(); s != "1W" {
		t.Fatalf("%v", s)
	}
	if s := KiloWatt.String(); s != "1kW" {
		t.Fatalf("%v", s)
	}
	if s := MegaWatt.String(); s != "1MW" {
		t.Fatalf("%v", s)
	}
	if s := GigaWatt.String(); s != "1GW" {
		t.Fatalf("%v", s)
	}
}
func TestEnergy_String(t *testing.T) {
	if s := NanoJoule.String(); s != "1nJ" {
		t.Fatalf("%v", s)
	}
	if s := MicroJoule.String(); s != "1µJ" {
		t.Fatalf("%v", s)
	}
	if s := MilliJoule.String(); s != "1mJ" {
		t.Fatalf("%v", s)
	}
	if s := Joule.String(); s != "1J" {
		t.Fatalf("%v", s)
	}
	if s := KiloJoule.String(); s != "1kJ" {
		t.Fatalf("%v", s)
	}
	if s := MegaJoule.String(); s != "1MJ" {
		t.Fatalf("%v", s)
	}
	if s := GigaJoule.String(); s != "1GJ" {
		t.Fatalf("%v", s)
	}
}

func TestCapacitance_String(t *testing.T) {
	if s := PicoFarad.String(); s != "1pF" {
		t.Fatalf("%v", s)
	}
	if s := NanoFarad.String(); s != "1nF" {
		t.Fatalf("%v", s)
	}
	if s := MicroFarad.String(); s != "1µF" {
		t.Fatalf("%v", s)
	}
	if s := MilliFarad.String(); s != "1mF" {
		t.Fatalf("%v", s)
	}
	if s := Farad.String(); s != "1F" {
		t.Fatalf("%v", s)
	}
	if s := KiloFarad.String(); s != "1kF" {
		t.Fatalf("%v", s)
	}
	if s := MegaFarad.String(); s != "1MF" {
		t.Fatalf("%v", s)
	}
}

func TestLuminousIntensity_String(t *testing.T) {
	if s := NanoCandela.String(); s != "1ncd" {
		t.Fatalf("%v", s)
	}
	if s := MicroCandela.String(); s != "1µcd" {
		t.Fatalf("%v", s)
	}
	if s := MilliCandela.String(); s != "1mcd" {
		t.Fatalf("%v", s)
	}
	if s := Candela.String(); s != "1cd" {
		t.Fatalf("%v", s)
	}
	if s := KiloCandela.String(); s != "1kcd" {
		t.Fatalf("%v", s)
	}
	if s := MegaCandela.String(); s != "1Mcd" {
		t.Fatalf("%v", s)
	}
	if s := GigaCandela.String(); s != "1Gcd" {
		t.Fatalf("%v", s)
	}
}

func TestFlux_String(t *testing.T) {
	if s := NanoLumen.String(); s != "1nlm" {
		t.Fatalf("%v", s)
	}
	if s := MicroLumen.String(); s != "1µlm" {
		t.Fatalf("%v", s)
	}
	if s := MilliLumen.String(); s != "1mlm" {
		t.Fatalf("%v", s)
	}
	if s := Lumen.String(); s != "1lm" {
		t.Fatalf("%v", s)
	}
	if s := KiloLumen.String(); s != "1klm" {
		t.Fatalf("%v", s)
	}
	if s := MegaLumen.String(); s != "1Mlm" {
		t.Fatalf("%v", s)
	}
	if s := GigaLumen.String(); s != "1Glm" {
		t.Fatalf("%v", s)
	}
}

func TestPicoAsString(t *testing.T) {
	data := []struct {
		in       int64
		expected string
	}{
		{0, "0"}, // 0
		{1, "1p"},
		{-1, "-1p"},
		{900, "900p"},
		{-900, "-900p"},
		{999, "999p"},
		{-999, "-999p"},
		{1000, "1n"},
		{-1000, "-1n"},
		{1100, "1.100n"},
		{-1100, "-1.100n"}, // 10
		{999999, "999.999n"},
		{-999999, "-999.999n"},
		{1000000, "1µ"},
		{-1000000, "-1µ"},
		{1000501, "1.001µ"},
		{-1000501, "-1.001µ"},
		{1100000, "1.100µ"},
		{-1100000, "-1.100µ"},
		{999999501, "1m"},
		{-999999501, "-1m"},
		{999999999, "1m"},
		{-999999999, "-1m"},
		{1000000000, "1m"},
		{-1000000000, "-1m"}, // 20
		{1100000000, "1.100m"},
		{-1100000000, "-1.100m"},
		{999999499999, "999.999m"},
		{-999999499999, "-999.999m"},
		{999999500001, "1"},
		{-999999500001, "-1"},
		{1000000000000, "1"},
		{-1000000000000, "-1"},
		{1100000000000, "1.100"},
		{-1100000000000, "-1.100"},
		{999999499999999, "999.999"},
		{-999999499999999, "-999.999"},
		{999999500000001, "1k"},
		{-999999500000001, "-1k"},
		{1000000000000000, "1k"}, //30
		{-1000000000000000, "-1k"},
		{1100000000000000, "1.100k"},
		{-1100000000000000, "-1.100k"},
		{999999499999999999, "999.999k"},
		{-999999499999999999, "-999.999k"},
		{999999500000000001, "1M"},
		{-999999500000000001, "-1M"},
		{1000000000000000000, "1M"},
		{-1000000000000000000, "-1M"},
		{1100000000000000000, "1.100M"},
		{-1100000000000000000, "-1.100M"},
		{-1999499999999999999, "-1.999M"},
		{1999499999999999999, "1.999M"},
		{-1999500000000000001, "-2M"},
		{1999500000000000001, "2M"},
		{9223372036854775807, "9.223M"},
		{-9223372036854775807, "-9.223M"},
		{-9223372036854775808, "-9.223M"},
	}
	for i, line := range data {
		if s := picoAsString(line.in); s != line.expected {
			t.Fatalf("%d: picoAsString(%d).String() = %s != %s", i, line.in, s, line.expected)
		}
	}
}

func TestNanoAsString(t *testing.T) {
	data := []struct {
		in       int64
		expected string
	}{
		{0, "0"}, // 0
		{1, "1n"},
		{-1, "-1n"},
		{900, "900n"},
		{-900, "-900n"},
		{999, "999n"},
		{-999, "-999n"},
		{1000, "1µ"},
		{-1000, "-1µ"},
		{1100, "1.100µ"},
		{-1100, "-1.100µ"}, // 10
		{999999, "999.999µ"},
		{-999999, "-999.999µ"},
		{1000000, "1m"},
		{-1000000, "-1m"},
		{1100000, "1.100m"},
		{1100100, "1.100m"},
		{1101000, "1.101m"},
		{-1100000, "-1.100m"},
		{1100499, "1.100m"},
		{1199999, "1.200m"},
		{4999501, "5m"},
		{1999501, "2m"},
		{-1100501, "-1.101m"},
		{111100501, "111.101m"},
		{999999499, "999.999m"},
		{999999501, "1"},
		{999999999, "1"},
		{1000000000, "1"},
		{-1000000000, "-1"}, // 20
		{1100000000, "1.100"},
		{-1100000000, "-1.100"},
		{1100499000, "1.100"},
		{-1100501000, "-1.101"},
		{999999499000, "999.999"},
		{999999501000, "1k"},
		{999999999999, "1k"},
		{-999999999999, "-1k"},
		{1000000000000, "1k"},
		{-1000000000000, "-1k"},
		{1100000000000, "1.100k"},
		{-1100000000000, "-1.100k"},
		{1100499000000, "1.100k"},
		{1199999000000, "1.200k"},
		{-1100501000000, "-1.101k"},
		{999999499000000, "999.999k"},
		{999999501000000, "1M"},
		{999999999999999, "1M"},
		{-999999999999999, "-1M"}, // 30
		{1000000000000000, "1M"},
		{-1000000000000000, "-1M"},
		{1100000000000000, "1.100M"},
		{-1100000000000000, "-1.100M"},
		{1100499000000000, "1.100M"},
		{-1100501000000000, "-1.101M"},
		{999999499000000000, "999.999M"},
		{999999501100000000, "1G"},
		{999999999999999999, "1G"},
		{-999999999999999999, "-1G"},
		{1000000000000000000, "1G"},
		{-1000000000000000000, "-1G"},
		{1100000000000000000, "1.100G"},
		{-1100000000000000000, "-1.100G"},
		{1999999999999999999, "2G"},
		{-1999999999999999999, "-2G"},
		{1100499000000000000, "1.100G"},
		{-1100501000000000000, "-1.101G"},
		{9223372036854775807, "9.223G"},
		{-9223372036854775807, "-9.223G"},
		{-9223372036854775808, "-9.223G"},
	}
	for i, line := range data {
		if s := nanoAsString(line.in); s != line.expected {
			t.Fatalf("%d: nanoAsString(%d).String() = %s != %s", i, line.in, s, line.expected)
		}
	}
}

func TestMicroAsString(t *testing.T) {
	data := []struct {
		in       int64
		expected string
	}{
		{0, "0"}, // 0
		{1, "1µ"},
		{-1, "-1µ"},
		{900, "900µ"},
		{-900, "-900µ"},
		{999, "999µ"},
		{-999, "-999µ"},
		{1000, "1m"},
		{-1000, "-1m"},
		{1100, "1.100m"},
		{-1100, "-1.100m"}, // 10
		{999999, "999.999m"},
		{-999999, "-999.999m"},
		{1000000, "1"},
		{-1000000, "-1"},
		{1000501, "1.001"},
		{-1000501, "-1.001"},
		{1100000, "1.100"},
		{-1100000, "-1.100"},
		{999999501, "1k"},
		{-999999501, "-1k"},
		{999999999, "1k"},
		{-999999999, "-1k"},
		{1000000000, "1k"},
		{-1000000000, "-1k"}, // 20
		{1100000000, "1.100k"},
		{-1100000000, "-1.100k"},
		{999999499999, "999.999k"},
		{-999999499999, "-999.999k"},
		{999999500001, "1M"},
		{-999999500001, "-1M"},
		{1000000000000, "1M"},
		{-1000000000000, "-1M"},
		{1100000000000, "1.100M"},
		{-1100000000000, "-1.100M"},
		{999999499999999, "999.999M"},
		{-999999499999999, "-999.999M"},
		{999999500000001, "1G"},
		{-999999500000001, "-1G"},
		{1000000000000000, "1G"}, //30
		{-1000000000000000, "-1G"},
		{1100000000000000, "1.100G"},
		{-1100000000000000, "-1.100G"},
		{999999499999999999, "999.999G"},
		{-999999499999999999, "-999.999G"},
		{999999500000000001, "1T"},
		{-999999500000000001, "-1T"},
		{1000000000000000000, "1T"},
		{-1000000000000000000, "-1T"},
		{1100000000000000000, "1.100T"},
		{-1100000000000000000, "-1.100T"},
		{-1999499999999999999, "-1.999T"},
		{1999499999999999999, "1.999T"},
		{-1999500000000000001, "-2T"},
		{1999500000000000001, "2T"},
		{9223372036854775807, "9.223T"},
		{-9223372036854775807, "-9.223T"},
		{-9223372036854775808, "-9.223T"},
	}
	for i, line := range data {
		if s := microAsString(line.in); s != line.expected {
			t.Fatalf("%d: microAsString(%d).String() = %s != %s", i, line.in, s, line.expected)
		}
	}
}

func BenchmarkCelsiusString(b *testing.B) {
	v := 10*Celsius + ZeroCelsius
	buf := bytes.Buffer{}
	for i := 0; i < b.N; i++ {
		buf.WriteString(v.String())
		buf.Reset()
	}
}

func BenchmarkCelsiusFloatf(b *testing.B) {
	v := float64(10)
	buf := bytes.Buffer{}
	for i := 0; i < b.N; i++ {
		buf.WriteString(fmt.Sprintf("%.1f°C", v))
		buf.Reset()
	}
}

func BenchmarkCelsiusFloatg(b *testing.B) {
	v := float64(10)
	buf := bytes.Buffer{}
	for i := 0; i < b.N; i++ {
		buf.WriteString(fmt.Sprintf("%g°C", v))
		buf.Reset()
	}
}

func TestAngle_Set(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want Angle
	}{
		{"Degrees", "1Degrees", Degree},
		{"Degrees", "-1Degrees", -1 * Degree},
		{"Degrees", "180.00Degrees", 180 * Degree},
		{"Degrees", "0.5Degrees", 8726646 * NanoRadian},
		{"Degrees", "0.5°", 8726646 * NanoRadian},
		{"nRadians", "1nRadians", NanoRadian},
		{"uRadians", "1uRadians", MicroRadian},
		{"mRadians", "1mRadians", MilliRadian},
		{"uRadians", "0.5uRadians", 500 * NanoRadian},
		{"mRadians", "0.5mRadians", 500 * MicroRadian},
		{"Radians", "1Radians", Radian},
		{"Pi", "1Pi", Pi},
		{"2Pi", "2Pi", 2 * Pi},
		{"2Pi", "-2Pi", -2 * Pi},
		{"0.5Pi", "0.5Pi", 1570796326 * NanoRadian},
		{"200", "200", 200 * Radian},
		{"200u", "200u", 200 * MicroRadian},
		{"1", "1", 1 * Radian},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got Angle
			fs := flag.NewFlagSet("Tests", flag.ExitOnError)
			fs.Var(&got, "angle", "value of angle")
			fs.Parse([]string{"-angle", tt.s})
			if got != tt.want {
				t.Errorf("%s wanted: %v but got: %v(%d)", tt.name, tt.want, got, got)
			}
		})

	}
}

func TestFrequency_Set(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want Frequency
	}{
		{"1uHz", "1uHz", 1 * MicroHertz},
		{"10uHz", "10uHz", 10 * MicroHertz},
		{"100uHz", "100uHz", 100 * MicroHertz},
		{"1µHz", "1µHz", 1 * MicroHertz},
		{"10µHz", "10µHz", 10 * MicroHertz},
		{"100µHz", "100µHz", 100 * MicroHertz},
		{"1mHz", "1mHz", 1 * MilliHertz},
		{"10mHz", "10mHz", 10 * MilliHertz},
		{"100mHz", "100mHz", 100 * MilliHertz},
		{"1Hz", "1Hz", 1 * Hertz},
		{"10Hz", "10Hz", 10 * Hertz},
		{"100Hz", "100Hz", 100 * Hertz},
		{"1kHz", "1kHz", 1 * KiloHertz},
		{"10kHz", "10kHz", 10 * KiloHertz},
		{"100kHz", "100kHz", 100 * KiloHertz},
		{"1MHz", "1MHz", 1 * MegaHertz},
		{"10MHz", "10MHz", 10 * MegaHertz},
		{"100MHz", "100MHz", 100 * MegaHertz},
		{"1GHz", "1GHz", 1 * GigaHertz},
		{"10GHz", "10GHz", 10 * GigaHertz},
		{"100GHz", "100GHz", 100 * GigaHertz},
		{"1THz", "1THz", 1 * TeraHertz},
		{"12.345Hz", "12.345Hz", 12345 * MilliHertz},
		{"-12.345Hz", "-12.345Hz", -12345 * MilliHertz},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got Frequency
			fs := flag.NewFlagSet("Tests", flag.ExitOnError)
			fs.Var(&got, "f", "value of angle")
			fs.Parse([]string{"-f", tt.s})
			if got != tt.want {
				t.Errorf("%s wanted: %v but got: %v(%d)", tt.name, tt.want, got, got)
			}
		})

	}
}

func TestDistance_Set(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want Distance
	}{
		{"1um", "1um", 1 * MicroMetre},
		{"10um", "10um", 10 * MicroMetre},
		{"100um", "100um", 100 * MicroMetre},
		{"1µm", "1µm", 1 * MicroMetre},
		{"10µm", "10µm", 10 * MicroMetre},
		{"100µm", "100µm", 100 * MicroMetre},
		{"1mm", "1mm", 1 * MilliMetre},
		{"10mm", "10mm", 10 * MilliMetre},
		{"100mm", "100mm", 100 * MilliMetre},
		{"1m", "1m", 1 * Metre},
		{"10m", "10m", 10 * Metre},
		{"100m", "100m", 100 * Metre},
		{"1km", "1km", 1 * KiloMetre},
		{"10km", "10km", 10 * KiloMetre},
		{"100km", "100km", 100 * KiloMetre},
		{"1Mm", "1Mm", 1 * MegaMetre},
		{"10Mm", "10Mm", 10 * MegaMetre},
		{"100Mm", "100Mm", 100 * MegaMetre},
		{"1Gm", "1Gm", 1 * GigaMetre},
		{"metre", "metre", 1 * Metre},
		{"Metre", "Metre", 1 * Metre},
		{"metres", "10metres", 10 * Metre},
		{"Metres", "10Metres", 10 * Metre},
		{"in", "in", 1 * Inch},
		{"In", "In", 1 * Inch},
		{"inch", "inch", 1 * Inch},
		{"Inch", "Inch", 1 * Inch},
		{"inches", "inches", 1 * Inch},
		{"Inches", "Inches", 1 * Inch},
		{"foot", "foot", 1 * Foot},
		{"Foot", "Foot", 1 * Foot},
		{"ft", "ft", 1 * Foot},
		{"Ft", "Ft", 1 * Foot},
		{"Feet", "10Feet", 10 * Foot},
		{"feet", "10feet", 10 * Foot},
		{"Yard", "Yard", 1 * Yard},
		{"yard", "yard", 1 * Yard},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got Distance
			fs := flag.NewFlagSet("Tests", flag.ExitOnError)
			fs.Var(&got, "d", "value of angle")
			fs.Parse([]string{"-d", tt.s})
			if got != tt.want {
				t.Errorf("%s wanted: %v but got: %v(%d)", tt.name, tt.want, got, got)
			}
		})

	}
}

func TestParseFrequency(t *testing.T) {
	tests := []struct {
		name    string
		s       string
		want    Frequency
		wantErr bool
	}{
		{"100µHz", "100µHz", 100 * MicroHertz, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseFrequency(tt.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseFrequency() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseFrequency() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestElectricalCapacitance_Set(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want ElectricalCapacitance
	}{

		{"1pF", "1pF", 1 * PicoFarad},
		{"10pF", "10pF", 10 * PicoFarad},
		{"100pF", "100pF", 100 * PicoFarad},
		{"1nF", "1nF", 1 * NanoFarad},
		{"10nF", "10nF", 10 * NanoFarad},
		{"100nF", "100nF", 100 * NanoFarad},
		{"1uF", "1uF", 1 * MicroFarad},
		{"10uF", "10uF", 10 * MicroFarad},
		{"100uF", "100uF", 100 * MicroFarad},
		{"1µF", "1µF", 1 * MicroFarad},
		{"10µF", "10µF", 10 * MicroFarad},
		{"100µF", "100µF", 100 * MicroFarad},
		{"1mF", "1mF", 1 * MilliFarad},
		{"10mF", "10mF", 10 * MilliFarad},
		{"100mF", "100mF", 100 * MilliFarad},
		{"1F", "1F", 1 * Farad},
		{"10F", "10F", 10 * Farad},
		{"100F", "100F", 100 * Farad},
		{"1kF", "1kF", 1 * KiloFarad},
		{"10kF", "10kF", 10 * KiloFarad},
		{"100kF", "100kF", 100 * KiloFarad},
		{"1f", "1f", 1 * Farad},
		{"1farad", "1farad", 1 * Farad},
		{"1Farad", "1Farad", 1 * Farad},
		{"10farads", "10farads", 10 * Farad},
		{"10Farads", "10Farads", 10 * Farad},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got ElectricalCapacitance
			fs := flag.NewFlagSet("Tests", flag.ExitOnError)
			fs.Var(&got, "farad", "value of capacitance")
			fs.Parse([]string{"-farad", tt.s})
			if got != tt.want {
				t.Errorf("%s wanted: %v but got: %v(%d)", tt.name, tt.want, got, got)
			}
		})
	}
}

func TestElectricCurrent_Set(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want ElectricCurrent
	}{

		{"1nA", "1nA", 1 * NanoAmpere},
		{"10nA", "10nA", 10 * NanoAmpere},
		{"100nA", "100nA", 100 * NanoAmpere},
		{"1uA", "1uA", 1 * MicroAmpere},
		{"10uA", "10uA", 10 * MicroAmpere},
		{"100uA", "100uA", 100 * MicroAmpere},
		{"1µA", "1µA", 1 * MicroAmpere},
		{"10µA", "10µA", 10 * MicroAmpere},
		{"100µA", "100µA", 100 * MicroAmpere},
		{"1mA", "1mA", 1 * MilliAmpere},
		{"10mA", "10mA", 10 * MilliAmpere},
		{"100mA", "100mA", 100 * MilliAmpere},
		{"1A", "1A", 1 * Ampere},
		{"10A", "10A", 10 * Ampere},
		{"100A", "100A", 100 * Ampere},
		{"1kA", "1kA", 1 * KiloAmpere},
		{"10kA", "10kA", 10 * KiloAmpere},
		{"100kA", "100kA", 100 * KiloAmpere},
		{"1MA", "1MA", 1 * MegaAmpere},
		{"10MA", "10MA", 10 * MegaAmpere},
		{"100MA", "100MA", 100 * MegaAmpere},
		{"1GA", "1GA", 1 * GigaAmpere},
		{"1a", "1a", 1 * Ampere},
		{"1Amp", "1Amp", 1 * Ampere},
		{"1amp", "1amp", 1 * Ampere},
		{"1amps", "1amps", 1 * Ampere},
		{"1Amps", "1Amps", 1 * Ampere},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got ElectricCurrent
			fs := flag.NewFlagSet("Tests", flag.ExitOnError)
			fs.Var(&got, "amps", "value of current")
			fs.Parse([]string{"-amps", tt.s})
			if got != tt.want {
				t.Errorf("%s wanted: %v but got: %v(%d)", tt.name, tt.want, got, got)
			}
		})

	}
}

func TestTemperature_Set(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want Temperature
	}{

		{"1nK", "1nK", 1 * NanoKelvin},
		{"10nK", "10nK", 10 * NanoKelvin},
		{"100nK", "100nK", 100 * NanoKelvin},
		{"1uK", "1uK", 1 * MicroKelvin},
		{"10uK", "10uK", 10 * MicroKelvin},
		{"100uK", "100uK", 100 * MicroKelvin},
		{"1µK", "1µK", 1 * MicroKelvin},
		{"10µK", "10µK", 10 * MicroKelvin},
		{"100µK", "100µK", 100 * MicroKelvin},
		{"1mK", "1mK", 1 * MilliKelvin},
		{"10mK", "10mK", 10 * MilliKelvin},
		{"100mK", "100mK", 100 * MilliKelvin},
		{"1K", "1K", 1 * Kelvin},
		{"10K", "10K", 10 * Kelvin},
		{"100K", "100K", 100 * Kelvin},
		{"1kK", "1kK", 1 * KiloKelvin},
		{"10kK", "10kK", 10 * KiloKelvin},
		{"100kK", "100kK", 100 * KiloKelvin},
		{"1MK", "1MK", 1 * MegaKelvin},
		{"10MK", "10MK", 10 * MegaKelvin},
		{"100MK", "100MK", 100 * MegaKelvin},
		{"1GK", "1GK", 1 * GigaKelvin},
		{"0C", "0C", ZeroCelsius},
		{"0°C", "0°C", ZeroCelsius},
		{"20C", "20C", ZeroCelsius + 20*Kelvin},
		{"-20C", "-20C", ZeroCelsius - 20*Kelvin},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got Temperature
			fs := flag.NewFlagSet("Tests", flag.ExitOnError)
			fs.Var(&got, "t", "value of temperature")
			fs.Parse([]string{"-t", tt.s})
			if got != tt.want {
				t.Errorf("%s wanted: %v but got: %v(%d)", tt.name, tt.want, got, got)
			}
		})

	}
}

func TestElectricPotential_Set(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want ElectricPotential
	}{

		{"1nV", "1nV", 1 * NanoVolt},
		{"10nV", "10nV", 10 * NanoVolt},
		{"100nV", "100nV", 100 * NanoVolt},
		{"1uV", "1uV", 1 * MicroVolt},
		{"10uV", "10uV", 10 * MicroVolt},
		{"100uV", "100uV", 100 * MicroVolt},
		{"1µV", "1µV", 1 * MicroVolt},
		{"10µV", "10µV", 10 * MicroVolt},
		{"100µV", "100µV", 100 * MicroVolt},
		{"1mV", "1mV", 1 * MilliVolt},
		{"10mV", "10mV", 10 * MilliVolt},
		{"100mV", "100mV", 100 * MilliVolt},
		{"1V", "1V", 1 * Volt},
		{"10V", "10V", 10 * Volt},
		{"100V", "100V", 100 * Volt},
		{"1kV", "1kV", 1 * KiloVolt},
		{"10kV", "10kV", 10 * KiloVolt},
		{"100kV", "100kV", 100 * KiloVolt},
		{"1MV", "1MV", 1 * MegaVolt},
		{"10MV", "10MV", 10 * MegaVolt},
		{"100MV", "100MV", 100 * MegaVolt},
		{"1GV", "1GV", 1 * GigaVolt},
		{"volt", "volt", 1 * Volt},
		{"volts", "volts", 1 * Volt},
		{"Volt", "Volt", 1 * Volt},
		{"Volts", "Volts", 1 * Volt},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got ElectricPotential
			fs := flag.NewFlagSet("Tests", flag.ExitOnError)
			fs.Var(&got, "v", "value of voltage")
			fs.Parse([]string{"-v", tt.s})
			if got != tt.want {
				t.Errorf("%s wanted: %v but got: %v(%d)", tt.name, tt.want, got, got)
			}
		})

	}
}

func TestElectricResistance_Set(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want ElectricResistance
	}{

		{"1nΩ", "1nΩ", 1 * NanoOhm},
		{"10nΩ", "10nΩ", 10 * NanoOhm},
		{"100nΩ", "100nΩ", 100 * NanoOhm},
		{"1uΩ", "1uΩ", 1 * MicroOhm},
		{"10uΩ", "10uΩ", 10 * MicroOhm},
		{"100uΩ", "100uΩ", 100 * MicroOhm},
		{"1µΩ", "1µΩ", 1 * MicroOhm},
		{"10µΩ", "10µΩ", 10 * MicroOhm},
		{"100µΩ", "100µΩ", 100 * MicroOhm},
		{"1mΩ", "1mΩ", 1 * MilliOhm},
		{"10mΩ", "10mΩ", 10 * MilliOhm},
		{"100mΩ", "100mΩ", 100 * MilliOhm},
		{"1Ω", "1Ω", 1 * Ohm},
		{"10Ω", "10Ω", 10 * Ohm},
		{"100Ω", "100Ω", 100 * Ohm},
		{"1kΩ", "1kΩ", 1 * KiloOhm},
		{"10kΩ", "10kΩ", 10 * KiloOhm},
		{"100kΩ", "100kΩ", 100 * KiloOhm},
		{"1MΩ", "1MΩ", 1 * MegaOhm},
		{"10MΩ", "10MΩ", 10 * MegaOhm},
		{"100MΩ", "100MΩ", 100 * MegaOhm},
		{"1GΩ", "1GΩ", 1 * GigaOhm},
		{"Ohm", "Ohm", 1 * Ohm},
		{"Ohms", "Ohms", 1 * Ohm},
		{"ohm", "ohm", 1 * Ohm},
		{"ohms", "ohms", 1 * Ohm},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got ElectricResistance
			fs := flag.NewFlagSet("Tests", flag.ExitOnError)
			fs.Var(&got, "r", "value of resistance")
			fs.Parse([]string{"-r", tt.s})
			if got != tt.want {
				t.Errorf("%s wanted: %v but got: %v(%d)", tt.name, tt.want, got, got)
			}
		})

	}
}

func TestPower_Set(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want Power
	}{

		{"1nW", "1nW", 1 * NanoWatt},
		{"10nW", "10nW", 10 * NanoWatt},
		{"100nW", "100nW", 100 * NanoWatt},
		{"1uW", "1uW", 1 * MicroWatt},
		{"10uW", "10uW", 10 * MicroWatt},
		{"100uW", "100uW", 100 * MicroWatt},
		{"1µW", "1µW", 1 * MicroWatt},
		{"10µW", "10µW", 10 * MicroWatt},
		{"100µW", "100µW", 100 * MicroWatt},
		{"1mW", "1mW", 1 * MilliWatt},
		{"10mW", "10mW", 10 * MilliWatt},
		{"100mW", "100mW", 100 * MilliWatt},
		{"1W", "1W", 1 * Watt},
		{"10W", "10W", 10 * Watt},
		{"100W", "100W", 100 * Watt},
		{"1kW", "1kW", 1 * KiloWatt},
		{"10kW", "10kW", 10 * KiloWatt},
		{"100kW", "100kW", 100 * KiloWatt},
		{"1MW", "1MW", 1 * MegaWatt},
		{"10MW", "10MW", 10 * MegaWatt},
		{"100MW", "100MW", 100 * MegaWatt},
		{"1GW", "1GW", 1 * GigaWatt},
		{"Watt", "Watt", 1 * Watt},
		{"Watts", "Watts", 1 * Watt},
		{"Watt", "Watt", 1 * Watt},
		{"Watts", "Watts", 1 * Watt},
		{"W", "W", 1 * Watt},
		{"w", "w", 1 * Watt},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got Power
			fs := flag.NewFlagSet("Tests", flag.ExitOnError)
			fs.Var(&got, "p", "value of power")
			fs.Parse([]string{"-p", tt.s})
			if got != tt.want {
				t.Errorf("%s wanted: %v but got: %v(%d)", tt.name, tt.want, got, got)
			}
		})

	}
}

func TestEnergy_Set(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want Energy
	}{

		{"1nJ", "1nJ", 1 * NanoJoule},
		{"10nJ", "10nJ", 10 * NanoJoule},
		{"100nJ", "100nJ", 100 * NanoJoule},
		{"1uJ", "1uJ", 1 * MicroJoule},
		{"10uJ", "10uJ", 10 * MicroJoule},
		{"100uJ", "100uJ", 100 * MicroJoule},
		{"1µJ", "1µJ", 1 * MicroJoule},
		{"10µJ", "10µJ", 10 * MicroJoule},
		{"100µJ", "100µJ", 100 * MicroJoule},
		{"1mJ", "1mJ", 1 * MilliJoule},
		{"10mJ", "10mJ", 10 * MilliJoule},
		{"100mJ", "100mJ", 100 * MilliJoule},
		{"1J", "1J", 1 * Joule},
		{"10J", "10J", 10 * Joule},
		{"100J", "100J", 100 * Joule},
		{"1kJ", "1kJ", 1 * KiloJoule},
		{"10kJ", "10kJ", 10 * KiloJoule},
		{"100kJ", "100kJ", 100 * KiloJoule},
		{"1MJ", "1MJ", 1 * MegaJoule},
		{"10MJ", "10MJ", 10 * MegaJoule},
		{"100MJ", "100MJ", 100 * MegaJoule},
		{"1GJ", "1GJ", 1 * GigaJoule},
		{"Joule", "Joule", 1 * Joule},
		{"Joules", "Joules", 1 * Joule},
		{"joule", "joule", 1 * Joule},
		{"joules", "joules", 1 * Joule},
		{"J", "J", 1 * Joule},
		{"j", "j", 1 * Joule},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got Energy
			fs := flag.NewFlagSet("Tests", flag.ExitOnError)
			fs.Var(&got, "e", "value of energy")
			fs.Parse([]string{"-e", tt.s})
			if got != tt.want {
				t.Errorf("%s wanted: %v but got: %v(%d)", tt.name, tt.want, got, got)
			}
		})

	}
}

func TestPressure(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want Pressure
	}{

		{"1nPa", "1nPa", 1 * NanoPascal},
		{"10nPa", "10nPa", 10 * NanoPascal},
		{"100nPa", "100nPa", 100 * NanoPascal},
		{"1uPa", "1uPa", 1 * MicroPascal},
		{"10uPa", "10uPa", 10 * MicroPascal},
		{"100uPa", "100uPa", 100 * MicroPascal},
		{"1µPa", "1µPa", 1 * MicroPascal},
		{"10µPa", "10µPa", 10 * MicroPascal},
		{"100µPa", "100µPa", 100 * MicroPascal},
		{"1mPa", "1mPa", 1 * MilliPascal},
		{"10mPa", "10mPa", 10 * MilliPascal},
		{"100mPa", "100mPa", 100 * MilliPascal},
		{"1Pa", "1Pa", 1 * Pascal},
		{"10Pa", "10Pa", 10 * Pascal},
		{"100Pa", "100Pa", 100 * Pascal},
		{"1kPa", "1kPa", 1 * KiloPascal},
		{"10kPa", "10kPa", 10 * KiloPascal},
		{"100kPa", "100kPa", 100 * KiloPascal},
		{"1MPa", "1MPa", 1 * MegaPascal},
		{"10MPa", "10MPa", 10 * MegaPascal},
		{"100MPa", "100MPa", 100 * MegaPascal},
		{"1GPa", "1GPa", 1 * GigaPascal},
		{"Pascal", "Pascal", 1 * Pascal},
		{"Pascals", "Pascals", 1 * Pascal},
		{"pascal", "pascal", 1 * Pascal},
		{"pascals", "pascals", 1 * Pascal},
		{"Pa", "Pa", 1 * Pascal},
		{"pa", "pa", 1 * Pascal},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got Pressure
			fs := flag.NewFlagSet("Tests", flag.ExitOnError)
			fs.Var(&got, "p", "value of presure")
			fs.Parse([]string{"-p", tt.s})
			if got != tt.want {
				t.Errorf("%s wanted: %v but got: %v(%d)", tt.name, tt.want, got, got)
			}
		})

	}
}

func TestLuminousIntensity_Set(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want LuminousIntensity
	}{

		{"1ncd", "1ncd", 1 * NanoCandela},
		{"10ncd", "10ncd", 10 * NanoCandela},
		{"100ncd", "100ncd", 100 * NanoCandela},
		{"1ucd", "1ucd", 1 * MicroCandela},
		{"10ucd", "10ucd", 10 * MicroCandela},
		{"100ucd", "100ucd", 100 * MicroCandela},
		{"1µcd", "1µcd", 1 * MicroCandela},
		{"10µcd", "10µcd", 10 * MicroCandela},
		{"100µcd", "100µcd", 100 * MicroCandela},
		{"1mcd", "1mcd", 1 * MilliCandela},
		{"10mcd", "10mcd", 10 * MilliCandela},
		{"100mcd", "100mcd", 100 * MilliCandela},
		{"1cd", "1cd", 1 * Candela},
		{"10cd", "10cd", 10 * Candela},
		{"100cd", "100cd", 100 * Candela},
		{"1kcd", "1kcd", 1 * KiloCandela},
		{"10kcd", "10kcd", 10 * KiloCandela},
		{"100kcd", "100kcd", 100 * KiloCandela},
		{"1Mcd", "1Mcd", 1 * MegaCandela},
		{"10Mcd", "10Mcd", 10 * MegaCandela},
		{"100Mcd", "100Mcd", 100 * MegaCandela},
		{"1Gcd", "1Gcd", 1 * GigaCandela},
		{"Candela", "Candela", 1 * Candela},
		{"Candelas", "Candelas", 1 * Candela},
		{"candela", "candela", 1 * Candela},
		{"candelas", "candelas", 1 * Candela},
		{"cd", "cd", 1 * Candela},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got LuminousIntensity
			fs := flag.NewFlagSet("Tests", flag.ExitOnError)
			fs.Var(&got, "l", "value of intensity")
			fs.Parse([]string{"-l", tt.s})
			if got != tt.want {
				t.Errorf("%s wanted: %v but got: %v(%d)", tt.name, tt.want, got, got)
			}
		})

	}
}

func TestLuminousFlux_Set(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want LuminousFlux
	}{

		{"1nlm", "1nlm", 1 * NanoLumen},
		{"10nlm", "10nlm", 10 * NanoLumen},
		{"100nlm", "100nlm", 100 * NanoLumen},
		{"1ulm", "1ulm", 1 * MicroLumen},
		{"10ulm", "10ulm", 10 * MicroLumen},
		{"100ulm", "100ulm", 100 * MicroLumen},
		{"1µlm", "1µlm", 1 * MicroLumen},
		{"10µlm", "10µlm", 10 * MicroLumen},
		{"100µlm", "100µlm", 100 * MicroLumen},
		{"1mlm", "1mlm", 1 * MilliLumen},
		{"10mlm", "10mlm", 10 * MilliLumen},
		{"100mlm", "100mlm", 100 * MilliLumen},
		{"1lm", "1lm", 1 * Lumen},
		{"10lm", "10lm", 10 * Lumen},
		{"100lm", "100lm", 100 * Lumen},
		{"1klm", "1klm", 1 * KiloLumen},
		{"10klm", "10klm", 10 * KiloLumen},
		{"100klm", "100klm", 100 * KiloLumen},
		{"1Mlm", "1Mlm", 1 * MegaLumen},
		{"10Mlm", "10Mlm", 10 * MegaLumen},
		{"100Mlm", "100Mlm", 100 * MegaLumen},
		{"1Glm", "1Glm", 1 * GigaLumen},
		{"Lumen", "Lumen", 1 * Lumen},
		{"Lumens", "Lumens", 1 * Lumen},
		{"lumen", "lumen", 1 * Lumen},
		{"lumens", "lumens", 1 * Lumen},
		{"lm", "lm", 1 * Lumen},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got LuminousFlux
			fs := flag.NewFlagSet("Tests", flag.ExitOnError)
			fs.Var(&got, "f", "value of flux")
			fs.Parse([]string{"-f", tt.s})
			if got != tt.want {
				t.Errorf("%s wanted: %v but got: %v(%d)", tt.name, tt.want, got, got)
			}
		})

	}
}

func TestSpeed_Set(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want Speed
	}{

		{"1nm/s", "1nm/s", 1 * NanoMetrePerSecond},
		{"10nm/s", "10nm/s", 10 * NanoMetrePerSecond},
		{"100nm/s", "100nm/s", 100 * NanoMetrePerSecond},
		{"1um/s", "1um/s", 1 * MicroMetrePerSecond},
		{"10um/s", "10um/s", 10 * MicroMetrePerSecond},
		{"100um/s", "100um/s", 100 * MicroMetrePerSecond},
		{"1µm/s", "1µm/s", 1 * MicroMetrePerSecond},
		{"10µm/s", "10µm/s", 10 * MicroMetrePerSecond},
		{"100µm/s", "100µm/s", 100 * MicroMetrePerSecond},
		{"1mm/s", "1mm/s", 1 * MilliMetrePerSecond},
		{"10mm/s", "10mm/s", 10 * MilliMetrePerSecond},
		{"100mm/s", "100mm/s", 100 * MilliMetrePerSecond},
		{"1m/s", "1m/s", 1 * MetrePerSecond},
		{"10m/s", "10m/s", 10 * MetrePerSecond},
		{"100m/s", "100m/s", 100 * MetrePerSecond},
		{"1km/s", "1km/s", 1 * KiloMetrePerSecond},
		{"10km/s", "10km/s", 10 * KiloMetrePerSecond},
		{"100km/s", "100km/s", 100 * KiloMetrePerSecond},
		{"1Mm/s", "1Mm/s", 1 * MegaMetrePerSecond},
		{"10Mm/s", "10Mm/s", 10 * MegaMetrePerSecond},
		{"100Mm/s", "100Mm/s", 100 * MegaMetrePerSecond},
		{"1Gm/s", "1Gm/s", 1 * GigaMetrePerSecond},
		{"m/s", "m/s", 1 * MetrePerSecond},
		{"km/h", "km/h", 1 * KilometrePerHour},
		{"mph", "mph", 1 * MilePerHour},
		{"fps", "fps", 1 * FootPerSecond},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got Speed
			fs := flag.NewFlagSet("Tests", flag.ExitOnError)
			fs.Var(&got, "s", "value of speed")
			fs.Parse([]string{"-s", tt.s})
			if got != tt.want {
				t.Errorf("%s wanted: %v but got: %v(%d)", tt.name, tt.want, got, got)
			}
		})

	}
}

func TestMass_Set(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want Mass
	}{

		{"1ng", "1ng", 1 * NanoGram},
		{"10ng", "10ng", 10 * NanoGram},
		{"100ng", "100ng", 100 * NanoGram},
		{"1ug", "1ug", 1 * MicroGram},
		{"10ug", "10ug", 10 * MicroGram},
		{"100ug", "100ug", 100 * MicroGram},
		{"1µg", "1µg", 1 * MicroGram},
		{"10µg", "10µg", 10 * MicroGram},
		{"100µg", "100µg", 100 * MicroGram},
		{"1mg", "1mg", 1 * MilliGram},
		{"10mg", "10mg", 10 * MilliGram},
		{"100mg", "100mg", 100 * MilliGram},
		{"1g", "1g", 1 * Gram},
		{"10g", "10g", 10 * Gram},
		{"100g", "100g", 100 * Gram},
		{"1kg", "1kg", 1 * KiloGram},
		{"10kg", "10kg", 10 * KiloGram},
		{"100kg", "100kg", 100 * KiloGram},
		{"1Mg", "1Mg", 1 * MegaGram},
		{"10Mg", "10Mg", 10 * MegaGram},
		{"100Mg", "100Mg", 100 * MegaGram},
		{"1Gg", "1Gg", 1 * GigaGram},
		{"gram", "gram", 1 * Gram},
		{"Gram", "Gram", 1 * Gram},
		{"grams", "grams", 1 * Gram},
		{"Grams", "Grams", 1 * Gram},
		{"ounce", "ounce", 1 * OunceMass},
		{"Ounce", "Ounce", 1 * OunceMass},
		{"Ounces", "Ounces", 1 * OunceMass},
		{"ounces", "ounces", 1 * OunceMass},
		{"tonne", "tonne", 1 * Tonne},
		{"tonnes", "tonnes", 1 * Tonne},
		{"Tonne", "Tonne", 1 * Tonne},
		{"Tonnes", "Tonnes", 1 * Tonne},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got Mass
			fs := flag.NewFlagSet("Tests", flag.ExitOnError)
			fs.Var(&got, "m", "value of mass")
			fs.Parse([]string{"-m", tt.s})
			if got != tt.want {
				t.Errorf("%s wanted: %v but got: %v(%d)", tt.name, tt.want, got, got)
			}
		})

	}
}

func TestForce_Set(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want Force
	}{

		{"1nN", "1nN", 1 * NanoNewton},
		{"10nN", "10nN", 10 * NanoNewton},
		{"100nN", "100nN", 100 * NanoNewton},
		{"1uN", "1uN", 1 * MicroNewton},
		{"10uN", "10uN", 10 * MicroNewton},
		{"100uN", "100uN", 100 * MicroNewton},
		{"1µN", "1µN", 1 * MicroNewton},
		{"10µN", "10µN", 10 * MicroNewton},
		{"100µN", "100µN", 100 * MicroNewton},
		{"1mN", "1mN", 1 * MilliNewton},
		{"10mN", "10mN", 10 * MilliNewton},
		{"100mN", "100mN", 100 * MilliNewton},
		{"1N", "1N", 1 * Newton},
		{"10N", "10N", 10 * Newton},
		{"100N", "100N", 100 * Newton},
		{"1kN", "1kN", 1 * KiloNewton},
		{"10kN", "10kN", 10 * KiloNewton},
		{"100kN", "100kN", 100 * KiloNewton},
		{"1MN", "1MN", 1 * MegaNewton},
		{"10MN", "10MN", 10 * MegaNewton},
		{"100MN", "100MN", 100 * MegaNewton},
		{"1GN", "1GN", 1 * GigaNewton},
		{"Newton", "Newton", 1 * Newton},
		{"newton", "newton", 1 * Newton},
		{"newtons", "newtons", 1 * Newton},
		{"Newtons", "Newtons", 1 * Newton},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got Force
			fs := flag.NewFlagSet("Tests", flag.ExitOnError)
			fs.Var(&got, "f", "value of force")
			fs.Parse([]string{"-f", tt.s})
			if got != tt.want {
				t.Errorf("%s wanted: %v but got: %v(%d)", tt.name, tt.want, got, got)
			}
		})

	}
}

func TestRelativeHumidity_Set(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want RelativeHumidity
	}{
		{"1rh", (1 * PercentRH).String(), 1 * PercentRH},
		{"rh", (1 * PercentRH).String(), 1 * PercentRH},
		{"rh", (1 * PercentRH).String(), 1 * PercentRH},
		{"0.00001%rH", "0.00001%rH", 1 * TenthMicroRH},
		{"0.0001%rH", "0.0001%rH", 1 * MicroRH},
		{"1mrH", "1mrH", 1 * MilliRH},
		{"1urH", "1urH", 1 * MicroRH},
		{"0.1%rH", "0.1%rH", 1 * MilliRH},
		{"1%rH", "1%rH", 1 * PercentRH},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got RelativeHumidity
			fs := flag.NewFlagSet("Tests", flag.ExitOnError)
			fs.Var(&got, "f", "value of humidity")
			fs.Parse([]string{"-f", tt.s})
			if got != tt.want {
				t.Errorf("%s wanted: %v but got: %v(%d)", tt.name, tt.want, got, got)
			}
		})

	}
}

func TestMeta_Set(t *testing.T) {
	var degree Angle
	var metre Distance
	var amp ElectricCurrent
	var volt ElectricPotential
	var ohm ElectricResistance
	var farad ElectricalCapacitance
	var newton Force
	var hertz Frequency
	var gram Mass
	var pascal Pressure
	var humidity RelativeHumidity
	var metresPerSecond Speed
	var celsius Temperature
	var watt Power
	var joule Energy
	var candela LuminousIntensity
	var lux LuminousFlux

	tests := []struct {
		name    string
		v       flag.Value
		s       string
		wantErr bool
	}{
		{"errAngle", &degree, "1.1.1.1", true},
		{"errDistance", &metre, "1.1.1.1", true},
		{"errElectricCurrent", &amp, "1.1.1.1", true},
		{"errElectricPotential", &volt, "1.1.1.1", true},
		{"errElectricResistance", &ohm, "1.1.1.1", true},
		{"errElectricalCapacitance", &farad, "1.1.1.1", true},
		{"errForce", &newton, "1.1.1.1", true},
		{"errFrequency", &hertz, "1.1.1.1", true},
		{"errMass", &gram, "1.1.1.1", true},
		{"errPressure", &pascal, "1.1.1.1", true},
		{"errRelativeHumidity", &humidity, "1.1.1.1", true},
		{"errSpeed", &metresPerSecond, "1.1.1.1", true},
		{"errTemperature", &celsius, "1.1.1.1", true},
		{"errPower", &watt, "1.1.1.1", true},
		{"errEnergy", &joule, "1.1.1.1", true},
		{"errLuminousIntensity", &candela, "1.1.1.1", true},
		{"errLuminousFlux", &lux, "1.1.1.1", true},
		{"errAngle", &degree, "1.1.1.1", true},
		//Mininmal Implementation un-comment for WIP.
		{"Angle", &degree, "1", false},
		{"Distance", &metre, "1", false},
		{"ElectricCurrent", &amp, "1", false},
		{"ElectricPotential", &volt, "1", false},
		{"ElectricResistance", &ohm, "1", false},
		{"ElectricalCapacitance", &farad, "1", false},
		{"Force", &newton, "1", false},
		{"Frequency", &hertz, "1", false},
		{"Mass", &gram, "1", false},
		{"Pressure", &pascal, "1", false},
		{"RelativeHumidity", &humidity, "1", false},
		{"Speed", &metresPerSecond, "1", false},
		{"Temperature", &celsius, "1", false},
		{"Power", &watt, "1", false},
		{"Energy", &joule, "1", false},
		{"LuminousIntensity", &candela, "1", false},
		{"LuminousFlux", &lux, "1", false},
	}

	for _, tt := range tests {
		got := tt.v.Set(tt.s)
		if tt.wantErr && got == nil {
			t.Errorf("%s expected error but got none", tt.name)
		}
		if !tt.wantErr && got != nil {
			t.Errorf("%s got unexpected error%v", tt.name, got)
		}
	}
}
