package dummy

import (
	"GoLab/dataframe"
	"GoLab/dataframe/series"
	"testing"
)

func TestNewDummyRegressor(t *testing.T) {
	dc := NewDummyRegressor()

	if dc.strategy != "mean" {
		t.Errorf("Expected strategy to be mean, got %v", dc.strategy)
	}

	if dc.result != nil {
		t.Errorf("Expected result to be nil, got %v", dc.result)
	}
}

func TestDummyRegressor_SetStrategy(t *testing.T) {
	dc := NewDummyRegressor()

	dc.SetStrategy("median")

	if dc.strategy != "median" {
		t.Errorf("Expected strategy to be median, got %v", dc.strategy)
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected SetStrategy to panic, but it did not")
		}
	}()

	dc.SetStrategy("not a valid criterion")
}

func TestDummyRegressor_SetQuantile(t *testing.T) {
	dc := NewDummyRegressor()

	dc.SetQuantile(0.5)

	if dc.quantile != 0.5 {
		t.Errorf("Expected quantile to be 0.5, got %v", dc.quantile)
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected SetQuantile to panic, but it did not")
		}
	}()

	dc.SetQuantile(-1)
}

func TestDummyRegressor_Fit(t *testing.T) {
	dc := NewDummyRegressor()

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

func TestDummyRegressor_Predict(t *testing.T) {
	dc := NewDummyRegressor()

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
}

func TestDummyRegressor_IsClassifier(t *testing.T) {
	dc := NewDummyRegressor()

	if dc.IsClassifier() {
		t.Errorf("Expected IsClassifier to return false, got true")
	}
}

func TestDummyRegressor_IsRegressor(t *testing.T) {
	dc := NewDummyRegressor()

	if !dc.IsRegressor() {
		t.Errorf("Expected IsRegressor to return true, got false")
	}
}
