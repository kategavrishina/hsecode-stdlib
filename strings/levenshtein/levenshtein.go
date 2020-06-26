package levenshtein

import "strings"

type Levenshtein struct {
	// contains filtered or unexported fields
	m   int
	n   int
	L   [][]int
	src string
	dst string
}

func min(x, y, z int) int {
	if x <= y && x <= z {
		return x
	} else if y <= z && y <= x {
		return y
	}
	return z
}

func New(src, dst string) *Levenshtein {
	m := len(src) + 1
	n := len(dst) + 1
	L := make([][]int, m)
	for i := range L {
		L[i] = make([]int, n)
	}
	for j := 0; j < n; j++ {
		L[0][j] = j
	}
	for i := 1; i < m; i++ {
		L[i][0] = i
		for j := 1; j < n; j++ {
			if src[i-1] == dst[j-1] {
				L[i][j] = L[i-1][j-1]
			} else {
				ins := L[i][j-1]
				del := L[i-1][j]
				rep := L[i-1][j-1]
				L[i][j] = min(ins, del, rep) + 1
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

func (ls *Levenshtein) Distance() [][]int {
	return ls.L
}

func (ls *Levenshtein) Transcript() string {
	L := ls.L
	m := ls.m
	n := ls.n
	result := make([]string, 0)
	for i, j := m-1, n-1; i > 0 && j > 0; {
		upperLeft := L[i-1][j-1]
		left := L[i-1][j]
		upper := L[i][j-1]
		minimum := min(upperLeft, left, upper)
		if ls.src[i-1] == ls.dst[j-1] {
			result = append(result, "M")
			i--
			j--
		} else if minimum == left {
			result = append(result, "D")
			i--
		} else if minimum == upper {
			result = append(result, "I")
			j--
		} else {
			result = append(result, "R")
			i--
			j--
		}
	}
	converted := make([]string, len(result))
	for l := len(result); l > 0; l-- {
		converted[len(result)-l] = result[l-1]
	}
	return strings.Join(converted, "")
}
