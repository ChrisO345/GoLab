package series

// TODO: Move to a separate package??

import "fmt"

type Series struct {
	Name     string
	elements Elements
	t        Type
}

type Elements interface {
	Elem(int) Element
	Len() int
	Values() []interface{}
}

type Element interface {
	Set(interface{})
	Get() interface{}

	IsNA() bool
	IsNumeric() bool
	Type() Type
}

type intElements []intElement

func (i intElements) Len() int           { return len(i) }
func (i intElements) Elem(j int) Element { return &i[j] }
func (i intElements) Values() []interface{} { // TODO: improve the way that this is implemented
	v := make([]interface{}, len(i))
	for j, e := range i {
		v[j] = e.e
	}
	return v
}

type floatElements []floatElement

func (f floatElements) Len() int           { return len(f) }
func (f floatElements) Elem(j int) Element { return &f[j] }
func (f floatElements) Values() []interface{} {
	v := make([]interface{}, len(f))
	for j, e := range f {
		v[j] = e.e
	}
	return v
}

type Type string

const (
	Int   Type = "int"
	Float Type = "float"
	//String  Type = "string"
	//Boolean Type = "bool"
)

func New(v interface{}, t Type, name string) Series {
	s := Series{Name: name, t: t}

	allocMemory := func(n int) {
		switch t {
		case Int:
			s.elements = make(intElements, n)
		case Float:
			s.elements = make(floatElements, n)
		}
	}

	if v == nil {
		allocMemory(1)
		s.elements.Elem(0).Set(nil)
		return s
	}

	switch v_ := v.(type) {
	case []string:
		panic("not implemented")
	case []int:
		l := len(v_)
		allocMemory(l)
		for i, e := range v_ {
			s.elements.Elem(i).Set(e)
		}
	case []float64:
		l := len(v_)
		allocMemory(l)
		for i, e := range v_ {
			s.elements.Elem(i).Set(e)
		}
	case []bool:
		panic("not implemented")

	default:
		panic("unsupported type")
	}

	return s
}

func (s Series) Copy() Series {
	name := s.Name
	t := s.t

	var elements Elements
	switch s.t {
	case Int:
		elements = make(intElements, s.elements.Len())
		copy(elements.(intElements), s.elements.(intElements))
	case Float:
		elements = make(floatElements, s.elements.Len())
		copy(elements.(floatElements), s.elements.(floatElements))
	}

	return Series{
		Name:     name,
		elements: elements,
		t:        t,
	}
}

func (s Series) Len() int {
	return s.elements.Len()
}

// String implements the stringer interface for Series
func (s Series) String() string {
	return fmt.Sprintf("{%v %v %v}", s.Name, s.elements.Values(), s.t)
}

func (s Series) Val(i int) interface{} {
	return s.elements.Elem(i).Get()
}

func (s Series) Elem(i int) Element {
	return s.elements.Elem(i)
}

func (s Series) HasNa() bool {
	for i := 0; i < s.Len(); i++ {
		if s.Elem(i).IsNA() {
			return true
		}
	}
	return false
}

func (s Series) Head(n int) Series {
	se := Series{Name: s.Name, t: s.t}

	allocMemory := func(n int) {
		switch s.t {
		case Int:
			se.elements = make(intElements, n)
		case Float:
			se.elements = make(floatElements, n)
		}
	}
	allocMemory(n)

	for i := 0; i < n; i++ {
		se.Elem(i).Set(s.Val(i))
	}
	return se
}

func (s Series) Tail(n int) Series {
	se := Series{Name: s.Name, t: s.t}

	allocMemory := func(n int) {
		switch s.t {
		case Int:
			se.elements = make(intElements, n)
		case Float:
			se.elements = make(floatElements, n)
		}
	}
	allocMemory(n)

	for i := 0; i < n; i++ {
		se.Elem(i).Set(s.Val(s.Len() - n + i))
	}
	return se
}
