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

## Methods



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
