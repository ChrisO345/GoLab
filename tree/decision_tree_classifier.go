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
	maxDepth int
	// minSamplesSplit int
	// minSamplesLeaf int

	criterionString string
	criterion       criterionFunction
	tree            *DecisionTree

	features []string
	target   string
}

// NewDecisionTreeClassifier creates a new DecisionTreeClassifier with default values
func NewDecisionTreeClassifier() *DecisionTreeClassifier {
	return &DecisionTreeClassifier{
		criterionString: "gini",
		criterion:       gini,
		maxDepth:        -1,
		tree:            nil,
	}
}

// SetCriterion sets the criterion for the DecisionTreeClassifier
func (dtc *DecisionTreeClassifier) SetCriterion(criterion string) {
	if dtc.tree != nil {
		panic(fmt.Errorf("cannot set criterion after fit"))
	}

	criterionStrings := []string{"gini", "entropy"}
	possibleCriteria := map[string]criterionFunction{
		"gini":    gini,
		"entropy": entropy,
	}

	for k, c := range possibleCriteria {
		if k == criterion {
			dtc.criterion = c
			dtc.criterionString = criterion
			return
		}
	}

	panic(fmt.Errorf("criterion must be one of %v, but got %v", criterionStrings, criterion))
}

// SetCriterionFunction sets the criterion for the DecisionTreeClassifier from a criterion function
func (dtc DecisionTreeClassifier) SetCriterionFunction(criterion criterionFunction) {
	if dtc.tree != nil {
		panic(fmt.Errorf("cannot set criterion after fit"))
	}

	dtc.criterionString = "other"
	dtc.criterion = criterion
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

			impurity := dtc.criterion(dfYLeft, dfYRight)

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
	numSamples, _ := dfX.Shape()
	numOutputs := dfY.Len()

	if numSamples != numOutputs {
		panic(fmt.Errorf("number of samples %v and number of outputs %v must be equal", numSamples, numOutputs))
	}

	objects := dfX.SelectObjectNames()
	if objects != nil {
		panic(fmt.Errorf("cannot fit with object columns %v", dfX.SelectObjectNames()))
	}

	dtc.tree = dtc.fitBranch(dfX, dfY, 1)
	dtc.features = dfX.Names()
	dtc.target = dfY.Name
}

func (dtc DecisionTreeClassifier) predict(df dataframe.DataFrame, idx int) int {
	current := dtc.tree
	for current.hasChildren() {
		if df.At(idx, current.Axis).(float64) <= current.Value {
			current = current.Left
		} else {
			current = current.Right
		}
	}

	// If leaf return the label otherwise return the
	if current.Leaf {
		return current.Label
	}

	panic(fmt.Errorf("current node is not a leaf"))
}

// Predict predicts the target values of the given dataframe.DataFrame
func (dtc DecisionTreeClassifier) Predict(df dataframe.DataFrame) series.Series {
	if dtc.tree == nil {
		panic(fmt.Errorf("must fit model before predicting"))
	}

	numSamples, _ := df.Shape()

	for idx, name := range df.Names() {
		if name != dtc.features[idx] {
			panic(fmt.Errorf("column %v does not match fit column %v", name, dtc.features[idx]))
		}
	}

	predictions := make([]int, numSamples)
	for i := 0; i < numSamples; i++ {
		// Slice the dataframe...
		predictions[i] = dtc.predict(df, i)
	}

	return series.New(predictions, series.Int, dtc.target)
}

// IsClassifier returns true as DecisionTreeClassifier is a classifier
func (dtc DecisionTreeClassifier) IsClassifier() bool {
	return true
}

// IsRegressor returns true as DecisionTreeClassifier is a not regressor
func (dtc DecisionTreeClassifier) IsRegressor() bool {
	return false
}
