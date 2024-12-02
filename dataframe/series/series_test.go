package series

import (
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
