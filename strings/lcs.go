package strings

import (
	"math"
	"strings"
)

func LCS(s1, s2 string) string {
	m := len(s1)
	n := len(s2)
	// C := matrix.New(m+1, n+1)
	C := make([][]int, m+1)
	for i := range C {
		C[i] = make([]int, n+1)
	}
	for i := 0; i < m+1; i++ {
		for j := 0; j < n+1; j++ {
			if i == 0 || j == 0 {
				C[i][j] = 0
				// C.Set(i, j, 0)
			} else if s1[i-1] == s2[j-1] {
				C[i][j] = C[i-1][j-1] + 1
				// C.Set(i, j, C.Get(i-1, j-1)+1)
			} else {
				C[i][j] = int(math.Max(float64(C[i][j-1]), float64(C[i-1][j])))
				// C.Set(i, j, int(math.Max(float64(C.Get(i, j-1)), float64(C.Get(i-1, j)))))
			}
		}
	}
	// index := C.Get(m, n)
	index := C[m][n]
	if index == 0 {
		return ""
	}
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
			// } else if C.Get(i, j-1) > C.Get(i-1, j) {
		} else if C[i][j-1] > C[i-1][j] {
			j--
		} else {
			i--
		}
	}
	return strings.Join(result[:len(result)-1], "")
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
