package rb_tree

type Color = bool

const (
	RED   Color = false
	BLACK Color = true
)

type Node struct {
	Key                 int
	Color               Color
	Left, Right, Parent *Node
}

var NilNode *Node = &Node{
	Key:    0,
	Color:  BLACK,
	Left:   nil,
	Right:  nil,
	Parent: nil,
}

func (t *RBTree) NewNode(key int) *Node {
	n := &Node{
		Key:    key,
		Color:  BLACK,
		Left:   t.nilNode,
		Right:  t.nilNode,
		Parent: t.nilNode,
	}
	return n
}

type RBTree struct {
	root    *Node
	nilNode *Node
}
