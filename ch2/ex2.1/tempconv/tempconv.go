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