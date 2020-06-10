package vector_test

import (
	"hsecode.com/stdlib/vector/int"
	"testing"
)

func TestPop(t *testing.T) {
	vec := vector.New(0)
	vec.Push(10)
	vec.Push(11)
	vec.Push(12)
	if vec.Pop() != 12 {
		t.Fatalf("Pop doesn't work: %v", vec)
	}
	if vec.Pop() != 11 {
		t.Fatalf("Pop doesn't work: %v", vec)
	}
	if vec.Pop() != 10 {
		t.Fatalf("Pop doesn't work: %v", vec)
	}
	vec.Push(13)
	vec.Insert(0, 15)
	vec.Insert(1, 14)
	vec.Delete(0)
	if vec.Get(0) != 14 {
		t.Fatalf("Delete doesn't work: %v", vec)
	}
}

func TestInsert(t *testing.T) {
	vec := vector.New(0)
	vec.Push(42)
	vec.Insert(0, 41)
	if (vec.Get(0) != 41) || (vec.Get(1) != 42) {
		t.Fatalf("Push or insert doesn't work: %v", vec)
	}
}

func TestPush(t *testing.T) {
	vec := vector.New(0)
	for _, p := range []int{2, 3, 5, 7, 11} {
		vec.Push(p)
	}
	if vec.Get(2) != 5 {
		t.Fatalf("Push doesn't work: %v", vec)
	}
}
