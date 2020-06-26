package levenshtein_test

import (
	"hsecode.com/stdlib/strings/levenshtein"
	"testing"
)

func TestLevenshtein_Distance(t *testing.T) {
	ls := levenshtein.New("fg", "gd")
	t.Log(ls.Distance())
}

func TestLevenshtein_Transcript(t *testing.T) {
	ls := levenshtein.New("algorithm", "altruistic")
	t.Log(ls.Transcript())
	t.Log(len(ls.Transcript()))
	ls = levenshtein.New("fg", "gd")
	t.Log(ls.Transcript())
}
