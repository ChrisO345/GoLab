package base

import (
	"GoLab/dataframe"
	"GoLab/dataframe/series"
)

type Model interface {
	Fit(dfX dataframe.DataFrame, dfY series.Series)
	Predict(df ...dataframe.DataFrame) series.Series
	//Score() float64

	IsClassifier() bool
	IsRegressor() bool
}

//type ProbabilisticClassifier interface {
//	Model
//
//	PredictProbability(df ...dataframe.DataFrame) series.Series
//}
