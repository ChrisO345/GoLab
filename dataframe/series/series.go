package series

import (
	"fmt"
)

// Series is a collection of elements of the same type and
// is the basic building block of a DataFrame
type Series struct {
	Name     string
	elements Elements
	t        Type
}

// Elements is an interface that defines the methods that a collection of elements must implement
type Elements interface {
	Elem(int) Element
	Len() int
	Values() []interface{}
}

// Element is an interface that defines the methods that an element must implement
type Element interface {
	Set(interface{})
	Get() interface{}

	IsNA() bool
	IsNumeric() bool
	Type() Type
}

// intElement is the implementation of the Element interface for int types
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

// floatElement is the implementation of the Element interface for float types
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

// booleanElements is the implementation of the Element interface for float types
type booleanElements []booleanElement

func (b booleanElements) Len() int           { return len(b) }
func (b booleanElements) Elem(j int) Element { return &b[j] }
func (b booleanElements) Values() []interface{} {
	v := make([]interface{}, len(b))
	for j, e := range b {
		v[j] = e.e
	}
	return v
}

// Type defines the type of the series
type Type string

const (
	Int     Type = "int"
	Float   Type = "float"
	Boolean Type = "bool"
	String  Type = "string"
	Runic Type = "rune"
)

// NewSeries creates a new series from a slice of values of type t, and a name
func NewSeries(v interface{}, t Type, name string) Series {
	s := Series{Name: name, t: t}

	allocMemory := func(n int) {
		switch t {
		case Int:
			s.elements = make(intElements, n)
		case Float:
			s.elements = make(floatElements, n)
		case Boolean:
			s.elements = make(booleanElements, n)
		case String:
			panic("not implemented")
		case Runic:
			panic("not implemented")
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
		l := len(v_)
		allocMemory(l)
		for i, e := range v_ {
			s.elements.Elem(i).Set(e)
		}
	case []rune:
		panic("not implemented")

	default:
		panic("unsupported type")
	}

	return s
}

// Copy returns a memory copy of the series
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
	case Boolean:

	case String:
		panic("not implemented")
	case Runic:
		panic("not implemented")
	}

	return Series{
		Name:     name,
		elements: elements,
		t:        t,
	}
}

// Len returns the number of elements in the series
func (s Series) Len() int {
	return s.elements.Len()
}

// String returns the Stringer implementation of the series
func (s Series) String() string {
	return fmt.Sprintf("{%v %v %v}", s.Name, s.elements.Values(), s.t)
}

// Val returns the value of the element at index i
func (s Series) Val(i int) interface{} {
	return s.elements.Elem(i).Get()
}

// Elem returns the element at index i
func (s Series) Elem(i int) Element {
	return s.elements.Elem(i)
}

// HasNa returns true if the series has any NA values
func (s Series) HasNa() bool {
	for i := 0; i < s.Len(); i++ {
		if s.Elem(i).IsNA() {
			return true
		}
	}
	return false
}

// Slice returns a copy of the series from index a to index b
func (s Series) Slice(a, b int) Series {
	if a < 0 {
		panic(fmt.Errorf("a index %v out of range", a))
	}

	if b > s.Len() || a > b {
		panic(fmt.Errorf("b index %v out of range", b))
	}

	se := Series{Name: s.Name, t: s.t}
	n := b - a

	allocMemory := func(n int) {
		switch s.t {
		case Int:
			se.elements = make(intElements, n)
		case Float:
			se.elements = make(floatElements, n)
		}
	}
	allocMemory(n)

	for i := a; i < b; i++ {
		se.Elem(i - a).Set(s.Val(i))
	}
	return se
}

// Head returns a slice of the first n elements of the series
func (s Series) Head(n int) Series {
	return s.Slice(0, n)
}

// Tail returns a slice of the last n elements of the series
func (s Series) Tail(n int) Series {
	return s.Slice(s.Len()-n, s.Len())
}

// Sort sorts the series in place via bubble sort TODO: replace with merge sort later
func (s Series) Sort() {
	n := s.Len()
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			switch s.t {
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
			case Boolean:
				panic("not implemented")
			case String:
				panic("not implemented")
			case Runic:
				panic("not implemented")
			}
		}
	}
}

// SortedIndex returns the indices of the series sorted in ascending order
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
			switch s.t {
			case Int:
				if s.Val(index[j]).(int) > s.Val(index[j+1]).(int) {
					swap = true
				}
			case Float:
				if s.Val(index[j]).(float64) > s.Val(index[j+1]).(float64) {
					swap = true
				}
			case Boolean:
				panic("not implemented")
			case String:
				panic("not implemented")
			case Runic:
				panic("not implemented")
			}
			if swap {
				index[j], index[j+1] = index[j+1], index[j]
			}
		}
	}

	return index
}

// Order returns the series with the elements ordered according to the positions slice
func (s Series) Order(positions ...int) Series {
	if len(positions) != s.Len() {
		panic(fmt.Errorf("series and new positions must be the same length"))
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

// Count returns the number of occurrences of the value v in the series
func (s Series) Count(v interface{}) int {
	count := 0
	for i := 0; i < s.Len(); i++ {
		if s.Val(i) == v {
			count++
		}
	}
	return count
}

// Unique returns the true if there are no duplicates in the series
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

// Homogeneous returns true if there is only one value in the series
func (s Series) Homogeneous() bool {
	if s.Len() == 0 {
		panic(fmt.Errorf("cannot check homogeneity of an empty series"))
	}

	first := s.Val(0)
	for i := 1; i < s.Len(); i++ {
		if s.Val(i) != first {
			return false
		}
	}
	return true
}

// NUnique returns the number of unique values in the series
func (s Series) NUnique() int {
	seen := make(map[interface{}]struct{})
	for i := 0; i < s.Len(); i++ {
		seen[s.Val(i)] = struct{}{}
	}
	return len(seen)
}

// ValueCounts returns a slice of the unique values in the series
func (s Series) ValueCounts() map[interface{}]int {
	seen := make(map[interface{}]int)
	for i := 0; i < s.Len(); i++ {
		seen[s.Val(i)] = seen[s.Val(i)] + 1
	}

	return seen
}

// Type returns the type of the series
func (s Series) Type() Type {
	return s.t
}

// IsNumeric returns true if the series is of a numeric type (int, float, bool)
func (s Series) IsNumeric() bool {
	return s.t == Int || s.t == Float || s.t == Boolean
}

// IsObject returns true if the series is of a non-numeric type (string, rune, object)
func (s Series) IsObject() bool {
	return s.t == String || s.t == Runic
}
