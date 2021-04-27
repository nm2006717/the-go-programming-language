package lengthconv

import "fmt"

type Meter float64
type Foot float64

const (
	SymbolM        = "m"
	SymbolFt       = "ft"
	OneMeter Foot  = 39.370079
	OneFoot  Meter = 0.0254
)

func (m Meter) String() string {
	return fmt.Sprintf("%g%s", m, SymbolM)
}

func (ft Foot) String() string {
	return fmt.Sprintf("%g%s", ft, SymbolFt)
}

func MtoF(m Meter) Foot {
	return Foot(m) * OneMeter
}

func FToM(f Foot) Meter {
	return Meter(f) * OneFoot
}
