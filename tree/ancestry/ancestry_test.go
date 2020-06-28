package ancestry_test

import (
	"hsecode.com/stdlib/tree"
	"hsecode.com/stdlib/tree/ancestry"
	"testing"
)

func TestAncestry(t *testing.T) {
	T := &tree.Tree{
		Value: 6,
		Left:  &tree.Tree{Value: 2, Left: &tree.Tree{Value: 1}, Right: &tree.Tree{Value: 3}},
		Right: &tree.Tree{Value: 8},
	}

	A := ancestry.New(T)
	t.Log(A.Hash[2])

	t.Log(A.IsDescendant(2, 2)) // false, not a _proper_ descentant
	t.Log(A.IsDescendant(2, 3)) // true
	t.Log(A.IsDescendant(6, 3)) // true
	t.Log(A.IsDescendant(1, 8)) // false
}
