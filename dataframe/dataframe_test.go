package dataframe

import (
	"GoLab/dataframe/series"
	"testing"
)

func TestDataFrame_New_Int(t *testing.T) {
	expected := "   Integers1  Integers2\n0          1          4\n1          2          5\n2          3          6"

	df := NewDataFrame(
		series.NewSeries([]int{1, 2, 3}, series.Int, "Integers1"),
		series.NewSeries([]int{4, 5, 6}, series.Int, "Integers2"),
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

	df := NewDataFrame(
		series.NewSeries([]float64{1.1, 2.2, 3.3}, series.Float, "Floats1"),
		series.NewSeries([]float64{4.4, 5.5, 6.6}, series.Float, "Floats2"),
	)

	if df.String() != expected {
		t.Errorf("Expected:\n%v\nGot:\n%v", expected, df.String())
	}
}

func TestDataFrame_New_Mixed(t *testing.T) {
	expected := "   Integers  Floats\n0         1     4.4\n1         2     5.5\n2         3     6.6"

	df := NewDataFrame(
		series.NewSeries([]int{1, 2, 3}, series.Int, "Integers"),
		series.NewSeries([]float64{4.4, 5.5, 6.6}, series.Float, "Floats"),
	)

	if df.String() != expected {
		t.Errorf("Expected:\n%v\nGot:\n%v", expected, df.String())
	}
}

func TestDataFrame_Shape(t *testing.T) {
	df := NewDataFrame(
		series.NewSeries([]int{1, 2, 3}, series.Int, "Integers"),
		series.NewSeries([]float64{4.4, 5.5, 6.6}, series.Float, "Floats"),
	)

	rows, cols := df.Shape()

	if rows != 3 || cols != 2 {
		t.Errorf("Expected (3, 2), got (%v, %v)", rows, cols)
	}
}

func TestDataFrame_Head(t *testing.T) {
	expected := "   Integers  Floats  Integers2\n0         1     4.4          7\n1         2     5.5          8"

	df := NewDataFrame(
		series.NewSeries([]int{1, 2, 3}, series.Int, "Integers"),
		series.NewSeries([]float64{4.4, 5.5, 6.6}, series.Float, "Floats"),
		series.NewSeries([]int{7, 8, 9}, series.Int, "Integers2"),
	)
	result := df.Head(2).String()

	if result != expected {
		t.Errorf("Expected:\n%v\nGot:\n%v", expected, result)
	}
}

func TestDataFrame_Tail(t *testing.T) {
	expected := "   Integers  Floats  Integers2\n1         2     5.5          8\n2         3     6.6          9"

	df := NewDataFrame(
		series.NewSeries([]int{1, 2, 3}, series.Int, "Integers"),
		series.NewSeries([]float64{4.4, 5.5, 6.6}, series.Float, "Floats"),
		series.NewSeries([]int{7, 8, 9}, series.Int, "Integers2"),
	)
	result := df.Tail(2).String()

	if result != expected {
		t.Errorf("Expected:\n%v\nGot:\n%v", expected, result)
	}
}

func TestDataFrame_SetIndex(t *testing.T) {
	df := NewDataFrame(
		series.NewSeries([]int{1, 2, 3}, series.Int, "Integers"),
		series.NewSeries([]float64{4.4, 5.5, 6.6}, series.Float, "Floats"),
	)

	df = df.SetIndex(series.NewSeries([]int{7, 8, 9}, series.Int, "Integers2"))

	if df.Index().String() != "{Integers2 [7 8 9] int}" {
		t.Errorf("Expected index to be [7, 8, 9], got %v", df.Index().String())
	}
}

func TestDataFrame_ResetIndex(t *testing.T) {
	df := NewDataFrame(
		series.NewSeries([]int{1, 2, 3}, series.Int, "Integers"),
		series.NewSeries([]float64{4.4, 5.5, 6.6}, series.Float, "Floats"),
	)

	df = df.SetIndex(series.NewSeries([]int{7, 8, 9}, series.Int, "Integers2"))
	df = df.ResetIndex()

	if df.Index().String() != "{Index [0 1 2] int}" {
		t.Errorf("Expected index to be [0, 1, 2], got %v", df.Index().String())
	}
}

func TestDataFrame_Columns(t *testing.T) {
	a := series.NewSeries([]int{1, 2, 3}, series.Int, "Integers")
	b := series.NewSeries([]float64{4.4, 5.5, 6.6}, series.Float, "Floats")

	df := NewDataFrame(
		a,
		b,
	)

	cols := df.Columns()

	if len(cols) != 2 {
		t.Errorf("Expected 2 columns, got %v", len(cols))
	}

	if cols[0].String() != a.String() {
		t.Errorf("Expected first column to be %v, got %v", a.String(), cols[0].String())
	}
	if cols[1].String() != b.String() {
		t.Errorf("Expected second column to be %v, got %v", b.String(), cols[1].String())
	}
}

func TestDataFrame_Index(t *testing.T) {
	df := NewDataFrame(
		series.NewSeries([]int{1, 2, 3}, series.Int, "Integers"),
		series.NewSeries([]float64{4.4, 5.5, 6.6}, series.Float, "Floats"),
	)

	if df.Index().String() != "{Index [0 1 2] int}" {
		t.Errorf("Expected index to be nil, got %v", df.Index())
	}
}

func TestDataFrame_Sort(t *testing.T) {
	expected := "   Integers  Floats\n1         1     6.6\n2         2     5.5\n0         3     4.4"

	df := NewDataFrame(
		series.NewSeries([]int{3, 1, 2}, series.Int, "Integers"),
		series.NewSeries([]float64{4.4, 6.6, 5.5}, series.Float, "Floats"),
	)
	df.Sort("Integers")

	if df.String() != expected {
		t.Errorf("Expected:\n%v\nGot:\n%v", expected, df.String())
	}

	expected = "   Integers  Floats\n0         3     4.4\n2         2     5.5\n1         1     6.6"

	df.Sort("Floats")

	if df.String() != expected {
		t.Errorf("Expected:\n%v\nGot:\n%v", expected, df.String())
	}
}

func TestDataFrame_Order(t *testing.T) {
	expected := "   Integers  Floats\n2         3     6.6\n1         2     5.5\n0         1     4.4"

	df := NewDataFrame(
		series.NewSeries([]int{1, 2, 3}, series.Int, "Integers"),
		series.NewSeries([]float64{4.4, 5.5, 6.6}, series.Float, "Floats"),
	)
	df = df.Order(2, 1, 0)

	if df.String() != expected {
		t.Errorf("Expected:\n%v\nGot:\n%v", expected, df.String())
	}
}

func TestDataFrame_Append(t *testing.T) {
	expected := "   Integers  Floats\n0         1     4.4\n1         2     5.5\n2         3     6.6"

	df1 := NewDataFrame(
		series.NewSeries([]int{1, 2, 3}, series.Int, "Integers"),
	)
	s := series.NewSeries([]float64{4.4, 5.5, 6.6}, series.Float, "Floats")

	df1.Append(s)

	if df1.String() != expected {
		t.Errorf("Expected:\n%v\nGot:\n%v", expected, df1.String())
	}
}