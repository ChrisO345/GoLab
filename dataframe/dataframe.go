package dataframe

import (
	"GoLab/dataframe/series"
	"fmt"
	"strings"
)

type DataFrame struct {
	index   series.Series
	columns []series.Series
	ncols   int
	nrows   int
}

func NewDataFrame(se ...series.Series) DataFrame {
	if se == nil || len(se) == 0 {
		panic("empty Series")
	}

	// Create index
	indices := make([]int, se[0].Len())
	for i := 0; i < se[0].Len(); i++ {
		indices[i] = i
	}

	index := series.NewSeries(indices, series.Int, "Index")

	columns := make([]series.Series, len(se))
	for i, s := range se {
		columns[i] = s.Copy()
	}
	ncols, nrows, err := checkColumnDimensions(columns...)
	if err != nil {
		panic(err)
	}

	df := DataFrame{
		index:   index,
		columns: columns,
		ncols:   ncols,
		nrows:   nrows,
	}

	// Currently assuming that column names are unique and non-empty

	return df
}

func checkColumnDimensions(se ...series.Series) (ncols int, nrows int, err error) {
	ncols = len(se)
	nrows = -1
	if se == nil || ncols == 0 {
		err = fmt.Errorf("empty Series")
		return
	}

	for i, s := range se {
		if nrows == -1 {
			nrows = s.Len()
		} else if nrows != s.Len() {
			err = fmt.Errorf("series %v has length %v, expected %v", i, s.Len(), nrows)
			return
		}
	}
	return
}

func (df DataFrame) String() string {
	var sb strings.Builder

	maxIndexOffset := len(fmt.Sprint(df.index.Val(df.nrows - 1)))
	for i := 0; i < df.index.Len(); i++ {
		temp := len(fmt.Sprint(df.index.Val(i)))
		if temp > maxIndexOffset {
			maxIndexOffset = temp
		}
	}

	for i, s := range df.columns {
		if i == 0 {
			for j := 0; j < maxIndexOffset; j++ {
				sb.WriteString(" ")
			}
		}
		sb.WriteString("  ")
		sb.WriteString(s.Name)
	}

	sb.WriteString("\n")
	// TODO: fix formatting implementation to be tidier
	for i := 0; i < df.nrows; i++ {
		indexOffset := maxIndexOffset - len(fmt.Sprint(df.index.Val(i)))
		for j := 0; j < indexOffset; j++ {
			sb.WriteString(" ")
		}
		sb.WriteString(fmt.Sprint(df.index.Val(i)))
		sb.WriteString("  ")

		for k, s := range df.columns {
			// Get length of column
			target := fmt.Sprint(s.Val(i))

			for j := 0; j < len(s.Name)-len(target); j++ {
				sb.WriteString(" ")
			}
			sb.WriteString(target)
			if k < df.ncols-1 {
				sb.WriteString("  ")
			}
		}
		if i < df.nrows-1 {
			sb.WriteString("\n")
		}
	}

	return sb.String()
}

func (df DataFrame) Shape() (int, int) {
	return df.nrows, df.ncols
}

func (df DataFrame) Columns() []series.Series { // TODO: change to a closure?? be value of DataFrame??
	return df.columns
}

// TODO: allow indexing by both strings and integers
//type indexer interface {
//	string | int
//}

func (df DataFrame) Column(name string) *series.Series {
	for _, s := range df.columns {
		if s.Name == name {
			return &s
		}
	}
	panic(fmt.Errorf("column %v not found", name))
}

func (df DataFrame) SetIndex(s series.Series) DataFrame {
	if df.nrows != s.Len() {
		panic(fmt.Errorf("index length %v does not match DataFrame length %v", s.Len(), df.nrows))
	}

	df.index = s.Copy()
	return df
}

func (df DataFrame) ResetIndex() DataFrame {
	indices := make([]int, df.nrows)
	for i := 0; i < df.nrows; i++ {
		indices[i] = i
	}

	df.index = series.NewSeries(indices, series.Int, "Index")
	return df
}

func (df DataFrame) Index() series.Series {
	return df.index
}

func (df DataFrame) Slice(from, to int) DataFrame{
	if from < 0 || to > df.nrows {
		panic(fmt.Errorf("to index %v out of range", to))
	}

	if from > to {
		panic(fmt.Errorf("from index %v greater than to index %v", from, to))
	}

	var s []series.Series
	for _, se := range df.columns {
		s = append(s, se.Slice(from, to))
	}

	dfNew := NewDataFrame(s...)
	dfNew.index = df.index.Slice(from, to)
	return dfNew
}

func (df DataFrame) Head(n ...int) DataFrame {
	if len(n) > 1 {
		panic("only one argument allowed")
	}
	if n == nil || len(n) == 0 {
		n = []int{5}
	}

	return df.Slice(0, n[0])
}

func (df DataFrame) Tail(n ...int) DataFrame {
	if len(n) > 1 {
		panic("only one argument allowed")
	}
	if n == nil || len(n) == 0 {
		n = []int{5}
	}

	fmt.Println(df.nrows, n[0])

	return df.Slice(df.nrows-n[0], df.nrows)
}

func (df DataFrame) At(i, j int) interface{} {
	if i < 0 || i >= df.nrows {
		panic(fmt.Errorf("index %v out of range", i))
	}
	if j < 0 || j >= df.ncols {
		panic(fmt.Errorf("column %v out of range", j))
	}

	return df.columns[j].Val(i)
}

func (df DataFrame) Swap(row1, row2 int) {
	// Swap index
	temp := df.index.Val(row1)
	df.index.Elem(row1).Set(df.index.Val(row2))
	df.index.Elem(row2).Set(temp)

	// Swap rows
	for k := 0; k < df.ncols; k++ {
		temp := df.columns[k].Val(row1)
		df.columns[k].Elem(row1).Set(df.columns[k].Val(row2))
		df.columns[k].Elem(row2).Set(temp)
	}
}

func (df DataFrame) Sort(columns ...string) {
	if len(columns) == 0 {
		panic("no columns specified")
	}

	if len(columns) > 1 {
		panic("> 1 column not yet implemented")
	}

	column := df.Column(columns[0])

	// Sort via bubble sort according to specified column
	for i := 0; i < df.nrows; i++ {
		for j := 0; j < df.nrows-i-1; j++ {
			switch column.Type() {
			case series.Int:
				if column.Val(j).(int) > column.Val(j+1).(int) {
					df.Swap(j, j+1)
				}
			case series.Float:
				if column.Val(j).(float64) > column.Val(j+1).(float64) {
					df.Swap(j, j+1)
				}
			}
		}
	}
}

func (df DataFrame) Order(positions ...int) DataFrame{
	if len(positions) != df.nrows {
		panic("positions must be the same length as the DataFrame")
	}

	for newPos, oldPos := range positions {
		if oldPos == newPos {
			continue
		}

		df.Swap(oldPos, newPos)

		for i, pos := range positions {
			if pos == newPos {
				positions[i] = oldPos
				positions[newPos] = newPos
				break
			}
		}
	}

	return df
}

func (df *DataFrame) Append(s series.Series) {
	if s.Len() != df.nrows {
		panic(fmt.Errorf("series length %v does not match DataFrame length %v", s.Len(), df.nrows))
	}

	df.columns = append(df.columns, s)
	df.ncols++
}

func (df DataFrame) Copy() DataFrame {
	var s []series.Series
	for _, se := range df.columns {
		s = append(s, se.Copy())
	}

	dfNew := NewDataFrame(s...)
	dfNew.index = df.index.Copy()
	return dfNew
}
