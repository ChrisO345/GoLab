package dummy

import (
	"GoLab/dataframe"
	"GoLab/dataframe/series"
	"testing"
)

func TestNewDummyClassifier(t *testing.T) {
	dc := NewDummyClassifier()

	if dc.strategy != "most_frequent" {
		t.Errorf("Expected strategy to be most_frequent, got %v", dc.strategy)
	}

	if dc.result != nil {
		t.Errorf("Expected result to be nil, got %v", dc.result)
	}
}

func TestDummyClassifier_SetStrategy(t *testing.T) {
	dc := NewDummyClassifier()

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected SetStrategy to panic, but it did not")
		}
	}()

	dc.SetStrategy("not a valid criterion")
}

func TestDummyClassifier_Fit(t *testing.T) {
	dc := NewDummyClassifier()

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Expected no panic, but got %v", r)
		}
	}()

	dfX := dataframe.New(
		series.New([]int{1, 2, 3}, series.Int, "Integers"),
		series.New([]float64{4.4, 5.5, 6.6}, series.Float, "Floats"),
	)
	dfY := series.New([]int{2, 2, 3}, series.Int, "Integers")

	dc.Fit(dfX, dfY)
}

func TestDummyClassifier_Predict(t *testing.T) {
	dc := NewDummyClassifier()

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Expected no panic, but got %v", r)
		}
	}()

	dfX := dataframe.New(
		series.New([]int{1, 2, 3}, series.Int, "Integers"),
		series.New([]float64{4.4, 5.5, 6.6}, series.Float, "Floats"),
	)
	dfY := series.New([]int{2, 2, 3}, series.Int, "Integers")

	dc.Fit(dfX, dfY)

	predictions := dc.Predict(dfX)

	if !predictions.Homogeneous() {
		t.Errorf("Expected predictions to be homogeneous, got %v", predictions.Homogeneous())
	}

	if predictions.Len() != 3 {
		t.Errorf("Expected predictions to have length 3, got %v", predictions.Len())
	}

	for i := 0; i < 3; i++ {
		if predictions.Val(i) != 2 {
			t.Errorf("Expected prediction to be 2, got %v", predictions.Val(i))
		}
	}
}

func TestDummyClassifier_IsClassifier(t *testing.T) {
	dc := NewDummyClassifier()

	if !dc.IsClassifier() {
		t.Errorf("Expected IsClassifier to return true, got false")
	}
}

func TestDummyClassifier_IsRegressor(t *testing.T) {
	dc := NewDummyClassifier()

	if dc.IsRegressor() {
		t.Errorf("Expected IsRegressor to return false, got true")
	}
}
