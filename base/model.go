package base

import (
	"GoLab/dataframe"
	"GoLab/dataframe/series"
)

type Model interface {
	Fit()
	Predict(df ...dataframe.DataFrame) series.Series
	//Score() float64

	IsClassifier() bool
	IsRegressor() bool
}
