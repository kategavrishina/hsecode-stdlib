package pick_test

import (
	"hsecode.com/stdlib/pick"
	"sort"
	"testing"
)

func TestNthElement(t *testing.T) {
	data := []int{13, 2, 29, 3, 7, 23, 31}
	m := len(data) / 2 // median index if the array was sorted
	pick.NthElement(sort.IntSlice(data), m)
	if data[m] != 13 {
		t.Fatal("Wrong answer")
	}
}
