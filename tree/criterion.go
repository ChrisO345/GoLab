package tree

import (
	"GoLab/dataframe/series"
	"math"
)

type criterionFunction func(dfLeftY series.Series, dfRightY series.Series) float64

func gini(dfLeftY series.Series, dfRightY series.Series) float64 {
	panic("gini not implemented")
}

func entropy(dfLeftY series.Series, dfRightY series.Series) float64 {
	leftLength := float64(dfLeftY.Len())
	rightLength := float64(dfRightY.Len())
	totalLength := leftLength + rightLength

	// Calculate entropy
	uniqueLeft := dfLeftY.Uniques()
	uniqueRight := dfRightY.Uniques()

	leftEntropy := 0.0
	if len(uniqueLeft) != 1 {
		for _, u := range uniqueLeft {
			countU := float64(dfLeftY.Count(u))
			leftEntropy -= (countU/leftLength)*math.Log2(countU/leftLength)
		}
	}

	rightEntropy := 0.0
	if len(uniqueRight) != 1 {
		for _, u := range uniqueRight {
			countU := float64(dfRightY.Count(u))
			rightEntropy -= (countU/rightLength)*math.Log2(countU/rightLength)
		}
	}

	impurityDrop := (leftLength/totalLength)*leftEntropy + (rightLength/totalLength)*rightEntropy

	return impurityDrop
}
