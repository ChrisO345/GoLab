package preprocessing

import (
	"fmt"
	"github.com/chriso345/golab/dataframe"
	"github.com/chriso345/golab/dataframe/series"
)

type Encoder interface {
	Fit(dfX dataframe.DataFrame)
	Transform(df dataframe.DataFrame) dataframe.DataFrame
	FitTransform(dfX dataframe.DataFrame) dataframe.DataFrame

	InverseTransform(df dataframe.DataFrame) dataframe.DataFrame

	GetFeatureNames() []string
}

// force implementation of Encoder interface
var _ Encoder = (*OneHotEncoder)(nil)

// OneHotEncoder is a struct that represents a one-hot encoder
type OneHotEncoder struct {
	featureNames []string
	encoder      map[string][]any
	nUnique      int
}

// NewOneHotEncoder creates a new OneHotEncoder with default values
func NewOneHotEncoder() *OneHotEncoder {
	return &OneHotEncoder{
		featureNames: nil,
		encoder:      nil,
		nUnique:      0,
	}
}

func (ohe *OneHotEncoder) Fit(dfX dataframe.DataFrame) {
	ohe.featureNames = dfX.Names()

	ohe.encoder = make(map[string][]any, len(ohe.featureNames))

	ohe.nUnique = 0
	for _, name := range ohe.featureNames {
		uniqueValues := dfX.Column(name).ValueCounts()
		ohe.nUnique += len(uniqueValues)

		ohe.encoder[name] = make([]any, len(uniqueValues))

		cnt := 0
		for key, _ := range uniqueValues {
			ohe.encoder[name][cnt] = key
			cnt++
		}
	}
}

func (ohe OneHotEncoder) Transform(df dataframe.DataFrame) dataframe.DataFrame {
	if ohe.encoder == nil {
		panic(fmt.Errorf("OneHotEncoder is not fitted"))
	}

	for _, name := range ohe.featureNames {
		found := false
		for _, col := range df.Columns() {
			if col.Name == name {
				found = true
				break
			}
		}
		if !found {
			panic(fmt.Errorf("column %v not found in DataFrame", name))
		}
	}

	numSamples, _ := df.Shape()
	result := make([][]int, ohe.nUnique)

	for i := 0; i < numSamples; i++ {
		cnt := 0
		for _, name := range ohe.featureNames {
			col := df.Column(name)
			for _, val := range ohe.encoder[name] {
				if val == col.Val(i) {
					result[cnt] = append(result[cnt], 1)
				} else {
					result[cnt] = append(result[cnt], 0)
				}
				cnt++
			}
		}
	}

	var names []string
	for _, name := range ohe.featureNames {
		for _, val := range ohe.encoder[name] {
			names = append(names, fmt.Sprintf("%v_%v", name, val))
		}
	}

	se := make([]series.Series, ohe.nUnique)
	for i := 0; i < ohe.nUnique; i++ {
		se[i] = series.New(result[i], series.Int, fmt.Sprintf("%v", names[i]))
	}

	return dataframe.New(se...)
}

func (ohe *OneHotEncoder) FitTransform(dfX dataframe.DataFrame) dataframe.DataFrame {
	ohe.Fit(dfX)
	return ohe.Transform(dfX)
}

func (ohe OneHotEncoder) InverseTransform(df dataframe.DataFrame) dataframe.DataFrame {
	//TODO implement me
	panic("implement me")
}

func (ohe OneHotEncoder) GetFeatureNames() []string {
	return ohe.featureNames
}
