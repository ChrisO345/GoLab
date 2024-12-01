package tree

import (
	"GoLab/dataframe"
	"GoLab/dataframe/series"
	"fmt"
	"testing"
)

func TestNewDecisionTreeClassifier(t *testing.T) {
	dtc := NewDecisionTreeClassifier()

	if dtc.criterion != "gini" {
		t.Errorf("Expected criterion to be gini, got %v", dtc.criterion)
	}

	if dtc.maxDepth != -1 {
		t.Errorf("Expected maxDepth to be -1, got %v", dtc.maxDepth)
	}

	if dtc.tree != nil {
		t.Errorf("Expected tree to be nil, got %v", dtc.tree)
	}
}

func TestDecisionTreeClassifier_SetCriterion(t *testing.T) {
	dtc := NewDecisionTreeClassifier()

	dtc.SetCriterion("entropy")

	if dtc.criterion != "entropy" {
		t.Errorf("Expected criterion to be entropy, got %v", dtc.criterion)
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	dtc.SetCriterion("not a valid criterion")
}

func TestDecisionTreeClassifier_SetMaxDepth(t *testing.T) {
	dtc := NewDecisionTreeClassifier()

	dtc.SetMaxDepth(5)

	if dtc.maxDepth != 5 {
		t.Errorf("Expected maxDepth to be 5, got %v", dtc.maxDepth)
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	dtc.SetMaxDepth(0)
}

func TestDecisionTreeClassifier_Fit(t *testing.T) {
	dtc := NewDecisionTreeClassifier()

	defer func() {
		if r := recover(); r != nil {
			t.Errorf(fmt.Sprintf("%v", r))
		}
	}()

	dfX := dataframe.NewDataFrame(
		series.NewSeries([]int{1, 3, 2}, series.Int, "Feature1"),
		series.NewSeries([]int{9, 8, 7}, series.Int, "Feature2"),
	)
	dfY := series.NewSeries([]int{1, 0, 1}, series.Int, "Target")
	dtc.Fit(dfX, dfY)
}

func TestDecisionTreeClassifier_Predict(t *testing.T) {
	dtc := NewDecisionTreeClassifier()

	defer func() {
		if r := recover(); r != nil {
			t.Errorf(fmt.Sprintf("%v", r))
		}
	}()

	dtc.Predict()
}

func TestDecisionTreeClassifier_IsClassifier(t *testing.T) {
	dtc := NewDecisionTreeClassifier()

	if !dtc.IsClassifier() {
		t.Errorf("Expected IsClassifier to return true, got false")
	}
}

func TestDecisionTreeClassifier_IsRegressor(t *testing.T) {
	dtc := NewDecisionTreeClassifier()

	if dtc.IsRegressor() {
		t.Errorf("Expected IsRegressor to return false, got true")
	}
}
