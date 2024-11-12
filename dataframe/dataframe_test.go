package dataframe

import (
	"testing"
)

func TestNewDataFrame(t *testing.T) {
	expected := "   Integers1  Integers2\n0          1          4  \n1          2          5  \n2          3          6  \n"

	df := NewDataFrame(
		NewSeries([]int{1, 2, 3}, Int, "Integers1"),
		NewSeries([]int{4, 5, 6}, Int, "Integers2"),
	)
	if df.Err != nil {
		t.Error(df.Err)
	}
	if df.String() != expected {
		t.Errorf("Expected:\n%v\nGot:\n%v", expected, df.String())
	}

	// Expected:
	//	   Integers1  Integers2
	//	0          1          4
	//	1          2          5
	//	2          3          6
}
