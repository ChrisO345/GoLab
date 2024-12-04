package tree

import (
	"fmt"
	"math"
	"strings"
)

type DecisionTree struct {
	Left  *DecisionTree
	Right *DecisionTree

	Leaf bool
	Axis int
	Value float64
	Label int
}

func (dt *DecisionTree) getStringer(depth int) (string, int, int) {
	if dt == nil {
		return "", 0, 0
	}

	var s strings.Builder

	for i := 0; i < depth; i++ {
		s.WriteString("    ")
	}

	if dt.Leaf {
		s.WriteString("Leaf: ")
		s.WriteString(fmt.Sprintf("%v", dt.Label))
		s.WriteString("\n")
		return s.String(), 1, 1
	}
	s.WriteString("Axis: ")
	s.WriteString(fmt.Sprintf("%v", dt.Axis))
	s.WriteString(", Value: ")
	s.WriteString(fmt.Sprintf("%v", dt.Value))
	s.WriteString("\n")
	leftString, leftLeaf, leftDepth := dt.Left.getStringer(depth + 1)
	s.WriteString(leftString)
	rightString, rightLeaf, rightDepth := dt.Right.getStringer(depth + 1)
	s.WriteString(rightString)

	maxDepth := 1 + math.Max(float64(leftDepth), float64(rightDepth))

	return s.String(), leftLeaf + rightLeaf, int(maxDepth)
}

func (dt *DecisionTree) String() string {
	if dt == nil {
		return "Empty Decision Tree"
	}

	s, leafs, depth := dt.getStringer(0)
	return fmt.Sprintf("Leafs: %v, Depth: %v\n%v", leafs, depth, s)
}
