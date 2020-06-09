package fulltext

import (
	"sort"
	"strings"
)

type Index struct {
	idx map[string][]int
}

func unique(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}
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
	for k, v := range idx.idx {
		sort.Ints(v)
		idx.idx[k] = unique(v)
	}
	return idx
}

func (idx *Index) Search(query string) []int {
	if query == "" {
		return []int{}
	}

	wquery := strings.Split(query, " ")
	// queryhash := make(map[string][]int)
	length := len(wquery)
	indices := make(map[int]int)
	found := make([]int, 0)
	for _, word := range wquery {
		// docIdx := idx.idx[word]
		// queryhash[word] = docIdx
		docIdx := idx.idx[word]
		for _, id := range docIdx {
			indices[id]++
		}
	}
	//first := queryhash[wquery[0]]
	/*
		for _, id := range first {
			indices[id] = 0
		}

		for _, v := range queryhash {
			for _, id := range v {
				_, ok := indices[id]
				if ok {
					indices[id]++
				}
			}
		}
	*/
	for k, v := range indices {
		if v == length {
			found = append(found, k)
		}
	}
	sort.Ints(found)
	return found
}
