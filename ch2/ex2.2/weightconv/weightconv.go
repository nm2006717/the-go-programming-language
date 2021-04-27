package weightconv

import "fmt"

type Pound float64
type Kilogram float64

const (
	SymbolLb             = "lb"
	SymbolKg             = "kg"
	oneKilogram Pound    = 2.204623
	onePound    Kilogram = 0.453592
)

func (p Pound) String() string {
	return fmt.Sprintf("%g%s", p, SymbolLb)
}

func (kg Kilogram) String() string {
	return fmt.Sprintf("%g%s", kg, SymbolKg)
}

func KgToLb(kg Kilogram) Pound {
	return Pound(kg) * oneKilogram
}

func LbToKg(lb Pound) Kilogram {
	return Kilogram(lb) * onePound
}
