package tree

import (
	"GoLab/dataframe"
	"GoLab/dataframe/series"
	"fmt"
	"strings"
	"testing"
)

func TestNewDecisionTreeClassifier(t *testing.T) {
	dtc := NewDecisionTreeClassifier()

	if dtc.criterionString != "gini" {
		t.Errorf("Expected criterion to be gini, got %v", dtc.criterionString)
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

	if dtc.criterionString != "entropy" {
		t.Errorf("Expected criterion to be entropy, got %v", dtc.criterionString)
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected SetCriterion to panic, but it did not")
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
			t.Errorf("Expected SetMaxDepth to panic, but it did not")
		}
	}()

	dtc.SetMaxDepth(0)
}

func TestDecisionTreeClassifier_FitGini(t *testing.T) {
	var expected strings.Builder
	expected.WriteString("Leafs: 6, Depth: 5\n")
	expected.WriteString("Axis: 0, Value: 0.9074\n")
	expected.WriteString("    Axis: 0, Value: 0.4786\n")
	expected.WriteString("        Axis: 1, Value: 0.5709\n")
	expected.WriteString("            Leaf: 0\n")
	expected.WriteString("            Leaf: 1\n")
	expected.WriteString("        Axis: 1, Value: 0.7734\n")
	expected.WriteString("            Leaf: 1\n")
	expected.WriteString("            Axis: 0, Value: 0.5635\n")
	expected.WriteString("                Leaf: 1\n")
	expected.WriteString("                Leaf: 0\n")
	expected.WriteString("    Leaf: 0\n")

	dtc := NewDecisionTreeClassifier()

	dfX := dataframe.New(
		series.New([]float64{0.9074, 0.9529, 0.5635, 0.9567, 0.8162, 0.3279, 0.0179, 0.4246, 0.4770, 0.3394, 0.0788, 0.4853, 0.4786, 0.2427, 0.4001, 0.8530, 0.5159, 0.6385, 0.5231, 0.5486}, series.Float, "Feature1"),
		series.New([]float64{0.5488, 0.6392, 0.7734, 0.9788, 0.9824, 0.3789, 0.3716, 0.1961, 0.3277, 0.0856, 0.5709, 0.7109, 0.9579, 0.8961, 0.9797, 0.4117, 0.3474, 0.1585, 0.4751, 0.0172}, series.Float, "Feature2"),
	)
	dfY := series.New([]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, series.Int, "Target")

	dtc.Fit(dfX, dfY)

	if dtc.tree.String() != expected.String() {
		t.Errorf("Expected:\n%v\nGot:\n%v", expected.String(), dtc.tree.String())
	}
}

func TestDecisionTreeClassifier_FitEntropy(t *testing.T) {
	var expected strings.Builder
	expected.WriteString("Leafs: 4, Depth: 4\n")
	expected.WriteString("Axis: 0, Value: 0.5978\n")
	expected.WriteString("    Axis: 0, Value: 0.4356\n")
	expected.WriteString("        Leaf: 1\n")
	expected.WriteString("        Axis: 0, Value: 0.4487\n")
	expected.WriteString("            Leaf: 0\n")
	expected.WriteString("            Leaf: 1\n")
	expected.WriteString("    Leaf: 0\n")

	dtc := NewDecisionTreeClassifier()
	dtc.SetCriterion("entropy")
	dfX := dataframe.New(
		series.New([]float64{0.1245, 0.6589, 0.4487, 0.4578, 0.5978, 0.2534, 0.4356, 0.3215}, series.Float, "Feature1"),
		series.New([]float64{0.2523, 0.8767, 0.1786, 0.5978, 0.9873, 0.5768, 0.3987, 0.1394}, series.Float, "Feature2"),
	)
	dfY := series.New([]int{1, 0, 1, 1, 0, 1, 0, 1}, series.Int, "Target")

	dtc.Fit(dfX, dfY)

	if dtc.tree.String() != expected.String() {
		t.Errorf("Expected:\n%v\nGot:\n%v", expected.String(), dtc.tree.String())
	}
}

func TestDecisionTreeClassifier_FitEntropy2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("expected no panic, got %v", r)
		}
	}()

	dtc := NewDecisionTreeClassifier()
	dtc.SetCriterion("entropy")

	dfX := dataframe.New(
		series.New([]float64{0.9074, 0.9529, 0.5635, 0.9567, 0.8162, 0.3279, 0.0179, 0.4246, 0.4770, 0.3394, 0.0788, 0.4853, 0.4786, 0.2427, 0.4001, 0.8530, 0.5159, 0.6385, 0.5231, 0.5486}, series.Float, "Feature1"),
		series.New([]float64{0.5488, 0.6392, 0.7734, 0.9788, 0.9824, 0.3789, 0.3716, 0.1961, 0.3277, 0.0856, 0.5709, 0.7109, 0.9579, 0.8961, 0.9797, 0.4117, 0.3474, 0.1585, 0.4751, 0.0172}, series.Float, "Feature2"),
		series.New([]float64{0.9878, 0.1234, 0.5678, 0.9876, 0.3456, 0.4567, 0.2345, 0.6789, 0.1234, 0.5678, 0.9876, 0.3456, 0.4567, 0.2345, 0.6789, 0.1234, 0.5678, 0.9876, 0.3456, 0.4567}, series.Float, "Feature3"),
	)
	dfY := series.New([]int{0, 2, 1, 2, 0, 1, 2, 1, 2, 0, 0, 0, 1, 2, 1, 2, 0, 1, 2, 0}, series.Int, "Target")
	
	dtc.Fit(dfX, dfY)
	
	fmt.Println(dtc.tree.String())
}

func TestDecisionTreeClassifier_Predict(t *testing.T) {
	expected := "{Target [1 1 0 1 1 0 1 1] int}"

	dtc := NewDecisionTreeClassifier()
	dtc.SetCriterion("entropy")
	dfX := dataframe.New(
		series.New([]float64{0.1245, 0.6589, 0.4487, 0.4578, 0.5978, 0.2534, 0.4356, 0.3215}, series.Float, "Feature1"),
		series.New([]float64{0.2523, 0.8767, 0.1786, 0.5978, 0.9873, 0.5768, 0.3987, 0.1394}, series.Float, "Feature2"),
	)
	dfY := series.New([]int{1, 0, 1, 1, 0, 1, 0, 1}, series.Int, "Target")

	dfPredict := dataframe.New(
		series.New([]float64{0.3276, 0.2345, 0.6789, 0.1234, 0.5678, 0.9876, 0.3456, 0.4567}, series.Float, "Feature1"),
		series.New([]float64{0.47, 0.89, 0.12, 0.34, 0.56, 0.78, 0.23, 0.45}, series.Float, "Feature2"),
	)

	dtc.Fit(dfX, dfY)
	predictions := dtc.Predict(dfPredict)

	if predictions.String() != expected {
		t.Errorf("Expected:\n%v\nGot:\n%v", expected, predictions.String())
	}
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
