package lru_test

import (
	"hsecode.com/stdlib/lru"
	"testing"
)

func TestEasyGetLRU(t *testing.T) {
	c := lru.New(3) // cache of size 3
	c.Put(1, 100)
	c.Put(2, 200)
	c.Put(3, 300)
	ansInt, ansBool := c.Get(1)
	if ansInt != 100 {
		t.Fatalf("Get existing key doesn't work: value %d", ansInt)
	}
	if ansBool != true {
		t.Fatalf("Get existing key doesn't work: bool %v", ansBool)
	}
	ansInt1, ansBool1 := c.Get(5)
	if ansInt1 != 0 {
		t.Fatalf("Get not existing key doesn't work: value %d", ansInt1)
	}
	if ansBool1 != false {
		t.Fatalf("Get not existing key doesn't work: bool %v", ansBool1)
	}
}

func TestGetPutLRU(t *testing.T) {
	c := lru.New(2) // cache of size 2

	c.Put(1, 100)
	c.Put(2, 200)
	c.Put(3, 300)
	ansInt, ansBool := c.Get(1)
	if ansInt != 0 {
		t.Fatalf("Get removed key doesn't work: value %d", ansInt)
	}
	if ansBool != false {
		t.Fatalf("Get removed key doesn't work: bool %v", ansBool)
	}

	c.Get(2)      // "use" the key
	c.Put(4, 400) // remove `3` because `2` was used recently.
	c.Put(2, 200) // "use" `2`
	c.Put(5, 500) // remove `4`, `2` is still there
	ansInt1, ansBool1 := c.Get(2)
	if ansInt1 != 200 {
		t.Fatalf("Get recently used key doesn't work: value %d", ansInt1)
	}
	if ansBool1 != true {
		t.Fatalf("Get recently used key doesn't work: bool %v", ansBool1)
	}

	ansInt2, ansBool2 := c.Get(4)
	if ansInt2 != 0 {
		t.Fatalf("Get recently removed key doesn't work: value %d", ansInt2)
	}
	if ansBool2 != false {
		t.Fatalf("Get recently removed key doesn't work: bool %v", ansBool2)
	}
}

func TestSameKeys(t *testing.T) {
	c := lru.New(3)
	c.Put(1, 1)
	c.Put(2, 2)
	c.Put(1, 10)
	ansInt0, _ := c.Get(1)
	if ansInt0 != 10 {
		t.Fatalf("Change values doesn't work: %v", c)
	}
}
