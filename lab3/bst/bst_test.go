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
	TreeInsert(bst, &node{key: 2})
	TreeInsert(bst, &node{key: 1})
	TreeInsert(bst, &node{key: 3})
	fmt.Println(*bst.root)
}
