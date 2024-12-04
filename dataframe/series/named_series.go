package series

// NewRangedSeries creates a new Series defined for a range of integers.
func NewRangedSeries(start, end int, t Type, name string) Series {
	numRange := make([]int, end-start)
	for i := start; i < end; i++ {
		numRange[i-start] = i
	}

	return NewSeries(numRange, t, name)
}