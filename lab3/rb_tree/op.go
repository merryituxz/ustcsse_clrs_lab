package rb_tree

func NewRBTree() *RBTree {
	rbTree := &RBTree{
		root:    NilNode,
		nilNode: NilNode,
	}
	return rbTree
}

//TreeMinimum 查找RBTree中最小值的结点
func TreeMinimum(t *RBTree, x *Node) *Node {
	for x.Left != t.nilNode {
		x = x.Left
	}
	return x
}

//TreeSearch 从x开始查找k
func TreeSearch(x *Node, k int) *Node {
	if x == nil || k == x.Key {
		return x
	}
	if k < x.Key {
		return TreeSearch(x.Left, k)
	} else {
		return TreeSearch(x.Right, k)
	}
}

//LeftRotate 以结点x为支点进行左旋
func LeftRotate(t *RBTree, x *Node) {
	y := x.Right
	x.Right = y.Left
	if y.Left != t.nilNode {
		y.Left.Parent = x
	}
	y.Parent = x.Parent
	if x.Parent == t.nilNode {
		t.root = y
	} else if x == x.Parent.Left {
		x.Parent.Left = y
	} else {
		x.Parent.Right = y
	}
	y.Left = x
	x.Parent = y
}

//RightRotate 以结点x为支点进行右旋
func RightRotate(t *RBTree, x *Node) {
	y := x.Left
	x.Left = y.Right
	if y.Right != t.nilNode {
		y.Right.Parent = x
	}
	y.Parent = x.Parent
	if x.Parent == t.nilNode {
		t.root = y
	} else if x == x.Parent.Right {
		x.Parent.Right = y
	} else {
		x.Parent.Left = y
	}
	y.Right = x
	x.Parent = y
}

//RBInsert 向RBTree中插入结点z
func RBInsert(t *RBTree, z *Node) {
	y := t.nilNode
	x := t.root
	for x != t.nilNode {
		y = x
		if z.Key < x.Key {
			x = x.Left
		} else {
			x = x.Right
		}
	}
	z.Parent = y
	if y == t.nilNode {
		t.root = z
	} else if z.Key < y.Key {
		y.Left = z
	} else {
		y.Right = z
	}
	z.Left = t.nilNode
	z.Right = t.nilNode
	z.Color = RED
	RBInsertFixup(t, z)
}

//RBInsertFixup 自结点z向上调整红黑树结构
func RBInsertFixup(t *RBTree, z *Node) {
	for z.Parent.Color == RED {
		if z.Parent == z.Parent.Parent.Left {
			y := z.Parent.Parent.Right
			if y.Color == RED {
				z.Parent.Color = BLACK
				y.Color = BLACK
				z.Parent.Parent.Color = RED
				z = z.Parent.Parent
			} else if z == z.Parent.Right {
				z = z.Parent
				LeftRotate(t, z)
			} else {
				z.Parent.Color = BLACK
				z.Parent.Parent.Color = RED
				RightRotate(t, z.Parent.Parent)
			}

		} else if z.Parent == z.Parent.Parent.Right {
			y := z.Parent.Parent.Left
			if y.Color == RED {
				z.Parent.Color = BLACK
				y.Color = BLACK
				z.Parent.Parent.Color = RED
				z = z.Parent.Parent
			} else if z == z.Parent.Left {
				z = z.Parent
				RightRotate(t, z)
			} else {
				z.Parent.Color = BLACK
				z.Parent.Parent.Color = RED
				LeftRotate(t, z.Parent.Parent)
			}
		}
	}
	t.root.Color = BLACK
}

func RBTransplant(t *RBTree, u, v *Node) {
	if u.Parent == t.nilNode {
		t.root = v
	} else if u == u.Parent.Left {
		u.Parent.Left = v
	} else {
		u.Parent.Right = v
	}
	v.Parent = u.Parent
}

func RBDelete(t *RBTree, z *Node) {
	y := z
	yOriginColor := y.Color
	var x *Node
	if z.Left == t.nilNode {
		x = z.Right
		RBTransplant(t, z, z.Right)
	} else if z.Right == t.nilNode {
		x = z.Left
		RBTransplant(t, z, z.Left)
	} else {
		y = TreeMinimum(t, z.Right)
		yOriginColor = y.Color
		x = y.Right
		if y.Parent == z {
			x.Parent = y
		} else {
			RBTransplant(t, y, y.Right)
			y.Right = z.Right
			y.Right.Parent = y
		}
		RBTransplant(t, z, y)
		y.Left = z.Left
		y.Left.Parent = y
		y.Color = z.Color
	}
	if yOriginColor == BLACK {
		RBDeleteFixup(t, z)
	}
}

func RBDeleteFixup(t *RBTree, x *Node) {
	for x != t.root && x.Color == BLACK {
		if x == x.Parent.Left {
			w := x.Parent.Right
			if w.Color == RED {
				w.Color = BLACK
				x.Parent.Color = RED
				LeftRotate(t, x.Parent)
				w = x.Parent.Right
			}
			if w.Left.Color == BLACK && w.Right.Color == BLACK {
				w.Color = RED
				x = x.Parent
			} else if w.Right.Color == BLACK {
				w.Left.Color = BLACK
				w.Color = RED
				RightRotate(t, w)
				w = x.Parent.Right
			} else {
				w.Color = x.Parent.Color
				x.Parent.Color = BLACK
				w.Right.Color = BLACK
				LeftRotate(t, x.Parent)
				x = t.root
			}
		} else if x == x.Parent.Right {
			w := x.Parent.Left
			if w.Color == RED {
				w.Color = BLACK
				x.Parent.Color = RED
				RightRotate(t, x.Parent)
				w = x.Parent.Left
			}
			if w.Left.Color == BLACK && w.Right.Color == BLACK {
				w.Color = RED
				x = x.Parent
			} else if w.Left.Color == BLACK {
				w.Right.Color = BLACK
				w.Color = RED
				LeftRotate(t, w)
				w = x.Parent.Left
			} else {
				w.Color = x.Parent.Color
				x.Parent.Color = BLACK
				w.Left.Color = BLACK
				RightRotate(t, x.Parent)
				x = t.root
			}
		}
	}
	x.Color = BLACK
}

//InorderTraverse 中序遍历, 返回结点序列
func InorderTraverse(t *RBTree) []*Node {
	var inorder func(r *Node)
	var inorderSeq []*Node
	inorder = func(r *Node) {
		if r == nil || r == t.nilNode {
			return
		}
		inorder(r.Left)
		inorderSeq = append(inorderSeq, r)
		inorder(r.Right)
	}
	inorder(t.root)
	return inorderSeq
}
