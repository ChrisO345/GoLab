package dataframe

import (
	"testing"
)

func TestNewSeries(t *testing.T) {
	expected := "{Integers [1 2 3] int}"
	s := NewSeries([]int{1, 2, 3}, Int, "Integers")

	if s.String() != expected {
		t.Errorf("Expected:\n%v\nGot:\n%v", expected, s.String())
	}
	// Output: {Integers [1 2 3] int}
}
