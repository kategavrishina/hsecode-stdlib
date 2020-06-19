package tree_test

import (
	"fmt"
	"hsecode.com/stdlib/tree"
	"testing"
)

func TestNewBST(t *testing.T) {
	elements := []int{11, 2, 5, 7, 3, 3, 5, 2}
	T := tree.NewBST(elements)
	T.InOrder(func(node *tree.Tree) { fmt.Println(node.Value) })
}
