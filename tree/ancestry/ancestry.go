package ancestry

import (
	"hsecode.com/stdlib/tree"
)

var time int

type Ancestry struct {
	// contains filtered or unexported fields
	Hash map[int]*Time
}

type Time struct {
	enter, exit int
}

func New(T *tree.Tree) *Ancestry {
	A := &Ancestry{
		Hash: make(map[int]*Time),
	}
	DFS(T, A)
	return A
}

func DFS(T *tree.Tree, A *Ancestry) {
	if T == nil {
		return
	}
	A.Hash[T.Value] = &Time{}
	A.Hash[T.Value].enter = time
	time++
	DFS(T.Left, A)
	DFS(T.Right, A)
	A.Hash[T.Value].exit = time
	time++
}

func (A *Ancestry) IsDescendant(a, b int) bool {
	hash := A.Hash
	if hash[a].enter < hash[b].enter && hash[a].exit > hash[b].exit {
		return true
	}
	return false
}
