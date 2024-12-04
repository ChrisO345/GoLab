package tree

import (
	"GoLab/dataframe"
	"GoLab/dataframe/series"
	"math"
)

type criterionFunction func(dfLeftY series.Series, dfRightY series.Series) float64

func gini(frame dataframe.DataFrame) float64 {
	panic("gini not implemented")
}

func entropy(dfLeftY series.Series, dfRightY series.Series) float64 {
	leftCount := dfLeftY.Len()
	rightCount := dfRightY.Len()

	leftPositive := dfLeftY.Count(1)
	leftNegative := dfLeftY.Count(0)

	rightPositive := dfRightY.Count(1)
	rightNegative := dfRightY.Count(0)

	// Calculate entropy
	leftEntropy := -((float64(leftPositive)/float64(leftCount))*math.Log2(float64(leftPositive)/float64(leftCount)) + (float64(leftNegative)/float64(leftCount))*math.Log2(float64(leftNegative)/float64(leftCount)))
	if leftPositive == 0 || leftNegative == 0 {
		leftEntropy = 0.0
	}

	rightEntropy := -((float64(rightPositive)/float64(rightCount))*math.Log2(float64(rightPositive)/float64(rightCount)) + (float64(rightNegative)/float64(rightCount))*math.Log2(float64(rightNegative)/float64(rightCount)))
	if rightPositive == 0 || rightNegative == 0 {
		rightEntropy = 0.0
	}

	impurityDrop := - (float64(leftCount)/float64(leftCount+rightCount))*leftEntropy - (float64(rightCount)/float64(leftCount+rightCount))*rightEntropy

	return impurityDrop
}
