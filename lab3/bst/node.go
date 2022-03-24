package bst

type node struct {
	key                 int
	left, right, parent *node
}

type BST struct {
	root *node
}

func NilNode() *node {
	return &node{}
}
