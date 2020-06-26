package levenshtein

import (
	"strings"
)

type Levenshtein struct {
	// contains filtered or unexported fields
	m   int
	n   int
	L   [][]int
	src string
	dst string
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func New(src, dst string) *Levenshtein {
	m := len(src) + 1
	n := len(dst) + 1
	L := make([][]int, n)
	for i := range L {
		L[i] = make([]int, m)
	}
	for i := 1; i < m; i++ {
		L[0][i] = i
		for j := 1; j < n; j++ {
			L[j][0] = j
			if src[i-1] == dst[j-1] {
				L[j][i] = L[j-1][i-1]
			} else {
				ins := L[j][i-1]
				del := L[j-1][i]
				rep := L[j-1][i-1]
				L[j][i] = min(rep, min(del, ins)) + 1
			}
		}
	}
	return &Levenshtein{
		m:   m,
		n:   n,
		L:   L,
		src: src,
		dst: dst,
	}
}

func (ls *Levenshtein) Distance() int {
	return ls.L[ls.n-1][ls.m-1]
	//return ls.L
}

func (ls *Levenshtein) Transcript() string {
	L := ls.L
	m := ls.m - 1
	n := ls.n - 1
	result := make([]string, 0)
	for i, j := m, n; !(i < 0 || j < 0) && !(i == 0 && j == 0); {
		if j == 0 {
			result = append(result, "D")
			i--
		} else if i == 0 {
			result = append(result, "I")
			j--
		} else {
			minimum := min(L[j-1][i-1], min(L[j-1][i], L[j][i-1]))
			if ls.src[i-1] == ls.dst[j-1] {
				result = append(result, "M")
				i--
				j--
			} else if minimum == L[j-1][i] {
				result = append(result, "I")
				j--
			} else if minimum == L[j][i-1] {
				result = append(result, "D")
				i--
			} else {
				if L[j-1][i-1] == L[j][i] {
					result = append(result, "M")
				} else {
					result = append(result, "R")
				}
				i--
				j--
			}
		}
	}
	converted := make([]string, len(result))
	for l := len(result); l > 0; l-- {
		converted[len(result)-l] = result[l-1]
	}
	return strings.Join(converted, "")
}
