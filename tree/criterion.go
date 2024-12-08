package tree

import (
	"GoLab/dataframe/series"
	"math"
)

type criterionFunction func(dfLeftY series.Series, dfRightY series.Series) float64

func gini(dfLeftY series.Series, dfRightY series.Series) float64 {
	// Mathematical formulation of Gini impurity:
	// $1 - \sum_{k}p_{mk}^{2}$
	leftLength := float64(dfLeftY.Len())
	rightLength := float64(dfRightY.Len())
	totalLength := leftLength + rightLength

	// Calculate Gini impurity
	uniqueLeft := dfLeftY.ValueCounts()
	uniqueRight := dfRightY.ValueCounts()

	leftGini := 0.0
	if len(uniqueLeft) != 1 {
		leftGini = 1.0
		for _, u := range uniqueLeft {
			countU := float64(dfLeftY.Count(u))
			p := countU / leftLength
			leftGini -= math.Pow(p, 2)
		}
	}

	rightGini := 0.0
	if len(uniqueRight) != 1 {
		rightGini = 1.0
		for _, u := range uniqueRight {
			countU := float64(dfRightY.Count(u))
			p := countU / rightLength
			rightGini -= math.Pow(p, 2)
		}
	}

	split := (leftLength/totalLength)*leftGini + (rightLength/totalLength)*rightGini

	return split
}

func entropy(dfLeftY series.Series, dfRightY series.Series) float64 {
	// Mathematical formulation of entropy:
	// $-\sum_{k}p_{mk}\log_{2}(p_{mk})$
	leftLength := float64(dfLeftY.Len())
	rightLength := float64(dfRightY.Len())
	totalLength := leftLength + rightLength

	// Calculate entropy
	uniqueLeft := dfLeftY.ValueCounts()
	uniqueRight := dfRightY.ValueCounts()

	leftEntropy := 0.0
	if len(uniqueLeft) != 1 {
		for _, u := range uniqueLeft {
			countU := float64(dfLeftY.Count(u))
			p := countU / leftLength
			leftEntropy -= p * math.Log2(p)
		}
	}

	rightEntropy := 0.0
	if len(uniqueRight) != 1 {
		for _, u := range uniqueRight {
			countU := float64(dfRightY.Count(u))
			p := countU / rightLength
			rightEntropy -= p * math.Log2(p)
		}
	}

	split := (leftLength/totalLength)*leftEntropy + (rightLength/totalLength)*rightEntropy

	return split
}
