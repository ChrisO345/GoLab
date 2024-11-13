package series

type floatElement struct {
	e float64
	nan bool
}

// force implementation of Element interface
var _ Element = (*floatElement)(nil)

func (f *floatElement) Set(value interface{}) {
	f.nan = false

	switch v := value.(type) {
	case float64:
		f.e = v
	case int:
		f.e = float64(v)
	case bool:
		if v {
			f.e = 1.0
		} else {
			f.e = 0.0
		}
	default:
		f.nan = true
		return
	}
}

func (f floatElement) Get() interface{} {
	return f.e
}

func (f floatElement) IsNA() bool {
	return f.nan
}

func (f floatElement) Type() Type {
	return Float
}

func (f floatElement) IsNumeric() bool {
	return true
}
