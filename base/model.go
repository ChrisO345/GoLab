package base

import (
	"GoLab/dataframe"
	"GoLab/dataframe/series"
)

// Model defines the interface for the all machine learning models
type Model interface {
	Fit(dfX dataframe.DataFrame, dfY series.Series)
	Predict(df ...dataframe.DataFrame) series.Series
	//Score() float64

	IsClassifier() bool
	IsRegressor() bool
}

// ProbabilisticClassifier defines the interface for the all probabilistic classifiers
type ProbabilisticClassifier interface {
	// Model interface implementation
	Model

	PredictProbability(df ...dataframe.DataFrame) series.Series
}
