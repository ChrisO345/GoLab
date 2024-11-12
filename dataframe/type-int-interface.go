package dataframe

type intElement struct {
	e   int
	nan bool
}

// force implementation of Element interface
var _ Element = (*intElement)(nil)

func (i *intElement) Set(value interface{}) {
	i.nan = false

	switch v := value.(type) {
	case int:
		i.e = v
	case bool:
		if v {
			i.e = 1
		} else {
			i.e = 0
		}
	default:
		i.nan = true
		return
	}
}

func (i intElement) Get() interface{} {
	return i.e
}

func (i intElement) IsNA() bool {
	return i.nan
}

func (i intElement) Type() Type {
	return Int
}
