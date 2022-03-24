package bst

func NewBST() *BST {
	bst := &BST{root: &node{
		key:   0,
		left:  nil,
		right: nil,
	}}
	return bst
}

//TreeSearch 从x开始查找k
func TreeSearch(x *node, k int) *node {
	if x == nil || k == x.key {
		return x
	}
	if k < x.key {
		return TreeSearch(x.left, k)
	} else {
		return TreeSearch(x.right, k)
	}
}

//TreeMinimum 查找BST中最小值的结点
func TreeMinimum(x *node) *node {
	for x.left != nil {
		x = x.left
	}
	return x
}

//TreeMaximum 查找BST中最大值的结点
func TreeMaximum(x *node) *node {
	for x.right != nil {
		x = x.right
	}
	return x
}

func TreeInsert(t *BST, z *node) {
	y := &node{}
	x := t.root
	for x != nil {
		y = x
		if z.key < x.key {
			x = x.left
		} else {
			x = x.right
		}
	}
	z.parent = y
	if y == nil {
		t.root = z
	} else if z.key < y.key {
		y.left = z
	} else {
		y.right = z
	}
}

func Transplant(t *BST, u, v *node) {
	if u.parent == NilNode() {
		t.root = v
	} else if u == u.parent.left {
		u.parent.left = v
	} else {
		u.parent.right = v
	}
	if v != NilNode() {
		v.parent = u.parent
	}
}

func TreeDelete(t *BST, z *node) {
	if z.left == NilNode() {
		Transplant(t, z, z.right)
	} else if z.right == NilNode() {
		Transplant(t, z, z.left)
	} else {

	}
}
