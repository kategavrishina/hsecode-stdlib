// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package vector

import "math"

type Vector struct {
	Len int // actual number of occupied elements in the underlying slice
	// contains filtered or unexported fields
	vector []int
}

func New(cap int) *Vector {
	return &Vector{
		Len:    0,
		vector: make([]int, cap),
	}
}

func (a *Vector) Delete(idx int) {
	if idx >= a.Len || idx < 0 {
		panic("Index is out of range")
	}
	if a.Len < len(a.vector)*3/4 {
		newVector := make([]int, int(math.Ceil(float64(len(a.vector))*3/4)))
		if idx > 0 {
			for i := 0; i < idx; i++ {
				newVector[i] = a.vector[i]
			}
		}
		for i := idx; i < a.Len-1; i++ {
			newVector[i] = a.vector[i+1]
		}
		a.vector = newVector
	} else {
		for i := idx; i < a.Len-1; i++ {
			a.vector[i] = a.vector[i+1]
		}
		a.vector[a.Len-1] = 0
	}
	a.Len--
}

func (a *Vector) Insert(idx int, x int) {
	if idx > a.Len {
		panic("Index is bigger than length")
	}
	if a.Len == len(a.vector) {
		newVector := make([]int, 0)
		if len(a.vector) == 0 {
			newVector = make([]int, 1)
		} else {
			newVector = make([]int, int(math.Ceil(float64(len(a.vector))*4/3)))
		}
		if idx > 0 {
			for i := idx - 1; i >= 0; i-- {
				newVector[i] = a.vector[i]
			}
		}
		if idx < a.Len {
			for i := a.Len; i > idx; i-- {
				newVector[i] = a.vector[i-1]
			}
		}
		newVector[idx] = x
		a.vector = newVector
	} else {
		for i := a.Len; i > idx; i-- {
			a.vector[i] = a.vector[i-1]
		}
		a.vector[idx] = x
	}
	a.Len++
}

func (a *Vector) Pop() int {
	obj := a.Get(a.Len - 1)
	a.Delete(a.Len - 1)
	return obj
}

func (a *Vector) Push(x int) {
	a.Insert(a.Len, x)
}

func (a *Vector) Set(idx int, x int) {
	if idx >= a.Len {
		panic("Index is bigger than length")
	}
	a.vector[idx] = x
}

func (a *Vector) Get(idx int) int {
	if idx >= a.Len {
		panic("Index is bigger than length")
	}
	return a.vector[idx]
}
