package tree

import (
	"GoLab/base"
	"GoLab/dataframe"
	"GoLab/dataframe/series"
	"fmt"
)

type DecisionTree struct {

}

type DecisionTreeClassifier struct {
	criterion string
	maxDepth  int

	tree *DecisionTree
}

func NewDecisionTreeClassifier() *DecisionTreeClassifier {
	return &DecisionTreeClassifier{
		criterion: "gini",
		maxDepth:  -1,
		tree: nil,
	}
}

func (dtc *DecisionTreeClassifier) SetCriterion(criterion string) {
	if dtc.tree != nil {
		panic("cannot set criterion after fit")
	}

	possibleCriteria := []string{"gini", "entropy"}

	for _, c := range possibleCriteria {
		if c == criterion {
			dtc.criterion = criterion
			return
		}
	}
	panic(fmt.Sprintf("invalid criterion, must be one of %v", possibleCriteria))
}

func (dtc *DecisionTreeClassifier) SetMaxDepth(maxDepth int) {
	if dtc.tree != nil {
		panic("cannot set maxDepth after fit")
	}

	if maxDepth < -1 || maxDepth == 0 { // Maybe maxDepth of 0 should be allowed? Unsure of the implications
		panic("incorrect depth specified")
	}
	dtc.maxDepth = maxDepth
}

// force implementation of Model interface
var _ base.Model = (*DecisionTreeClassifier)(nil)

func (dtc DecisionTreeClassifier) Fit() {
	// TODO: Implement fit for gini and entropy
	panic("fit not implemented")
}

func (dtc DecisionTreeClassifier) Predict(df ...dataframe.DataFrame) series.Series {
	// TODO: Implement predict on DecisionTreeClassifier
	panic("predict not implemented")
}

func (dtc DecisionTreeClassifier) IsClassifier() bool {
	return true
}

func (dtc DecisionTreeClassifier) IsRegressor() bool {
	return false
}
