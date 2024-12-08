# Decision Trees

Decision trees are a popular model for both classification and regression tasks. They are fairly simple to understand
and interpret and can be visualized easily. Their goal is to create a model that will predict the value of a target 
variable by learning simple decision rules inferred from the data features.

---

## Models

The following models are available within the `tree` package:

- [x] [DecisionTreeClassifier](decision_tree_classifier.go)
- [ ] DecisionTreeRegressor
- [ ] ExtraTreeClassifier
- [ ] ExtraTreeRegressor

---

## Usage

Here's an example of how to use the `DecisionTreeClassifier` model:

```go
package main

import (
    "fmt"
    "github.com/chriso345/golab/tree"
    "github.com/chriso345/golab/dataframe"
)

func main() {
    // Define the features and target
    df := dataframe.NewDataFrame(
		series.NewSeries([]float64{0.1245, 0.6589, 0.4487, 0.4578, 0.5978, 0.2534, 0.4356, 0.3215}, series.Float, "Feature1"),
		series.NewSeries([]float64{0.2523, 0.8767, 0.1786, 0.5978, 0.9873, 0.5768, 0.3987, 0.1394}, series.Float, "Feature2"),
		series.NewSeries([]int{1, 0, 1, 1, 0, 1, 0, 1}, series.Int, "Target")
	)
    
    // Extract the feature column
	target := dfX.Drop("Target")
    
    // Create a DecisionTreeClassifier model
    dtc := tree.NewDecisionTreeClassifier()
    
    // Set the hyperparameters
    dtc.SetMaxDepth(3)
    dtc.SetCriterion("entropy")
    
    // Fit the model
    dtc.Fit(df, target)

    // Defining the prediction set
	dfPredict := dataframe.NewDataFrame(
		series.NewSeries([]float64{0.3276, 0.2345, 0.6789, 0.1234, 0.5678, 0.9876, 0.3456, 0.4567}, series.Float, "Feature1"),
		series.NewSeries([]float64{0.47, 0.89, 0.12, 0.34, 0.56, 0.78, 0.23, 0.45}, series.Float, "Feature2"),
	)
	
	// Make predictions
    predictions := dtc.Predict(dfPredict)
    
    fmt.Println(predictions.String())
    // Output: {Target [1 1 0 1 1 0 1 1] int}
}
```

---

## Mathematical formulation

The Decision Tree algorithm can be formulated as follows:

Given a dataset $D$ with $n$ samples and $m$ features, the algorithm recursively splits the data based on the feature that
minimises the loss function. Each data point is represented by $D_i$ (where $i = 1, 2, \dots, n$) and has $m$ features.
```math
\begin{align}
D^{\text{left}} &= \{ D_i \mid D_{i,m} \leq \text{split\_value} \} \\
D^{\text{right}} &= \{ D_i \mid D_{i,m} > \text{split\_value} \}
\end{align}
```
where $D_i$ represents a data point and $D_{i,m}$ is the value of the $m$-th feature for the data point $D_i$

The quality of this split $G(D_{n,m}) can be determined via the loss function as:
```math
G(D) = \frac{|D^{\text{left}}|}{|D|} H(D^{\text{left}}) + \frac{|D^{\text{right}}|}{|D|} H(D^{\text{right}})
```
Where we wish to select the split value such that it minimises the loss function:
```math
\text{argmin}_{m, \text{split\_value}} G(D)
```

### Classification

Where $p_{mk}$ is the represents the probability of class $k$ in the dataset $D$, common measures of impurity are:
- Gini impurity
```math
H(D) = 1 - \sum_{k}p_{mk}^{2}
```
- Entropy (also known as Logarithmic loss)
```math
H(D) = -\sum_{k}p_{mk}\log(p_{mk})
```
Note that when $p_{mk} = 0$, that split is considered pure and the algorithm will not split further.
