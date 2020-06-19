package tree

import "sort"

func NewBST(elements []int) *Tree {
	unique := make([]int, 0)
	hash := make(map[int]struct{})
	for _, el := range elements {
		if _, ok := hash[el]; !ok {
			hash[el] = struct{}{}
			unique = append(unique, el)
		}
	}
	sort.Ints(unique)
	return fromSorted(unique)
}

func fromSorted(unique []int) *Tree {
	if len(unique) > 0 {
		mid := len(unique) / 2
		return &Tree{
			Value: unique[mid],
			Left:  NewBST(unique[:mid]),
			Right: NewBST(unique[mid+1:]),
		}
	} else {
		return nil
	}
}
