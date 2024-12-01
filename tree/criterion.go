package tree

import (
	"GoLab/dataframe"
	"GoLab/dataframe/series"
)

type criterionFunction func(dataframe.DataFrame) float64

func gini(frame dataframe.DataFrame) float64 {
	panic("gini not implemented")
}

func entropy(frame dataframe.DataFrame, series series.Series) float64 {
	panic("entropy not implemented")
}
