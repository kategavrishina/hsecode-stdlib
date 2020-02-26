package matrix

import "github.com/cheekybits/genny/generic"

//go:generate genny -in=$GOFILE -out=int/dont_edit.go gen "ValueType=int"

type ValueType generic.Type

type Matrix struct {
	Rows   int
	Cols   int
	matrix []ValueType
}

func New(n, m int) *Matrix {
	// mat := new(Matrix)
	// mat.matrix = make([]ValueType, n*m)
	// mat.Rows = n
	// mat.Cols = m

	return &Matrix{
		Rows:   n,
		Cols:   m,
		matrix: make([]ValueType, n*m),
	}

	// return mat
}

func (M *Matrix) Get(i, j int) ValueType {
	if i < 0 || i >= M.Rows || j < 0 || j >= M.Cols {
		panic("index error")
	}
	idx := M.Cols*i + j

	return M.matrix[idx]
}

func (M *Matrix) Set(i, j int, v ValueType) {
	idx := M.Cols*i + j
	M.matrix[idx] = v
}
