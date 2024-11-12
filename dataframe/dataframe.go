package dataframe

import (
	"fmt"
	"strings"
)

type DataFrame struct {
	columns []Series
	ncols   int
	nrows   int

	Err error
}

func NewDataFrame(se ...Series) DataFrame {
	if se == nil || len(se) == 0 {
		return DataFrame{Err: fmt.Errorf("empty DataFrame")}
	}

	columns := make([]Series, len(se))
	for i, s := range se {
		columns[i] = s.Copy()
	}
	ncols, nrows, err := checkColumnDimensions(columns...)
	if err != nil {
		return DataFrame{Err: err}
	}

	df := DataFrame{
		columns: columns,
		ncols:   ncols,
		nrows:   nrows,
	}

	// Currently assuming that column names are unique and non-empty

	return df
}

func checkColumnDimensions(se ...Series) (ncols int, nrows int, err error) {
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

	maxIndexOffset := len(fmt.Sprint(df.nrows - 1))

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
		indexOffset := maxIndexOffset - len(fmt.Sprint(i))
		for j := 0; j < indexOffset; j++ {
			sb.WriteString(" ")
		}
		sb.WriteString(fmt.Sprint(i))
		sb.WriteString("  ")

		for _, s := range df.columns {
			// Get length of column name
			target := fmt.Sprint(s.elements.Elem(i).Get())

			for j := 0; j < len(s.Name) - len(target); j++ {
				sb.WriteString(" ")
			}
			sb.WriteString(target)
			sb.WriteString("  ")
		}
		sb.WriteString("\n")
	}

	return sb.String()
}
