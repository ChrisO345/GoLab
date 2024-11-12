package dataframe
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
	Type() Type
}

type intElements []intElement

func (i intElements) Len() int {
	return len(i)
}

func (i intElements) Elem(j int) Element {
	return &i[j]
}

func (i intElements) Values() []interface{} {
	v := make([]interface{}, len(i))
	for j, e := range i {
		v[j] = e.e
	}
	return v
}

type Type string

const (
	Int Type = "int"
	//String  Type = "string"
	//Float   Type = "float"
	//Boolean Type = "bool"
)

func NewSeries(v interface{}, t Type, name string) Series {
	s := Series{Name: name, t: t}

	allocMemory := func(n int) {
		switch t {
		case Int:
			s.elements = make(intElements, n)
		}
	}

	if v == nil {
		allocMemory(1)
		s.elements.Elem(0).Set(nil)
		return s
	}

	switch v_ := v.(type) {
	case []int:
		l := len(v_)
		allocMemory(l)
		for i, e := range v_ {
			s.elements.Elem(i).Set(e)
		}
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
