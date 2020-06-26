package strings

import (
	"math"
	"strings"
)

func LCS(s1, s2 string) string {
	m := len(s1)
	n := len(s2)
	C := make([][]int, m+1)
	for i := range C {
		C[i] = make([]int, n+1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if s1[i] == s2[j] {
				C[i+1][j+1] = C[i][j] + 1
			} else {
				C[i+1][j+1] = int(math.Max(float64(C[i+1][j]), float64(C[i][j+1])))
			}
		}
	}
	index := C[m][n]
	if index == 0 {
		return ""
	}
	result := make([]string, index)
	for i, j := m, n; i > 0 && j > 0; {
		if s1[i-1] == s2[j-1] {
			result[index-1] = string(s1[i-1])
			i--
			j--
			index--
		} else if C[i][j-1] > C[i-1][j] {
			j--
		} else {
			i--
		}
	}
	return strings.Join(result, "")
}
