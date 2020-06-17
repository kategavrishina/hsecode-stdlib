package pick

import "container/heap"

type Heap struct {
	indices []int
	data    Ordered
}

func (h *Heap) Len() int {
	return len(h.indices)
}
func (h *Heap) Swap(i, j int) {
	h.indices[i], h.indices[j] = h.indices[j], h.indices[i]
}
func (h *Heap) Less(i, j int) bool {
	return !h.data.Less(h.indices[i], h.indices[j])
}
func (h *Heap) Push(x interface{}) {
	h.indices = append(h.indices, x.(int))
}
func (h *Heap) Pop() interface{} {
	old := h.indices
	n := len(old)
	x := old[n-1]
	h.indices = old[:n-1]
	return x
}

type Ordered interface {
	Len() int
	Less(i, j int) bool
}

func FirstN(data Ordered, n int) []int {
	var allInds []int
	for i := 0; i < data.Len(); i++ {
		allInds = append(allInds, i)
	}
	if data.Len() <= n {
		return allInds
	}
	h := Heap{
		indices: make([]int, 0),
		data:    data,
	}
	heap.Init(&h)
	for j := 0; j < n; j++ {
		heap.Push(&h, j)
	}
	for i := n; i < len(allInds); i++ {
		if !data.Less(h.indices[0], i) {
			heap.Remove(&h, 0)
			heap.Push(&h, i)
			heap.Fix(&h, n-1)
		}
	}
	return h.indices
}
