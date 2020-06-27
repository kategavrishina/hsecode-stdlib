package tree

func (T *Tree) NoLeft() *Tree {
	if T.Left == nil {
		return T
	}
	l := T.Left
	for l.Left.Left != nil {
		l = l.Left
	}

	R := &Tree{
		Value: T.Value}
	return R
}
