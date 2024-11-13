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

func New(se ...series.Series) DataFrame {
	if se == nil || len(se) == 0 {
		panic("empty Series")
	}

	// Create index
	indices := make([]int, se[0].Len())
	for i := 0; i < se[0].Len(); i++ {
		indices[i] = i
	}

	index := series.New(indices, series.Int, "Index")

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

func SetIndex(df DataFrame, s series.Series) DataFrame {
	if df.nrows != s.Len() {
		panic(fmt.Errorf("index length %v does not match DataFrame length %v", s.Len(), df.nrows))
	}

	df.index = s.Copy()
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

func (df DataFrame) Shape() (int, int) {
	return df.nrows, df.ncols
}

func (df DataFrame) String() string {
	var sb strings.Builder

	maxIndexOffset := len(fmt.Sprint(df.index.Val(df.nrows - 1)))

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

func (df DataFrame) Head(n ...int) DataFrame {
	if len(n) > 1 {
		panic("only one argument allowed")
	}
	if n == nil || len(n) == 0 {
		n = []int{5}
	}

	var s []series.Series
	for _, se := range df.columns {
		s = append(s, se.Head(n[0]))
	}

	dfNew := New(s...)
	dfNew.index = df.index.Head(n[0])
	return dfNew
}

func (df DataFrame) Tail(n ...int) DataFrame {
	if len(n) > 1 {
		panic("only one argument allowed")
	}
	if n == nil || len(n) == 0 {
		n = []int{5}
	}

	var s []series.Series
	for _, se := range df.columns {
		s = append(s, se.Tail(n[0]))
	}

	dfNew := New(s...)
	dfNew.index = df.index.Tail(n[0])
	return dfNew
}
