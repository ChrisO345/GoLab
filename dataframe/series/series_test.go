package series

import (
	"math"
	"testing"
)

func TestNewSeriesInt(t *testing.T) {
	expected := "{Integers [1 2 3] int}"
	s := NewSeries([]int{1, 2, 3}, Int, "Integers")

	if s.String() != expected {
		t.Errorf("Expected:\n%v\nGot:\n%v", expected, s.String())
	}
	// Output: {Integers [1 2 3] int}
}

func TestNewSeriesFloat(t *testing.T) {
	expected := "{Floats [1.1 2.2 3.3] float}"
	s := NewSeries([]float64{1.1, 2.2, 3.3}, Float, "Floats")

	if s.String() != expected {
		t.Errorf("Expected:\n%v\nGot:\n%v", expected, s.String())
	}
	// Output: {Floats [1.1 2.2 3.3] float}
}

func TestNewSeriesBool(t *testing.T) {
	expected := "{Booleans [true false true] bool}"
	s := NewSeries([]bool{true, false, true}, Boolean, "Booleans")

	if s.String() != expected {
		t.Errorf("Expected:\n%v\nGot:\n%v", expected, s.String())
	}
	// Output: {Booleans [true false true] bool}
}

func TestSeries_Slice(t *testing.T) {
	expected := "{Integers [2 3] int}"
	s := NewSeries([]int{1, 2, 3, 4}, Int, "Integers")
	se := s.Slice(1, 3)

	if se.String() != expected {
		t.Errorf("Expected:\n%v\nGot:\n%v", expected, s.String())
	}
	// Output: {Integers [2 3] int}
}

func TestSeries_Head(t *testing.T) {
	expected := "{Integers [1 2] int}"
	s := NewSeries([]int{1, 2, 3, 4}, Int, "Integers")
	se := s.Head(2)

	if se.String() != expected {
		t.Errorf("Expected:\n%v\nGot:\n%v", expected, s.String())
	}
	// Output: {Integers [1 2] int}
}

func TestSeries_Tail(t *testing.T) {
	expected := "{Integers [3 4] int}"
	s := NewSeries([]int{1, 2, 3, 4}, Int, "Integers")
	se := s.Tail(2)

	if se.String() != expected {
		t.Errorf("Expected:\n%v\nGot:\n%v", expected, s.String())
	}
	// Output: {Integers [3 4] int}
}

func TestSeries_SortInt(t *testing.T) {
	expected := "{Integers [1 2 3] int}"
	s := NewSeries([]int{3, 1, 2}, Int, "Integers")
	s.Sort()

	if s.String() != expected {
		t.Errorf("Expected:\n%v\nGot:\n%v", expected, s.String())
	}
	// Output: {Integers [1 2 3] int}
}

func TestSeries_SortFloat(t *testing.T) {
	expected := "{Floats [1.1 2.2 3.3] float}"
	s := NewSeries([]float64{3.3, 1.1, 2.2}, Float, "Floats")
	s.Sort()

	if s.String() != expected {
		t.Errorf("Expected:\n%v\nGot:\n%v", expected, s.String())
	}
	// Output: {Floats [1.1 2.2 3.3] float}
}

func TestSeries_OrderInt(t *testing.T) {
	expected := "{Integers [3 1 2] int}"
	s := NewSeries([]int{1, 2, 3}, Int, "Integers")
	se := s.Order([]int{2, 0, 1}...)

	if se.String() != expected {
		t.Errorf("Expected:\n%v\nGot:\n%v", expected, s.String())
	}
	// Output: {Integers [3 2 1] int}
}

func TestSeries_ValueCounts(t *testing.T) {
	expected := map[interface{}]int{1: 2, 2: 2, 3: 1}
	s := NewSeries([]int{1, 2, 3, 2, 1}, Int, "Integers")
	unique := s.ValueCounts()

	for k, v := range unique {
		if expected[k] != v {
			t.Errorf("Expected:\n%v\nGot:\n%v", expected, unique)
			break
		}
	}
	// Output: map[1:2 2:2 3:1]
}

func TestSeries_Sort(t *testing.T) {
	expected := "{Integers [1 2 3] int}"
	s := NewSeries([]int{3, 1, 2}, Int, "Integers")
	s.Sort()

	if s.String() != expected {
		t.Errorf("Expected:\n%v\nGot:\n%v", expected, s.String())
	}
	// Output: {Integers [1 2 3] int}

	expected = "{Integers [1 2 3 4 5 6 7 8 9] int}"
	s = NewSeries([]int{3, 9, 5, 7, 6, 8, 1, 2, 4}, Int, "Integers")
	s.Sort()

	if s.String() != expected {
		t.Errorf("Expected:\n%v\nGot:\n%v", expected, s.String())
	}
}

func TestSeries_Count(t *testing.T) {
	s := NewSeries([]int{1, 2, 3, 2, 1}, Int, "Integers")
	count := s.Count(1)

	if count != 2 {
		t.Errorf("Expected:\n%v\nGot:\n%v", 2, count)
	}
	// Output: 2
}

func TestSeries_Unique(t *testing.T) {
	s := NewSeries([]int{1, 4, 5, 2}, Int, "Integers")
	unique := s.Unique()

	if !unique {
		t.Errorf("Expected:\n%v\nGot:\n%v", true, unique)
	}
	// Output: true

	s = NewSeries([]int{1, 2, 3, 2, 1, 1}, Int, "Integers")
	unique = s.Unique()

	if unique {
		t.Errorf("Expected:\n%v\nGot:\n%v", false, unique)
	}
	// Output: false
}

func TestSeries_Homogeneous(t *testing.T) {
	s := NewSeries([]float64{1.1, 1.1, 1.1, 1.1}, Float, "Floats")
	homogeneous := s.Homogeneous()

	if !homogeneous {
		t.Errorf("Expected:\n%v\nGot:\n%v", true, homogeneous)
	}
	// Output: true

	s = NewSeries([]float64{1.1, 1.1, 1.1, 1.2}, Float, "Floats")
	homogeneous = s.Homogeneous()

	if homogeneous {
		t.Errorf("Expected:\n%v\nGot:\n%v", false, homogeneous)
	}
}

func TestSeries_Copy(t *testing.T) {
	s := NewSeries([]int{1, 2, 3}, Int, "Integers")
	se := s.Copy()

	if se.String() != s.String() {
		t.Errorf("Expected:\n%v\nGot:\n%v", s.String(), se.String())
	}
	// Output: {Integers [1 2 3] int}
}

func TestSeries_Elem(t *testing.T) {
	s := NewSeries([]int{1, 2, 3}, Int, "Integers")
	e := s.Elem(1)

	if e.Get() != 2 {
		t.Errorf("Expected:\n%v\nGot:\n%v", 2, e.Get())
	}
	// Output: 2

	s = NewSeries([]bool{true, false, true}, Boolean, "Booleans")
	e = s.Elem(1)

	if e.Get() != false {
		t.Errorf("Expected:\n%v\nGot:\n%v", false, e.Get())
	}
}

func TestSeries_IsNumeric(t *testing.T) {
	s := NewSeries([]int{1, 2, 3}, Int, "Integers")
	numeric := s.IsNumeric()

	if !numeric {
		t.Errorf("Expected:\n%v\nGot:\n%v", true, numeric)
	}
	// Output: true

	s = NewSeries([]bool{true, false, true}, Boolean, "Booleans")
	numeric = s.IsNumeric()

	if !numeric {
		t.Errorf("Expected:\n%v\nGot:\n%v", true, numeric)
	}
	// Output: true
}

func TestSeries_IsObject(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Error: %v", r)
		}
	}()

	panic("no object series implemented")
}

func TestSeries_Len(t *testing.T) {
	s := NewSeries([]int{1, 2, 3}, Int, "Integers")
	l := s.Len()

	if l != 3 {
		t.Errorf("Expected:\n%v\nGot:\n%v", 3, l)
	}
	// Output: 3
}

func TestSeries_HasNa(t *testing.T) {
	s := NewSeries([]int{1, 2, 3}, Int, "Integers")
	na := s.HasNa()

	if na {
		t.Errorf("Expected:\n%v\nGot:\n%v", false, na)
	}
	// Output: false

	s = NewSeries([]float64{1.1, 1.2, 1.3, 1.4, 1.5, math.Inf(1)}, Float, "Floats")
	na = s.HasNa()

	if !na {
		t.Errorf("Expected:\n%v\nGot:\n%v", true, na)
	}
	// Output: true

	s = NewSeries([]float64{1.1, 1.2, 1.3, 1.4, 1.5, math.NaN()}, Float, "Floats")
	na = s.HasNa()

	if !na {
		t.Errorf("Expected:\n%v\nGot:\n%v", true, na)
	}
}

func TestSeries_NUnique(t *testing.T) {
	s := NewSeries([]int{1, 2, 3, 2, 1}, Int, "Integers")
	unique := s.NUnique()

	if unique != 3 {
		t.Errorf("Expected:\n%v\nGot:\n%v", 3, unique)
	}
	// Output: 3
}

func TestSeries_Order(t *testing.T) {
	expected := "{Integers [3 2 1] int}"
	s := NewSeries([]int{1, 2, 3}, Int, "Integers")
	se := s.Order(2, 1, 0)

	if se.String() != expected {
		t.Errorf("Expected:\n%v\nGot:\n%v", expected, s.String())
	}
	// Output: {Integers [3 2 1] int}
}

func TestSeries_SortedIndex(t *testing.T) {
	expected := []int{1, 2, 0}

	s := NewSeries([]int{3, 1, 2}, Int, "Integers")
	index := s.SortedIndex()

	if index[0] != expected[0] || index[1] != expected[1] || index[2] != expected[2] {
		t.Errorf("Expected:\n%v\nGot:\n%v", expected, index)
	}
	// Output: [0 1 2]
}

func TestSeries_String(t *testing.T) {
	s := NewSeries([]int{1, 2, 3}, Int, "Integers")
	expected := "{Integers [1 2 3] int}"

	if s.String() != expected {
		t.Errorf("Expected:\n%v\nGot:\n%v", expected, s.String())
	}
	// Output: {Integers [1 2 3] int}
}

func TestSeries_Val(t *testing.T) {
	s := NewSeries([]int{1, 2, 3}, Int, "Integers")
	v := s.Val(1)

	if v != 2 {
		t.Errorf("Expected:\n%v\nGot:\n%v", 2, v)
	}
	// Output: 2
}

func TestSeries_Type(t *testing.T) {
	s := NewSeries([]int{1, 2, 3}, Int, "Integers")
	typ := s.Type()

	if typ != Int {
		t.Errorf("Expected:\n%v\nGot:\n%v", Int, typ)
	}
	// Output: int

	s = NewSeries([]float64{1.1, 1.2, 1.3}, Float, "Floats")
	typ = s.Type()

	if typ != Float {
		t.Errorf("Expected:\n%v\nGot:\n%v", Float, typ)
	}

	s = NewSeries([]bool{true, false, true}, Boolean, "Booleans")
	typ = s.Type()

	if typ != Boolean {
		t.Errorf("Expected:\n%v\nGot:\n%v", Boolean, typ)
	}
}

func TestNewRangedSeries(t *testing.T) {
	expected := "{Integers [1 2 3 4] int}"
	s := NewRangedSeries(1, 5, Int, "Integers")

	if s.String() != expected {
		t.Errorf("Expected:\n%v\nGot:\n%v", expected, s.String())
	}
	// Output: {Integers [1 2 3 4] int}
}
