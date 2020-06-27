package tree

import (
	"fmt"
	"testing"
)

func TestEncode(t *testing.T) {
	T := &Tree{
		Value: 2,
		Left:  &Tree{Value: 1},
		Right: &Tree{Value: 3, Right: &Tree{Value: 10}},
	}
	t.Log(T.Encode())
}

func TestDecode(t *testing.T) {
	T, err := Decode([]string{"2", "1", "3", "nil", "nil", "nil", "10"})

	// Resulting tree:
	//    2
	//   / \
	//  1    3
	//        \
	//         10
	t.Log(err)
	T.InOrder(func(node *Tree) { fmt.Println(node.Value) })
}
