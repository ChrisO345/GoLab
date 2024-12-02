package tree

import (
	"GoLab/dataframe"
	"GoLab/dataframe/series"
	"math"
)

type criterionFunction func(dataframe.DataFrame) float64

func gini(frame dataframe.DataFrame) float64 {
	panic("gini not implemented")
}

func entropy(dfLeftX dataframe.DataFrame, dfLeftY series.Series, dfRightX dataframe.DataFrame, dfRightY series.Series) float64 {
	leftCount := dfLeftY.Len()
	rightCount := dfRightY.Len()

	leftPositive := 0
	leftNegative := 0
	for j := 0; j < leftCount; j++ {
		if dfLeftY.Val(j).(int) == 1 {
			leftPositive++
		} else {
			leftNegative++
		}
	}

	rightPositive := 0
	rightNegative := 0
	for j := 0; j < rightCount; j++ {
		if dfRightY.Val(j).(int) == 1 {
			rightPositive++
		} else {
			rightNegative++
		}
	}

	// Calculate entropy
	leftEntropy := 0.0
	if leftPositive != 0 && leftNegative != 0 {
		leftEntropy = -((float64(leftPositive)/float64(leftCount))*math.Log2(float64(leftPositive)/float64(leftCount)) + (float64(leftNegative)/float64(leftCount))*math.Log2(float64(leftNegative)/float64(leftCount)))
	}

	rightEntropy := 0.0
	if rightPositive != 0 && rightNegative != 0 {
		rightEntropy = -((float64(rightPositive)/float64(rightCount))*math.Log2(float64(rightPositive)/float64(rightCount)) + (float64(rightNegative)/float64(rightCount))*math.Log2(float64(rightNegative)/float64(rightCount)))
	}

	impurityDrop := (float64(leftCount)/float64(leftCount+rightCount))*leftEntropy + (float64(rightCount)/float64(leftCount+rightCount))*rightEntropy

	return impurityDrop
}
