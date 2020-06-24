package strings_test

import (
	"hsecode.com/stdlib/strings"
	"testing"
)

func isWord(w string) bool {
	if w == "abcd" || w == "abc" || w == "aaa" || w == "bc" {
		return true
	} else {
		return false
	}
}

func TestSegmentation(t *testing.T) {
	answer, _ := strings.Segmentation("abcdaaabcaaabc", isWord)
	t.Log(answer)
}
