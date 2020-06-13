package pick

import (
	"math/rand"
	"sort"
)

func NthElement(data sort.Interface, nth int) {
	left := 0
	right := data.Len() - 1
	for {
		pivotIndex := rand.Intn(right+1-left) + left
		pivotIndex = partition(data, left, right, pivotIndex)
		if pivotIndex < nth {
			left = pivotIndex + 1
		} else if nth < pivotIndex {
			right = pivotIndex - 1
		} else {
			return
		}
	}
}

func partition(data sort.Interface, left, right, pivotIndex int) int {
	storeIndex := left
	data.Swap(pivotIndex, right)
	for i := left; i < right; i++ {
		if data.Less(i, right) {
			data.Swap(i, storeIndex)
			storeIndex++
		}
	}
	data.Swap(storeIndex, right)
	return storeIndex
}
