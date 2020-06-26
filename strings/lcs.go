package strings

import (
	matrix "hsecode.com/stdlib/matrix/int"
	"math"
	"strings"
)

func LCS(s1, s2 string) string {
	m := len(s1)
	n := len(s2)
	C := matrix.New(m+1, n+1)
	for i := 0; i < m+1; i++ {
		C.Set(i, 0, 0)
	}
	for j := 0; j < n+1; j++ {
		C.Set(0, j, 0)
	}
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if s1[i-1] == s2[j-1] {
				C.Set(i, j, C.Get(i-1, j-1)+1)
			} else {
				C.Set(i, j, int(math.Max(float64(C.Get(i, j-1)), float64(C.Get(i-1, j)))))
			}
		}
	}
	index := C.Get(m, n)
	result := make([]string, index+1)
	result[index] = "\000"

	i := m
	j := n
	for i > 0 && j > 0 {
		if s1[i-1] == s2[j-1] {
			result[index-1] = string(s1[i-1])
			i--
			j--
			index--
		} else if C.Get(i, j-1) > C.Get(i-1, j) {
			j--
		} else {
			i--
		}
	}
	return strings.Join(result, "")
}

/*
	return Backtrack(C, s1, s2, m, n)
}

func Backtrack(C *matrix.Matrix, s1, s2 string, i, j int) string {
	if i == 0 || j == 0 {
		return ""
	}
	if s1[i-1] == s2[j-1] {
		return Backtrack(C, s1, s2, i-1, j-1) + string(s1[i-1])
	}
	if C.Get(i, j-1) > C.Get(i-1, j) {
		return Backtrack(C, s1, s2, i, j-1)
	}
	return Backtrack(C, s1, s2, i-1, j)
}
*/
