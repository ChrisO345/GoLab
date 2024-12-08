package dataframe

import (
	"GoLab/dataframe/series"
	"fmt"
	"strings"
)

// DataFrame is a collection of series.Series with a shared index.
// It is similar to a table in a relational database, and is implemented
// similar to a dataframe in R or Python (pandas).
type DataFrame struct {
	index   series.Series
	columns []series.Series
	ncols   int
	nrows   int
}

// NewDataFrame creates a new DataFrame from a collection of series.Series.
// It has a shared index which defaults to a range of integers.
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

	// TODO: Currently assuming that column names are unique

	return df
}

// checkColumnDimensions checks that all series.Series have the same length.
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

// String is the Stringer implementation for DataFrame.
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

// Shape returns the dimensions of the DataFrame in the form (nrows, ncols).
func (df DataFrame) Shape() (int, int) {
	return df.nrows, df.ncols
}

// Columns returns a collection of the series.Series of the DataFrame.
func (df DataFrame) Columns() []series.Series {
	return df.columns
}

// Column returns a series.Series of the DataFrame by name.
func (df DataFrame) Column(name string) *series.Series {
	for _, s := range df.columns {
		if s.Name == name {
			return &s
		}
	}
	panic(fmt.Errorf("column %v not found", name))
}

// Names returns a collection of the names of the series.Series of the DataFrame.
func (df DataFrame) Names() []string {
	names := make([]string, df.ncols)
	for i, s := range df.columns {
		names[i] = s.Name
	}
	return names
}

// SetIndex sets the index of the DataFrame to a specified series.Series.
func (df DataFrame) SetIndex(s series.Series) DataFrame {
	if df.nrows != s.Len() {
		panic(fmt.Errorf("index length %v does not match DataFrame length %v", s.Len(), df.nrows))
	}

	df.index = s.Copy()
	return df
}

// ResetIndex resets the index of the DataFrame to a range of integers.
func (df DataFrame) ResetIndex() DataFrame {
	indices := make([]int, df.nrows)
	for i := 0; i < df.nrows; i++ {
		indices[i] = i
	}

	df.index = series.NewSeries(indices, series.Int, "Index")
	return df
}

// Index returns the index of the DataFrame.
func (df DataFrame) Index() series.Series {
	return df.index
}

// Slice returns a new DataFrame with rows from a to b
func (df DataFrame) Slice(a, b int) DataFrame{
	if a < 0 || b > df.nrows {
		panic(fmt.Errorf("b index %v out of range", b))
	}

	if a > b {
		panic(fmt.Errorf("a index %v greater than b index %v", a, b))
	}

	var s []series.Series
	for _, se := range df.columns {
		s = append(s, se.Slice(a, b))
	}

	dfNew := NewDataFrame(s...)
	dfNew.index = df.index.Slice(a, b)
	return dfNew
}

// Head returns a slice of the last n elements of the DataFrame. If n is not specified, it defaults to 5.
func (df DataFrame) Head(n ...int) DataFrame {
	if len(n) > 1 {
		panic("only one argument allowed")
	}
	if n == nil || len(n) == 0 {
		n = []int{5}
	}

	return df.Slice(0, n[0])
}

// Tail returns a slice of the last n elements of the DataFrame. If n is not specified, it defaults to 5.
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

// At returns the value at the specified row and column of the DataFrame.
func (df DataFrame) At(i, j int) interface{} {
	if i < 0 || i >= df.nrows {
		panic(fmt.Errorf("index %v out of range", i))
	}
	if j < 0 || j >= df.ncols {
		panic(fmt.Errorf("column %v out of range", j))
	}

	return df.columns[j].Val(i)
}

// Swap swaps the rows at index row1 and row2 of the DataFrame inplace.
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

// Sort sorts the DataFrame inplace according to the specified columns.
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

// Order orders the DataFrame inplace according to the specified positions.
func (df DataFrame) Order(positions ...int) DataFrame {
	if len(positions) != df.nrows {
		panic("positions must be the same length as the DataFrame")
	}

	// Need to copy otherwise positions collection will mutate
	newPositions := make([]int, df.nrows)
	for i, pos := range positions {
		newPositions[i] = pos
	}

	for newPos, oldPos := range newPositions {
		if oldPos == newPos {
			continue
		}

		df.Swap(oldPos, newPos)

		for i, pos := range newPositions {
			if pos == newPos {
				newPositions[i] = oldPos
				newPositions[newPos] = newPos
				break
			}
		}
	}

	return df
}

// Append appends a series.Series to right of the DataFrame.
func (df *DataFrame) Append(s series.Series) {
	if s.Len() != df.nrows {
		panic(fmt.Errorf("series length %v does not match DataFrame length %v", s.Len(), df.nrows))
	}

	df.columns = append(df.columns, s)
	df.ncols++
}

// Copy returns a deep copy of the DataFrame.
func (df DataFrame) Copy() DataFrame {
	var s []series.Series
	for _, se := range df.columns {
		s = append(s, se.Copy())
	}

	dfNew := NewDataFrame(s...)
	dfNew.index = df.index.Copy()
	return dfNew
}

// SelectObjectNames returns a collection of the names of object columns of the DataFrame.
func (df DataFrame) SelectObjectNames() []string {
	var objects []string
	for _, s := range df.columns {
		if s.IsObject() {
			objects = append(objects, s.Name)
		}
	}
	return objects
}

// SelectNumericNames returns a collection of the names of numeric columns of the DataFrame.
func (df DataFrame) SelectNumericNames() []string {
	var objects []string
	for _, s := range df.columns {
		if s.IsNumeric() {
			objects = append(objects, s.Name)
		}
	}
	return objects
}

// Drop removes the specified column from the DataFrame and returns it as a series.Series.
func (df *DataFrame) Drop(name string) series.Series {
	for i, s := range df.columns {
		if s.Name == name {
			df.columns = append(df.columns[:i], df.columns[i+1:]...)
			df.ncols--
			return s
		}
	}
	panic(fmt.Errorf("column %v not found", name))
}

