package levenshtein_test

import (
	"hsecode.com/stdlib/strings/levenshtein"
	"testing"
)

func TestLevenshtein_Distance(t *testing.T) {
	ls := levenshtein.New("abcdef", "azced")
	t.Log(ls.Distance())
}

func TestLevenshtein_Transcript(t *testing.T) {
	ls := levenshtein.New("vintner", "writers")
	t.Log(ls.Transcript())
	t.Log(len(ls.Transcript()))
}
