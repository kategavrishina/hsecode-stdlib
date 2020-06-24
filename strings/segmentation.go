package strings

import "errors"

func Segmentation(s string, isWord func(w string) bool) ([]string, error) {
	if len(s) == 0 {
		return []string{}, nil
	}
	n := len(s)
	split := make([]bool, len(s)+1)
	split[n] = true
	for r := n; r > 0; r-- {
		if split[r] {
			for l := 0; l < r; l++ {
				split[l] = isWord(s[l:r])
			}
		}
	}
	result := make([]string, 0)
	if !split[0] {
		return nil, errors.New("string is unsplittable")
	} else {
		l, r := 0, 0
		for i, j := range split {
			if j {
				r = i
				result = append(result, s[l:r])
				l = i
			}
		}
	}
	return result, nil
}
