package ancestry

import (
	"hsecode.com/stdlib/tree"
)

var time int
var A *Ancestry

type Ancestry struct {
	// contains filtered or unexported fields
	Hash map[int]*Time
}

type Time struct {
	enter, exit int
}

func New(T *tree.Tree) *Ancestry {
	A = &Ancestry{
		Hash: make(map[int]*Time),
	}
	DFS(T)
	return A
}

func DFS(T *tree.Tree) {
	if T == nil {
		return
	}
	A.Hash[T.Value] = &Time{time, 0}
	time++
	if T.Left != nil {
		DFS(T.Left)
	}
	if T.Right != nil {
		DFS(T.Right)
	}
	A.Hash[T.Value].exit = time
	time++
}

func (A *Ancestry) IsDescendant(a, b int) bool {
	if A.Hash[a].enter < A.Hash[b].enter && A.Hash[a].exit > A.Hash[b].exit {
		return true
	}
	return false
}
