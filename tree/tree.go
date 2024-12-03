package tree

type DecisionTree struct {
	Left  *DecisionTree
	Right *DecisionTree

	Leaf bool
	Axis int
	Position int
	Label int
}

//func (dt *DecisionTree) getStringer(depth int) string {
//	if dt == nil {
//		return ""
//	}
//	return ""
//}
//
//func (dt *DecisionTree) String() string {
//	if dt == nil {
//		return "Empty Decision Tree"
//	}
//	return dt.getStringer(0)
//}
