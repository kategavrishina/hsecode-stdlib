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

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
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
	queryhash := make(map[string][]int)
	for _, word := range wquery {
		doc_idx := idx.idx[word]
		queryhash[word] = doc_idx
	}
	length := len(wquery)
	found := make([]int, 0)
	first := queryhash[wquery[0]]
	for _, id := range first {
		l := 0
		for _, v := range queryhash {
			if contains(v, id) {
				l++
			}
		}
		if l == length {
			found = append(found, id)
		}
	}
	return found
}
