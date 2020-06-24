package strings_test

import (
	"hsecode.com/stdlib/strings"
	"testing"
)

func isWord(w string) bool {
	/*
		for i := range w {
			for j := range w {
				if j != i {
					return false
				}
			}
		}
		return true
	*/
	if w == "abcd" || w == "abc" || w == "aaa" || w == "bc" {
		return true
	} else {
		return false
	}
}

func TestSegmentation(t *testing.T) {
	answer, err := strings.Segmentation("aaaaaabcaaaaaabc", isWord)
	t.Log(answer, err)
}
