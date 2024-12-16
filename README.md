# golab

*Machine Learning in Go*

**golab** is a machine learning library for Go that provides tools for building and training machine learning models. It includes a variety of models, such as decision trees, linear regression, and k-means clustering, as well as tools for data preprocessing and evaluation.

___

## Installation

To install **golab**, you need to have Go installed on your machine. Then, you can run the following command:

```bash
go get github.com/chriso345/golab
```

___

## Usage

Here is an example using the `DecisionTreeClassifier` model to make predictions on a dataset:

```go
package main

import (
	"fmt"
	"github.com/chriso345/golab/dataframe"
	"github.com/chriso345/golab/dataframe/series"
	"github.com/chriso345/golab/tree"
)

func main() {
	// Define the training dataframe
	dfX := dataframe.New(
		series.New([]float64{0.1245, 0.6589, 0.4487, 0.4578, 0.5978, 0.2534, 0.4356, 0.3215}, series.Float, "Feature1"),
		series.New([]float64{0.2523, 0.8767, 0.1786, 0.5978, 0.9873, 0.5768, 0.3987, 0.1394}, series.Float, "Feature2"),
		series.New([]int{1, 0, 1, 1, 0, 1, 0, 1}, series.Int, "Target"),
	)

	// Drop the target column from the dataframe
	target := dfX.Drop("Target")

	// Create the DecisionTreeClassifier
	dtc := tree.NewDecisionTreeClassifier()

	// Set the criterion to entropy
	dtc.SetCriterion("entropy")

	// Fit the model to the training data
	dtc.Fit(dfX, target)

	// Define the prediction dataframe
	dfPredict := dataframe.New(
		series.New([]float64{0.3276, 0.2345, 0.6789, 0.1234, 0.5678, 0.9876, 0.3456, 0.4567}, series.Float, "Feature1"),
		series.New([]float64{0.47, 0.89, 0.12, 0.34, 0.56, 0.78, 0.23, 0.45}, series.Float, "Feature2"),
	)

	// Make predictions
	predictions := dtc.Predict(dfPredict)

	fmt.Println(predictions)
	// {Target [1 1 0 1 1 0 1 1] int}
}

```

___

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.