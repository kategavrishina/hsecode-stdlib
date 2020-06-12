package fulltext_test

import (
	fulltext "hsecode.com/stdlib/strings/fulltext"
	"testing"
)

func TestFullText(t *testing.T) {
	docs := fulltext.New([]string{
		"this is the house that jack built",
		"this the rat that ate the malt is",
		"this the is the",
		"rate the eat this jack is",
		"rate the eat this jack is",
	})
	t.Log(docs)
	s0 := docs.Search("")
	if len(s0) > 0 {
		t.Fatalf("Empty search doesn't work: %v", s0)
	}
	s1 := docs.Search("in the house that jack built")
	if len(s1) > 0 {
		t.Fatalf("No documents found doesn't work: %v", s1)
	}
	s2 := docs.Search("malt rat")
	if len(s2) != 1 {
		t.Fatalf("One document found doesn't work: %v", s2)
	}
	s3 := docs.Search("is this the")
	if s3[1] != 1 {
		t.Fatalf("Two documents found doesn't work: %v", s3)
	}
	if len(s3) != 5 {
		t.Fatalf("Two documents found doesn't work: %v", s3)
	}
	t.Log(s3)
}

func TestDuplicate(t *testing.T) {
	docs := fulltext.New([]string{
		"this is the house that jack built",
		"house that jack built this is the",
	})
	s0 := docs.Search("the jack")
	if s0[1] != 1 {
		t.Fatalf("Duplicate documents doesn't work: %v", s0)
	}
	if len(s0) != 2 {
		t.Fatalf("Duplicate documents doesn't work: %v", s0)
	}
}
