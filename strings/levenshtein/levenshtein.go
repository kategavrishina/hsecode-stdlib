package levenshtein

import (
	"math"
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
	L := make([][]int, n)
	for i := range L {
		L[i] = make([]int, m)
	}
	for j := 0; j < n; j++ {
		L[j][0] = j
	}
	for i := 1; i < m; i++ {
		L[0][i] = i
		for j := 1; j < n; j++ {
			if src[i-1] == dst[j-1] {
				L[j][i] = L[j-1][i-1]
			} else {
				ins := L[j][i-1]
				del := L[j-1][i]
				rep := L[j-1][i-1]
				L[j][i] = int(math.Min(float64(rep), math.Min(float64(ins), float64(del)))) + 1
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
	return ls.L[ls.m-1][ls.n-1]
	// return ls.L
}

func (ls *Levenshtein) Transcript() string {
	L := ls.L
	m := ls.m - 1
	n := ls.n - 1
	if m == 0 {
		return strings.Repeat("I", n)
	}
	if n == 0 {
		return strings.Repeat("D", m)
	}
	result := make([]string, 0)
	for i, j := m, n; i > 0 && j > 0; {
		current := L[j][i]
		upperLeft := L[j-1][i-1]
		left := L[j-1][i]
		upper := L[j][i-1]
		minimum := min(upperLeft, left, upper)
		if ls.src[i-1] == ls.dst[j-1] {
			result = append(result, "M")
			i--
			j--
		} else if minimum == left {
			result = append(result, "I")
			j--
		} else if minimum == upper {
			result = append(result, "D")
			i--
		} else {
			if upperLeft == current {
				result = append(result, "M")
			} else {
				result = append(result, "R")
			}
			i--
			j--
		}
		/*
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
		*/
	}
	converted := make([]string, len(result))
	for l := len(result); l > 0; l-- {
		converted[len(result)-l] = result[l-1]
	}
	return strings.Join(converted, "")
}
