package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100

	SymbolF = "℉"
	SymbolK = "K"
	SymbolC = "℃"
)

func (c Celsius) String() string    { return fmt.Sprintf("%g%s", c, SymbolC) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g%s", f, SymbolF) }
func (k Kelvin) String() string     { return fmt.Sprintf("%g%s", k, SymbolK) }

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*1.8 + 32) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 1.8) }

// KToC converts a Kelvin temperature to Celsius.
func KToC(k Kelvin) Celsius { return Celsius(k - 273.15) }

// CToK converts a Celsius  temperature to Kelvin.
func CToK(c Celsius) Kelvin { return Kelvin(c + 273.15) }

// FToK converts a Fahrenheit temperature to Kelvin.
func FToK(f Fahrenheit) Kelvin { return Kelvin((f + 459.67) / 1.8) }

// KToF converts a Kelvin  temperature to Fahrenheit.
func KToF(k Kelvin) Fahrenheit { return Fahrenheit(k*1.8 - 459.67) }
