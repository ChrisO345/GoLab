package preprocessing

import (
	"GoLab/dataframe"
	"GoLab/dataframe/series"
	"testing"
)

func TestOneHotEncoder_Fit(t *testing.T) {
	ohe := NewOneHotEncoder()

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Expected no panic, but got %v", r)
		}
	}()

	dfX := dataframe.New(
		series.New([]int{1, 2, 3}, series.Int, "Integers"),
		series.New([]float64{4.4, 5.5, 6.6}, series.Float, "Floats"),
		series.New([]string{"a", "b", "c"}, series.String, "Strings"),
	)

	ohe.Fit(dfX)

	if ohe.nUnique != 9 {
		t.Errorf("Expected nUnique to be 9, got %v", ohe.nUnique)
	}

	featureNames := []string{"Integers", "Floats", "Strings"}
	for i, name := range ohe.featureNames {
		if name != featureNames[i] {
			t.Errorf("Expected feature name to be %v, got %v", featureNames[i], name)
		}
	}

	for _, name := range ohe.featureNames {
		if _, ok := ohe.encoder[name]; !ok {
			t.Errorf("Expected encoder to have key %v", name)
		}
	}
}

func TestOneHotEncoder_Transform(t *testing.T) {
	ohe := NewOneHotEncoder()

	df := dataframe.New(
		series.New([]int{1, 2, 3}, series.Int, "Integers"),
		series.New([]float64{4.4, 5.5, 6.6}, series.Float, "Floats"),
		series.New([]string{"a", "b", "c"}, series.String, "Strings"),
	)

	ohe.Fit(df)
	result := ohe.Transform(df)

	numSamples, _ := result.Shape()
	if numSamples != 3 {
		t.Errorf("Expected numSamples to be 3, got %v", numSamples)
	}

	for _, name := range result.Names() {
		if name != "Integers_1" && name != "Integers_2" && name != "Integers_3" &&
			name != "Floats_4.4" && name != "Floats_5.5" && name != "Floats_6.6" &&
			name != "Strings_a" && name != "Strings_b" && name != "Strings_c" {
			t.Errorf("Unexpected column name %v", name)
		}
	}
}

func TestOneHotEncoder_InverseTransform(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Not Implemented")
		}
	}()
}

func TestOneHotEncoder_GetFeatureNames(t *testing.T) {
	ohe := NewOneHotEncoder()

	df := dataframe.New(
		series.New([]int{1, 2, 3}, series.Int, "Integers"),
		series.New([]float64{4.4, 5.5, 6.6}, series.Float, "Floats"),
		series.New([]string{"a", "b", "c"}, series.String, "Strings"),
	)

	ohe.Fit(df)
	featureNames := ohe.GetFeatureNames()

	if len(featureNames) != 3 {
		t.Errorf("Expected featureNames to have length 3, got %v", len(featureNames))
	}

	for i, name := range featureNames {
		if name != df.Names()[i] {
			t.Errorf("Expected feature name to be %v, got %v", df.Names()[i], name)
		}
	}
}
