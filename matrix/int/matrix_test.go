package matrix_test

import (
	matrix "hsecode.com/stdlib/matrix/int"
	"testing"
)

func Test(t *testing.T) {
	m := matrix.New(2, 2)
	m.Set(0, 1, 56)
	if m.Get(0, 1) != 56 {
		t.Fatal("error")
	}
}
