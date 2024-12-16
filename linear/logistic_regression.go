package linear

import (
	"GoLab"
	"GoLab/dataframe"
	"GoLab/dataframe/series"
)

type LogisticRegression struct {
	penalty string
	C       float64
	solver  string
}

func NewLogisticRegression() *LogisticRegression {
	return &LogisticRegression{
		penalty: "l2",
		C:       1.0,
		solver:  "lbfgs",
	}
}

// force implementation of Model interface
var _ GoLab.Model = (*LogisticRegression)(nil)

func (lr LogisticRegression) Fit(dfX dataframe.DataFrame, dfY series.Series) {
	// TODO: Implement fit for logistic regression
	panic("fit not implemented")
}

func (lr LogisticRegression) Predict(df dataframe.DataFrame) series.Series {
	// TODO: Implement predict for logistic regression
	panic("predict not implemented")
}

func (lr LogisticRegression) IsClassifier() bool {
	return false
}

func (lr LogisticRegression) IsRegressor() bool {
	return true
}
