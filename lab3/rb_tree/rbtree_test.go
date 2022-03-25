package rb_tree

import (
	"fmt"
	"testing"
)

var rbTree *RBTree = NewRBTree()

func TestInitRBTree(t *testing.T) {
	fmt.Println(*rbTree)
}

func TestRBInsert(t *testing.T) {
	RBInsert(rbTree, rbTree.NewNode(1))
	RBInsert(rbTree, rbTree.NewNode(5))
	RBInsert(rbTree, rbTree.NewNode(6))
	RBInsert(rbTree, rbTree.NewNode(7))
	RBInsert(rbTree, rbTree.NewNode(8))
	RBInsert(rbTree, rbTree.NewNode(9))
	RBInsert(rbTree, rbTree.NewNode(10))
	RBInsert(rbTree, rbTree.NewNode(11))
	RBInsert(rbTree, rbTree.NewNode(12))
	RBInsert(rbTree, rbTree.NewNode(13))
	RBInsert(rbTree, rbTree.NewNode(14))

	inorderSeq := InorderTraverse(rbTree)

	for _, node := range inorderSeq {
		if node.Color == RED {
			fmt.Printf("key: %d, color: red\n", node.Key)
		} else {
			fmt.Printf("key: %d, color: black\n", node.Key)
		}
	}
}

func TestRBDelete(t *testing.T) {
	RBInsert(rbTree, rbTree.NewNode(1))
	RBInsert(rbTree, rbTree.NewNode(5))
	RBInsert(rbTree, rbTree.NewNode(6))
	RBInsert(rbTree, rbTree.NewNode(7))
	RBInsert(rbTree, rbTree.NewNode(8))
	RBInsert(rbTree, rbTree.NewNode(9))
	RBInsert(rbTree, rbTree.NewNode(10))
	RBInsert(rbTree, rbTree.NewNode(11))
	RBInsert(rbTree, rbTree.NewNode(12))
	RBInsert(rbTree, rbTree.NewNode(13))
	RBInsert(rbTree, rbTree.NewNode(14))

	RBDelete(rbTree, TreeSearch(rbTree.root, 14))
	RBDelete(rbTree, TreeSearch(rbTree.root, 9))
	//RBDelete(rbTree, TreeSearch(rbTree.root, 5))

	inorderSeq := InorderTraverse(rbTree)

	for _, node := range inorderSeq {
		if node.Color == RED {
			fmt.Printf("key: %d, color: red\n", node.Key)
		} else {
			fmt.Printf("key: %d, color: black\n", node.Key)
		}
	}
}
