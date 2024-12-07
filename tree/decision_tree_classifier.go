package tree

import (
	"GoLab/base"
	"GoLab/dataframe"
	"GoLab/dataframe/series"
	"fmt"
	"math"
)

// DecisionTreeClassifier is a struct that represents a decision tree classifier
type DecisionTreeClassifier struct {
	criterion string
	maxDepth  int
	// minSamplesSplit int
	// minSamplesLeaf int

	//metric criterionFunction
	tree *DecisionTree
}

// NewDecisionTreeClassifier creates a new DecisionTreeClassifier with default values
func NewDecisionTreeClassifier() *DecisionTreeClassifier {
	return &DecisionTreeClassifier{
		criterion: "gini",
		maxDepth:  -1,
		tree:      nil,
	}
}

// SetCriterion sets the criterion for the DecisionTreeClassifier
func (dtc *DecisionTreeClassifier) SetCriterion(criterion string) {
	if dtc.tree != nil {
		panic(fmt.Errorf("cannot set criterion after fit"))
	}

	possibleCriteria := []string{"gini", "entropy"}

	for _, c := range possibleCriteria {
		if c == criterion {
			dtc.criterion = criterion
			return
		}
	}
	panic(fmt.Errorf("criterion must be one of %v, but got %v", possibleCriteria, criterion))
}

// SetCriterionFromFunction sets the criterion for the DecisionTreeClassifier from a criterion function
func (dtc DecisionTreeClassifier) SetCriterionFromFunction(criterion func()) {
	panic("SetCriterionFromFunction not implemented")
}

// SetMaxDepth sets the maximum depth of the DecisionTreeClassifier
func (dtc *DecisionTreeClassifier) SetMaxDepth(maxDepth int) {
	if dtc.tree != nil {
		panic(fmt.Errorf("cannot set maxDepth after fit"))
	}

	if maxDepth < -1 || maxDepth == 0 { // Maybe maxDepth of 0 should be allowed? Unsure of the implications
		panic(fmt.Errorf("maxDepth must be greater than 0 or -1 for no limit, but got %v", maxDepth))
	}
	dtc.maxDepth = maxDepth
}

// force implementation of Model interface
var _ base.Model = (*DecisionTreeClassifier)(nil)

func (dtc DecisionTreeClassifier) fitBranch(dfX dataframe.DataFrame, dfY series.Series, depth int) *DecisionTree {
	numSamples, _ := dfX.Shape()

	if numSamples == 0 {
		return nil
	}

	// If all samples are the same class, return a leaf node
	if dfY.Homogeneous() {
		return &DecisionTree{
			Leaf:  true,
			Label: dfY.Val(0).(int),
		}
	}

	minimumEntropy := math.Inf(1)
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
				impurity = gini(dfYLeft, dfYRight)
			} else if dtc.criterion == "entropy" {
				impurity = entropy(dfYLeft, dfYRight)
			}

			if impurity < minimumEntropy {
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
	var left, right *DecisionTree
	if dtc.maxDepth == -1 || depth < dtc.maxDepth {
		left = dtc.fitBranch(dfXLeft, dfYLeft, depth+1)
		right = dtc.fitBranch(dfXRight, dfYRight, depth+1)
	}

	return &DecisionTree{
		Leaf:  false,
		Axis:  bestSplitAxis,
		Value: dfX.At(bestSplitPosition, bestSplitAxis).(float64),
		Left:  left,
		Right: right,
	}
}

// Fit fits the DecisionTreeClassifier to the data and creates the DecisionTree
func (dtc *DecisionTreeClassifier) Fit(dfX dataframe.DataFrame, dfY series.Series) {
	// TODO: Implement fit for gini and entropy

	numSamples, numFeatures := dfX.Shape()
	numOutputs := dfY.Len()

	if numSamples != numOutputs {
		panic(fmt.Errorf("number of samples %v and number of outputs %v must be equal", numSamples, numOutputs))
	}

	if numFeatures > 2 {
		panic("fit not implemented for num_samples > 2") // TODO: implement...
	}

	objects := dfX.SelectObjectNames()
	if objects != nil {
		panic(fmt.Errorf("cannot fit with object columns %v", dfX.SelectObjectNames()))
	}

	dtc.tree = dtc.fitBranch(dfX, dfY, 1)
}

// Predict predicts the target values of the given dataframe.DataFrame
func (dtc DecisionTreeClassifier) Predict(df ...dataframe.DataFrame) series.Series {
	// TODO: Implement predict on DecisionTreeClassifier
	panic("predict not implemented")
}

// IsClassifier returns true as DecisionTreeClassifier is a classifier
func (dtc DecisionTreeClassifier) IsClassifier() bool {
	return true
}

// IsRegressor returns true as DecisionTreeClassifier is a not regressor
func (dtc DecisionTreeClassifier) IsRegressor() bool {
	return false
}
