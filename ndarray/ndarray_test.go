package ndarray_test

import (
	"hsecode.com/stdlib/ndarray"
	"testing"
)

func TestCompare(t *testing.T) {
	matrix := ndarray.New(3, 3)
	if matrix.Idx(1, 1) != 4 {
		t.Fatal("Wrong answer matrix")
	}

	array3d := ndarray.New(3, 3, 3)
	if array3d.Idx(2, 1, 1) != 22 {
		t.Fatal("Wrong answer 3D array")
	}

	m := ndarray.New(5, 5, 5, 5)
	if m.Idx(0, 0, 0, 0) != 0 {
		t.Fatal("Wrong answer 0")
	}

	n := ndarray.New(5, 5, 5, 5, 5, 5)
	if n.Idx(4, 4, 4, 4, 4, 4) != 5*5*5*5*5*5-1 {
		t.Fatal("Wrong answer N")
	}
}
