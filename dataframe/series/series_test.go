package series

import (
	"testing"
)

func TestNewSeriesInt(t *testing.T) {
	expected := "{Integers [1 2 3] int}"
	s := New([]int{1, 2, 3}, Int, "Integers")

	if s.String() != expected {
		t.Errorf("Expected:\n%v\nGot:\n%v", expected, s.String())
	}
	// Output: {Integers [1 2 3] int}
}

func TestNewSeriesFloat(t *testing.T) {
	expected := "{Floats [1.1 2.2 3.3] float}"
	s := New([]float64{1.1, 2.2, 3.3}, Float, "Floats")

	if s.String() != expected {
		t.Errorf("Expected:\n%v\nGot:\n%v", expected, s.String())
	}
	// Output: {Floats [1.1 2.2 3.3] float}
}
