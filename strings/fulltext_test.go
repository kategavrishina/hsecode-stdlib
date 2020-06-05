package fulltext_test

import (
	fulltext "hsecode.com/stdlib/strings"
	"testing"
)

func TestFullText(t *testing.T) {
	docs := fulltext.New([]string{
		"this is the house that jack built",
		"this is the rat that ate the malt",
	})
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
	if len(s3) != 2 {
		t.Fatalf("Two documents found doesn't work: %v", s3)
	}
}
