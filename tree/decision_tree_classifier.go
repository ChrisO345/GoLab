package tree

import (
	"GoLab/base"
	"GoLab/dataframe"
	"GoLab/dataframe/series"
	"fmt"
	"math"
)

type DecisionTree struct { // TODO: Move to another file. Should this be private?

}

type DecisionTreeClassifier struct {
	criterion string
	maxDepth  int
	// minSamplesSplit int
	// minSamplesLeaf int

	//metric criterionFunction
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

func (dtc DecisionTreeClassifier) SetCriterionFromFunction(criterion func()) {
	panic("SetCriterionFromFunction not implemented")
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

func (dtc DecisionTreeClassifier) Fit(dfX dataframe.DataFrame, dfY series.Series) {
	// TODO: Implement fit for gini and entropy

	numSamples, numFeatures := dfX.Shape()
	numOutputs := dfY.Len()

	if numSamples != numOutputs {
		panic("number of observations must be the same") // TODO: improve error message
	}

	if numFeatures > 2 {
		panic("fit not implemented for num_samples > 2") // TODO: implement...
	}

	for _, columns := range dfX.Columns() {
		if !columns.IsNumeric() {
			panic("data must be numeric") // TODO: improve error message
		}
	}

	// Split along each feature, calculate the gini/entropy, and choose the best split TODO: FUNCTION
	// Q_left  <= ...
	// Q_right >  ...

	optimalSplit := math.Inf(1)
	bestAxis := -1
	bestPosition := -1

	for axis, column := range dfX.Columns() {
		order := column.SortedIndex()

		dfX = dfX.Order(order...)
		dfY = dfY.Order(order...)

		for i := 0; i < numSamples - 1; i++ {
			fmt.Println(i)
			// Don't split if the values are the same
			current := dfX.At(i, 0)
			next := dfX.At(i+1, 0)
			if current == next {
				continue
			}

			dfLeftX := dfX.Slice(0, i)
			dfLeftY := dfY.Slice(0, i)
			dfRightX := dfX.Slice(i+1, numSamples)
			dfRightY := dfY.Slice(i+1, numSamples)

			// Calculate the split TODO: fix for more than 2 features
			impurity := math.Inf(1)
			if dtc.criterion == "gini" {
				// Calculate the gini
				// gini(dfX, dfY)
			} else if dtc.criterion == "entropy" {
				// Calculate the entropy
				impurity = -entropy(dfLeftX, dfLeftY, dfRightX, dfRightY)
			}
			fmt.Println(impurity)

			if impurity < optimalSplit {
				optimalSplit = impurity
				bestAxis = axis
				bestPosition = i
			}
		}
	}
	order := dfX.Columns()[bestAxis].SortedIndex()
	dfX = dfX.Order(order...)
	dfY = dfY.Order(order...)

	fmt.Println(bestAxis, bestPosition)
	fmt.Println(dfX.At(bestPosition, 0), dfX.At(bestPosition, 1))
	fmt.Println(optimalSplit)

	fmt.Println(dfX.String())

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