package bitset_test

import (
	"hsecode.com/stdlib/bitset"
	"testing"
)

func TestBitset_All(t *testing.T) {
	bs := bitset.New(2)
	ans := bs.All()
	if ans != false {
		t.Fatalf("Wrong answer: %v", ans)
	}
	_ = bs.Set(0, true)
	_ = bs.Set(1, true)
	ans = bs.All()
	if ans != true {
		t.Fatalf("Wrong answer: %v", ans)
	}
}

func TestBitset_Any(t *testing.T) {
	bs := bitset.New(3)
	ans := bs.Any()
	if ans != false {
		t.Fatalf("Wrong answer: %v", ans)
	}
	_ = bs.Set(1, true)
	ans = bs.Any()
	if ans != true {
		t.Fatalf("Wrong answer: %v", ans)
	}
}

func TestBitset_Count(t *testing.T) {
	bs := bitset.New(3)

	ans := bs.Count()
	if ans != 0 {
		t.Fatalf("Wrong answer: %v", ans)
	}
	_ = bs.Set(0, true)
	ans = bs.Count()
	if ans != 1 {
		t.Fatalf("Wrong answer: %v", ans)
	}
	_ = bs.Set(1, true)
	ans = bs.Count()
	if ans != 2 {
		t.Fatalf("Wrong answer: %v", ans)
	}
}

func TestBitset_Flip(t *testing.T) {
	bs := bitset.New(3)
	_ = bs.Set(0, true)
	_ = bs.Set(1, false)

	bs.Flip()
	v1, _ := bs.Test(0) // false
	v2, _ := bs.Test(1) // true
	if v1 != false {
		t.Fatalf("Wrong answer v1: %v", v1)
	}
	if v2 != true {
		t.Fatalf("Wrong answer v2: %v", v2)
	}
}

func TestBitset_Reset(t *testing.T) {
	bs := bitset.New(3)
	_ = bs.Set(1, true)
	bs.Reset()
	value, _ := bs.Test(1)
	if value != false {
		t.Fatalf("Wrong answer: %v", value)
	}
}

func TestBitset_Set(t *testing.T) {
	bs := bitset.New(3)
	_ = bs.Set(1, true)

	value, _ := bs.Test(1)
	if value != true {
		t.Fatalf("Wrong answer: %v", value)
	}
	err := bs.Set(3, false)
	if err == nil {
		t.Fatalf("Wrong error: %v", err)
	}
}

func TestBitset_Test(t *testing.T) {
	bs := bitset.New(3)
	_ = bs.Set(1, true)

	value, _ := bs.Test(0)
	if value != false {
		t.Fatalf("Wrong answer: %v", value)
	}
	value, _ = bs.Test(1)
	if value != true {
		t.Fatalf("Wrong answer: %v", value)
	}
	_, err := bs.Test(3)
	if err == nil {
		t.Fatalf("Wrong error: %v", err)
	}
}
