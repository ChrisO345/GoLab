package dataframe

import (
	"GoLab/dataframe/series"
	"testing"
)

func TestDataFrame_New_Int(t *testing.T) {
	expected := "   Integers1  Integers2\n0          1          4\n1          2          5\n2          3          6"

	df := New(
		series.New([]int{1, 2, 3}, series.Int, "Integers1"),
		series.New([]int{4, 5, 6}, series.Int, "Integers2"),
	)

	if df.String() != expected {
		t.Errorf("Expected:\n%v\nGot:\n%v", expected, df.String())
	}

	// Expected:
	//	   Integers1  Integers2
	//	0          1          4
	//	1          2          5
	//	2          3          6
}

func TestDataFrame_New_Float(t *testing.T) {
	expected := "   Floats1  Floats2\n0      1.1      4.4\n1      2.2      5.5\n2      3.3      6.6"

	df := New(
		series.New([]float64{1.1, 2.2, 3.3}, series.Float, "Floats1"),
		series.New([]float64{4.4, 5.5, 6.6}, series.Float, "Floats2"),
	)

	if df.String() != expected {
		t.Errorf("Expected:\n%v\nGot:\n%v", expected, df.String())
	}
}

func TestDataFrame_New_Mixed(t *testing.T) {
	expected := "   Integers  Floats\n0         1     4.4\n1         2     5.5\n2         3     6.6"

	df := New(
		series.New([]int{1, 2, 3}, series.Int, "Integers"),
		series.New([]float64{4.4, 5.5, 6.6}, series.Float, "Floats"),
	)

	if df.String() != expected {
		t.Errorf("Expected:\n%v\nGot:\n%v", expected, df.String())
	}
}

func TestDataFrame_Shape(t *testing.T) {
	df := New(
		series.New([]int{1, 2, 3}, series.Int, "Integers"),
		series.New([]float64{4.4, 5.5, 6.6}, series.Float, "Floats"),
	)

	rows, cols := df.Shape()

	if rows != 3 || cols != 2 {
		t.Errorf("Expected (3, 2), got (%v, %v)", rows, cols)
	}
}

func TestDataFrame_Head(t *testing.T) {
	expected := "   Integers  Floats  Integers2\n0         1     4.4          7\n1         2     5.5          8"

	df := New(
		series.New([]int{1, 2, 3}, series.Int, "Integers"),
		series.New([]float64{4.4, 5.5, 6.6}, series.Float, "Floats"),
		series.New([]int{7, 8, 9}, series.Int, "Integers2"),
	)
	result := df.Head(2).String()

	if result != expected {
		t.Errorf("Expected:\n%v\nGot:\n%v", expected, result)
	}
}

func TestDataFrame_Tail(t *testing.T) {
	expected := "   Integers  Floats  Integers2\n1         2     5.5          8\n2         3     6.6          9"

	df := New(
		series.New([]int{1, 2, 3}, series.Int, "Integers"),
		series.New([]float64{4.4, 5.5, 6.6}, series.Float, "Floats"),
		series.New([]int{7, 8, 9}, series.Int, "Integers2"),
	)
	result := df.Tail(2).String()

	if result != expected {
		t.Errorf("Expected:\n%v\nGot:\n%v", expected, result)
	}
}