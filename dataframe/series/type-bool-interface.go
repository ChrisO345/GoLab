package series

import "math"

type booleanElement struct {
	e   bool
	nan bool
}

// force implementation of Element interface
var _ Element = (*booleanElement)(nil)

func (b *booleanElement) Set(value interface{}) {
	b.nan = false

	switch v := value.(type) {
	case int:
		b.e = v != 0
	case bool:
		b.e = v
	case float64:
		if math.IsNaN(v) || math.IsInf(v, 0) {
			b.nan = true
			return
		}
		b.e = v != 0.0
	default:
		b.nan = true
		return
	}
}

func (b booleanElement) Get() interface{} {
	return b.e
}

func (b booleanElement) IsNA() bool {
	return b.nan
}

func (b booleanElement) Type() Type {
	return Boolean
}

func (b booleanElement) IsNumeric() bool {
	return true
}
