package series

// TODO: Should methods be pointers to allow in-place modification and better memory handling??

import (
	"fmt"
)

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
	//Runic Type = "rune"
)

func NewSeries(v interface{}, t Type, name string) Series {
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

func (s Series) Slice(from, to int) Series {
	if from < 0 || to > s.Len() {
		panic("out of bounds")
	}

	if from > to {
		panic("from must be less than to")
	}

	se := Series{Name: s.Name, t: s.t}
	n := to - from

	allocMemory := func(n int) {
		switch s.t {
		case Int:
			se.elements = make(intElements, n)
		case Float:
			se.elements = make(floatElements, n)
		}
	}
	allocMemory(n)

	for i := from; i < to; i++ {
		se.Elem(i - from).Set(s.Val(i))
	}
	return se
}

func (s Series) Head(n int) Series {
	return s.Slice(0, n)
}

func (s Series) Tail(n int) Series {
	return s.Slice(s.Len()-n, s.Len())
}

// Sort sorts the series in place via bubble sort TODO: replace with merge sort later
func (s Series) Sort() {
	n := s.Len()
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			switch s.t {  // TODO: expand for other types
			case Int:
				if s.Val(j).(int) > s.Val(j+1).(int) {
					temp := s.Val(j)
					s.Elem(j).Set(s.Val(j + 1))
					s.Elem(j + 1).Set(temp)
				}
			case Float:
				if s.Val(j).(float64) > s.Val(j+1).(float64) {
					temp := s.Val(j)
					s.Elem(j).Set(s.Val(j + 1))
					s.Elem(j + 1).Set(temp)
				}
			}
		}
	}
}

// SortedIndex returns the index of what would be a sorted series
func (s Series) SortedIndex() []int {
	n := s.Len()
	index := make([]int, n)
	for i := 0; i < n; i++ {
		index[i] = i
	}

	// Bubble Sort TODO: replace with more efficient sort such as merge sort
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			swap := false
			switch s.t { // TODO: expand for more types
			case Int:
				if s.Val(index[j]).(int) > s.Val(index[j+1]).(int) {
					swap = true
				}
			case Float:
				if s.Val(index[j]).(float64) > s.Val(index[j+1]).(float64) {
					swap = true
				}
			}
			if swap {
				index[j], index[j+1] = index[j+1], index[j]
			}
		}
	}

	return index
}

func (s Series) Order(positions ...int) Series {
	if len(positions) != s.Len() {
		panic("series and new positions must be the same length")
	}

	// Need to copy otherwise positions collection will mutate
	newPositions := make([]int, s.Len())
	for i, pos := range positions {
		newPositions[i] = pos
	}

	for newPos, oldPos := range newPositions {
		if oldPos == newPos {
			continue
		}

		temp := s.Val(oldPos)
		s.Elem(oldPos).Set(s.Val(newPos))
		s.Elem(newPos).Set(temp)

		for i, pos := range newPositions {
			if pos == newPos {
				newPositions[i] = oldPos
				newPositions[newPos] = newPos
				break
			}
		}
	}

	return s
}

func (s Series) Count(v interface{}) int {
	count := 0
	for i := 0; i < s.Len(); i++ {
		if s.Val(i) == v {
			count++
		}
	}
	return count
}

func (s Series) Unique() bool {
	seen := make(map[interface{}]struct{})
	for i := 0; i < s.Len(); i++ {
		if _, ok := seen[s.Val(i)]; ok {
			return false
		}
		seen[s.Val(i)] = struct{}{}
	}
	return true
}

func (s Series) NUnique() bool {
	if s.Len() == 0 {
		panic("empty series...")
	}

	first := s.Val(0)
	for i := 1; i < s.Len(); i++ {
		if s.Val(i) != first {
			return false
		}
	}
	return true
}

func (s Series) Type() Type {
	return s.t
}

func (s Series) IsNumeric() bool {
	return s.t == Int || s.t == Float // FIXME: when implementing other types
}

func (s Series) IsObject() bool {
	return false // FIXME: when implementing other types
}