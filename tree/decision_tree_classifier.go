package tree

import (
	"GoLab/base"
	"GoLab/dataframe"
	"GoLab/dataframe/series"
	"fmt"
	"math"
)

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

func (dtc DecisionTreeClassifier) fitBranch(dfX dataframe.DataFrame, dfY series.Series) *DecisionTree {
	numSamples, _ := dfX.Shape()

	if numSamples == 0 {
		return nil
	}

	// If all samples are the same class, return a leaf node
	if dfY.NUnique() {
		return &DecisionTree{
			Leaf: true,
			Label: dfY.Val(0).(int),
		}
	}

	minimumEntropy := math.Inf(-1)
	bestSplitAxis := 0
	bestSplitPosition := 0

	for axis, column := range dfX.Columns() {
		order := column.SortedIndex()
		dfX = dfX.Order(order...)
		dfY = dfY.Order(order...)

		for i := 0; i < numSamples; i++ {
			dfYLeft := dfY.Slice(0, i)
			dfYRight := dfY.Slice(i, numSamples)

			var impurity float64
			if dtc.criterion == "gini" {
				panic("Gini Not Implemented")
			} else if dtc.criterion == "entropy" {
				impurity = entropy(dfYLeft, dfYRight)
			}

			if impurity > minimumEntropy {
				minimumEntropy = impurity
				bestSplitAxis = axis
				bestSplitPosition = i
			}
		}
	}

	// Sort by best axis
	order := dfX.Columns()[bestSplitAxis].SortedIndex()
	dfX = dfX.Order(order...)
	dfY = dfY.Order(order...)

	// Split the data
	dfXLeft := dfX.Slice(0, bestSplitPosition)
	dfYLeft := dfY.Slice(0, bestSplitPosition)
	dfXRight := dfX.Slice(bestSplitPosition, numSamples)
	dfYRight := dfY.Slice(bestSplitPosition, numSamples)

	// Recursively fit the Left and Right branches
	left := dtc.fitBranch(dfXLeft, dfYLeft)
	right := dtc.fitBranch(dfXRight, dfYRight)

	return &DecisionTree{
		Leaf: false,
		Axis: bestSplitAxis,
		Value: dfX.At(bestSplitPosition, bestSplitAxis).(float64),
		Left: left,
		Right: right,
	}
}

func (dtc *DecisionTreeClassifier) Fit(dfX dataframe.DataFrame, dfY series.Series) {
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

	dtc.tree = dtc.fitBranch(dfX, dfY)
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