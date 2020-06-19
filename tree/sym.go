package tree

func (T *Tree) IsSym() bool {
	if T == nil {
		return true
	}
	return IsAntiSym(T.Right, T.Left)
}

func IsAntiSym(a, b *Tree) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return IsAntiSym(a.Left, b.Right) && IsAntiSym(a.Right, b.Left) && a.Value == b.Value
}
