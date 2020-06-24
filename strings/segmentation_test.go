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
	if w != "abcd" {
		return false
	} else {
		return true
	}
}

func TestSegmentation(t *testing.T) {
	answer, err := strings.Segmentation("abcdabcdabcd", isWord)
	t.Log(answer, err)
}
