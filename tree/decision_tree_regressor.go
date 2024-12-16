package tree

import (
	"GoLab"
	"GoLab/dataframe"
	"GoLab/dataframe/series"
	"fmt"
)

// DecisionTreeRegressor is a struct that represents a decision tree classifier
type DecisionTreeRegressor struct {
	maxDepth int
	// minSamplesSplit int
	// minSamplesLeaf int

	criterionString string
	criterion       criterionFunction
	tree            *DecisionTree

	features []string
	target   string
}

// force implementation of Model interface
var _ GoLab.Model = (*DecisionTreeRegressor)(nil)

// NewDecisionTreeRegressor creates a new DecisionTreeRegressor with default values
func NewDecisionTreeRegressor() *DecisionTreeRegressor {
	return &DecisionTreeRegressor{
		criterionString: "mse",
		//criterion:       mse, // TODO implement me
		maxDepth:        -1,
		tree:            nil,
	}
}

// SetCriterion sets the criterion for the DecisionTreeRegressor
func (dtr *DecisionTreeRegressor) SetCriterion(criterion string) {
	if dtr.tree != nil {
		panic(fmt.Errorf("cannot set criterion after fit"))
	}

	criterionStrings := []string{"mse"}
	possibleCriteria := map[string]criterionFunction{
		//"mse": mse, // TODO implement me
	}

	for k, c := range possibleCriteria {
		if k == criterion {
			dtr.criterion = c
			dtr.criterionString = criterion
			return
		}
	}

	panic(fmt.Errorf("criterion must be one of %v, but got %v", criterionStrings, criterion))
}

// SetCriterionFunction sets the criterion for the DecisionTreeRegressor from a criterion function
func (dtr DecisionTreeRegressor) SetCriterionFunction(criterion criterionFunction) {
	if dtr.tree != nil {
		panic(fmt.Errorf("cannot set criterion after fit"))
	}

	dtr.criterion = criterion
	dtr.criterionString = "custom"
}

// SetMaxDepth sets the maximum depth of the DecisionTreeRegressor
func (dtr *DecisionTreeRegressor) SetMaxDepth(maxDepth int) {
	if dtr.tree != nil {
		panic(fmt.Errorf("cannot set maxDepth after fit"))
	}

	if maxDepth < -1 || maxDepth == 0 { // Maybe maxDepth of 0 should be allowed? Unsure of the implications
		panic(fmt.Errorf("maxDepth must be greater than 0 or -1 for no limit, but got %v", maxDepth))
	}

	dtr.maxDepth = maxDepth
}

// Fit fits the DecisionTreeRegressor to the data
func (dtr DecisionTreeRegressor) Fit(dfX dataframe.DataFrame, dfY series.Series) {
	//TODO implement me
	panic("implement me")
}

// Predict predicts the target values for the given data
func (dtr DecisionTreeRegressor) Predict(df dataframe.DataFrame) series.Series {
	//TODO implement me
	panic("implement me")
}

// IsClassifier returns whether the model is a classifier
func (dtr DecisionTreeRegressor) IsClassifier() bool {
	return false
}

// IsRegressor returns whether the model is a regressor
func (dtr DecisionTreeRegressor) IsRegressor() bool {
	return true
}
