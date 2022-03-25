package bst

//NewBST 返回一个空的BST
func NewBST() *BST {
	bst := &BST{root: nil}
	return bst
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

//TreeMinimum 查找BST中最小值的结点
func TreeMinimum(x *Node) *Node {
	for x.Left != nil {
		x = x.Left
	}
	return x
}

//TreeMaximum 查找BST中最大值的结点
func TreeMaximum(x *Node) *Node {
	for x.Right != nil {
		x = x.Right
	}
	return x
}

//TreeInsert 向BST中插入结点z
func TreeInsert(t *BST, z *Node) {
	var y *Node = nil
	x := t.root
	for x != nil {
		y = x
		if z.Key < x.Key {
			x = x.Left
		} else {
			x = x.Right
		}
	}
	z.Parent = y
	if y == nil {
		t.root = z
	} else if z.Key < y.Key {
		y.Left = z
	} else {
		y.Right = z
	}
}

func Transplant(t *BST, u, v *Node) {
	if u.Parent == nil {
		t.root = v
	} else if u == u.Parent.Left {
		u.Parent.Left = v
	} else {
		u.Parent.Right = v
	}
	if v != nil {
		v.Parent = u.Parent
	}
}

//TreeDelete 从BST中删除结点z
func TreeDelete(t *BST, z *Node) {
	if z.Left == nil {
		Transplant(t, z, z.Right)
	} else if z.Right == nil {
		Transplant(t, z, z.Left)
	} else {
		y := TreeMinimum(z.Right)
		if y.Parent != z {
			Transplant(t, y, y.Right)
			y.Right = z.Right
			y.Right.Parent = y
		}
		Transplant(t, z, y)
		y.Left = z.Left
		y.Left.Parent = y
	}
}

//IsValidBST 判断BST是否合法
func IsValidBST(t *BST) ([]int, bool) {
	var inorder func(r *Node)
	var inorderSeq []int
	inorder = func(r *Node) {
		if r == nil {
			return
		}
		inorder(r.Left)
		inorderSeq = append(inorderSeq, r.Key)
		inorder(r.Right)
	}
	inorder(t.root)
	for i := 1; i < len(inorderSeq); i++ {
		if inorderSeq[i] < inorderSeq[i-1] {
			return inorderSeq, false
		}
	}
	return inorderSeq, true
}
