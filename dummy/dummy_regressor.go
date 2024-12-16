package dummy

import (
	"GoLab"
	"GoLab/dataframe"
	"GoLab/dataframe/series"
	"fmt"
)

// DummyRegressor is a struct that represents a dummy regressor
type DummyRegressor struct {
	strategy string
	quantile float64

	seriesType series.Type
	result     any

	features []string
	target   string
}

// NewDummyRegressor creates a new DummyRegressor with default values
func NewDummyRegressor() *DummyRegressor {
	return &DummyRegressor{
		strategy: "mean",
		quantile: 0.5,
		result:   nil,
	}
}

// SetStrategy sets the strategy for the DummyRegressor
func (dr *DummyRegressor) SetStrategy(strategy string) {
	strategyStrings := []string{"mean", "median", "quantile"} // TODO: implement more strategies

	for _, s := range strategyStrings {
		if s == strategy {
			if s == "median" {
				dr.quantile = 0.5
			}
			dr.strategy = strategy
			return
		}
	}

	panic(fmt.Errorf("strategy must be one of %v, but got %v", strategyStrings, strategy))
}

// SetQuantile sets the quantile for the DummyRegressor
func (dr *DummyRegressor) SetQuantile(quantile float64) {
	if quantile < 0 || quantile > 1 {
		panic(fmt.Errorf("quantile must be between 0 and 1, but got %v", quantile))
	}

	dr.quantile = quantile
}

// force implementation of Model interface
var _ GoLab.Model = (*DummyRegressor)(nil)

func (dr *DummyRegressor) Fit(dfX dataframe.DataFrame, dfY series.Series) {
	numSamples, _ := dfX.Shape()
	numOutputs := dfY.Len()

	if numSamples != numOutputs {
		panic(fmt.Errorf("number of samples %v and number of outputs %v must be equal", numSamples, numOutputs))
	}

	if dr.strategy == "mean" {
		dr.result = dfY.Mean()
	} else if dr.strategy == "median" {
		dr.result = dfY.Median()
	} else if dr.strategy == "quantile" {
		dr.result = dfY.Quantile(dr.quantile)
	}

	dr.features = dfX.Names()
	dr.target = dfY.Name
	dr.seriesType = dfY.Type()
}

func (dr *DummyRegressor) Predict(dfX dataframe.DataFrame) series.Series {
	if dr.result == nil {
		panic(fmt.Errorf("DummyRegressor is not fitted"))
	}

	numSamples, _ := dfX.Shape()

	predictions := make([]float64, numSamples)
	for i := 0; i < numSamples; i++ {
		predictions[i] = dr.result.(float64)
	}

	return series.New(predictions, dr.seriesType, dr.target)
}

func (dr *DummyRegressor) IsClassifier() bool {
	return false
}

func (dr *DummyRegressor) IsRegressor() bool {
	return true
}