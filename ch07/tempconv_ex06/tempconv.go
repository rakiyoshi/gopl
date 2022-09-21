package tempconv

import (
	"flag"
	"fmt"
)

type Celsius float64
type Fahrenheit float64
type Kelvin float64

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%g°F", k) }
func CToF(c Celsius) Fahrenheit     { return Fahrenheit(c*9/5 + 32) }
func CToK(c Celsius) Kelvin         { return Kelvin(c + 273.15) }
func FToC(f Fahrenheit) Celsius     { return Celsius((f - 32) * 5 / 9) }
func FToK(f Fahrenheit) Kelvin      { return CToK(FToC(f)) }
func KToC(k Kelvin) Celsius         { return Celsius(k - 273.15) }
func KToF(k Kelvin) Fahrenheit      { return CToF(KToC(k)) }

type celsiusFlag struct{ Celsius }

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	case "K":
		f.Celsius = KToC(Kelvin(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

// CelsiusFlag defines specified name, default value and usage
// and returns the address of theflag variable flag arguments are degree and unit.
// for example, "100C"
func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}
