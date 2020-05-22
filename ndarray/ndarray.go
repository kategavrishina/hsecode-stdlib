package ndarray

type NDArray struct {
	// contains filtered or unexported fields
	shape []int
}

func New(shape ...int) *NDArray {
	for _, s := range shape {
		if s < 0 {
			panic("Dimensions are negative")
		}
	}
	return &NDArray{
		shape: shape,
	}
}

func (nda *NDArray) Idx(indices ...int) int {
	if len(indices) != len(nda.shape) {
		panic("Dimensions mismatch (overall)")
	}

	for i, ind := range indices {
		if ind < 0 {
			panic("Index is negative")
		}
		if ind >= nda.shape[i] {
			panic("Dimensions mismatch (particular index)")
		}
	}

	c := indices[0]
	for i, ind := range indices[1:] {
		N := nda.shape[i+1]
		c = ind + N*c
	}
	return c
}
