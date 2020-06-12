package fulltext

import (
	"strings"
)

type Index struct {
	idx map[string][]int
}

func unique(intSlice []int) []int {
	keys := make(map[int]bool)
	list := make([]int, 0)
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func New(docs []string) *Index {
	idx := new(Index)
	idx.idx = make(map[string][]int)
	for i, doc := range docs {
		words := strings.Split(doc, " ")
		for _, word := range words {
			idx.idx[word] = append(idx.idx[word], i)
		}
	}
	return idx
}

func (idx *Index) Search(query string) []int {
	if query == "" {
		return []int{}
	}
	wquery := strings.Split(query, " ")
	indices := make(map[int]uint8)
	found := make([]int, 0)
	for i, word := range wquery {
		docIdx := unique(idx.idx[word])
		for _, id := range docIdx {
			if indices[id] < uint8(i) {
				break
			}
			indices[id]++
			if indices[id] == uint8(len(wquery)) {
				found = append(found, id)
			}
		}
	}
	return found
}
