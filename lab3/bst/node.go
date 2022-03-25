package bst

type Node struct {
	Key                 int
	Left, Right, Parent *Node
}

var NilNode = &Node{}

type BST struct {
	root *Node
}
