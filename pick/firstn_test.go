package pick_test

import (
	"hsecode.com/stdlib/pick"
	"sort"
	"testing"
)

func TestFirstN(t *testing.T) {
	data := []int{13, 2, 29, 3, 7, 23}
	indices := pick.FirstN(sort.Reverse(sort.IntSlice(data)), 3)
	sort.IntSlice(indices).Sort()
	if (indices[0] != 0) || (indices[1] != 2) || (indices[2] != 5) {
		t.Fatal("Wrong answer")
	}
}
