package dummy

import (
	"GoLab"
	"GoLab/dataframe"
	"GoLab/dataframe/series"
	"fmt"
)

// DummyClassifier is a struct that represents a dummy classifier
type DummyClassifier struct {
	strategy string

	seriesType series.Type
	result     any

	features []string
	target   string
}

// NewDummyClassifier creates a new DummyClassifier with default values
func NewDummyClassifier() *DummyClassifier {
	return &DummyClassifier{
		strategy: "most_frequent",
		result:   nil,
	}
}

// SetStrategy sets the strategy for the DummyClassifier
func (dc *DummyClassifier) SetStrategy(strategy string) {
	strategyStrings := []string{"most_frequent"} // TODO: implement more strategies

	for _, s := range strategyStrings {
		if s == strategy {
			dc.strategy = strategy
			return
		}
	}

	panic(fmt.Errorf("strategy must be one of %v, but got %v", strategyStrings, strategy))
}

// force implementation of Model interface
var _ GoLab.Model = (*DummyClassifier)(nil)

func (dc *DummyClassifier) Fit(dfX dataframe.DataFrame, dfY series.Series) {
	numSamples, _ := dfX.Shape()
	numOutputs := dfY.Len()

	if numSamples != numOutputs {
		panic(fmt.Errorf("number of samples %v and number of outputs %v must be equal", numSamples, numOutputs))
	}

	if dc.strategy == "most_frequent" {
		dc.result = dfY.Mode()
	}

	dc.features = dfX.Names()
	dc.target = dfY.Name
	dc.seriesType = dfY.Type()
}

func (dc *DummyClassifier) Predict(df dataframe.DataFrame) series.Series {
	if dc.result == nil {
		panic(fmt.Errorf("must fit model before predicting"))
	}

	numSamples, _ := df.Shape()

	for idx, name := range df.Names() {
		if name != dc.features[idx] {
			panic(fmt.Errorf("column %v does not match fit column %v", name, dc.features[idx]))
		}
	}

	switch dc.seriesType {
	case series.Int:
		predictions := make([]int, numSamples)
		for i := 0; i < numSamples; i++ {
			predictions[i] = dc.result.(int)
		}
		return series.New(predictions, dc.seriesType, dc.target)
	case series.Float:
		predictions := make([]float64, numSamples)
		for i := 0; i < numSamples; i++ {
			predictions[i] = dc.result.(float64)
		}
		return series.New(predictions, dc.seriesType, dc.target)
	case series.Boolean:
		predictions := make([]bool, numSamples)
		for i := 0; i < numSamples; i++ {
			predictions[i] = dc.result.(bool)
		}
		return series.New(predictions, dc.seriesType, dc.target)
	case series.String:
		predictions := make([]string, numSamples)
		for i := 0; i < numSamples; i++ {
			predictions[i] = dc.result.(string)
		}
		return series.New(predictions, dc.seriesType, dc.target)
	case series.Runic:
		predictions := make([]rune, numSamples)
		for i := 0; i < numSamples; i++ {
			predictions[i] = dc.result.(rune)
		}
		return series.New(predictions, dc.seriesType, dc.target)
	default:
		panic(fmt.Errorf("series type %v not supported", dc.seriesType))
	}
}

func (dc *DummyClassifier) IsClassifier() bool {
	return true
}

func (dc *DummyClassifier) IsRegressor() bool {
	return false
}
