package pick

import (
	"sort"
)

func NthElement(data sort.Interface, nth int) {
	q := partition(data)
	if q < nth {
		// partition(left) || partition(right)
	} // else break function
}

func partition(data sort.Interface) int {
	px := data.Len() / 2
	pv := data[px]
	last := data.Len() - 1
}
