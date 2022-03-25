package bst

import (
	"fmt"
	"testing"
)

var bst *BST = NewBST()

func TestInitBST(t *testing.T) {
	fmt.Println(*bst.root)
}

func TestTreeInsert(t *testing.T) {
	TreeInsert(bst, &Node{Key: 8})
	TreeInsert(bst, &Node{Key: 3})
	TreeInsert(bst, &Node{Key: 1})
	TreeInsert(bst, &Node{Key: 6})
	TreeInsert(bst, &Node{Key: 4})
	TreeInsert(bst, &Node{Key: 7})
	TreeInsert(bst, &Node{Key: 10})
	TreeInsert(bst, &Node{Key: 14})
	TreeInsert(bst, &Node{Key: 13})
	inorderSeq, valid := IsValidBST(bst)
	fmt.Println(inorderSeq, "valid:", valid)
}

func TestTreeDelete(t *testing.T) {
	TestTreeInsert(t)

	n := TreeSearch(bst.root, 3)

	TreeDelete(bst, n)

	inorderSeq, valid := IsValidBST(bst)
	fmt.Println(inorderSeq, "valid:", valid)
}
