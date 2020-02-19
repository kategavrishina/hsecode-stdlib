package cmp_test

import (
	cmp "hsecode.com/stdlib/cmp/int"
	"testing"
)

func TestCompare(t *testing.T) {
	if cmp.Min(4, 3, 2, 0, 10) != 0 {
		t.Fatal("unexpected Min() value")
	}
	a := []int{4, 3, 2, 1}
	r := cmp.Max(a...)
	if r != 4 {
		t.Fatalf("Max(%v...) = %v, expected 4", a, r)
	}
}
