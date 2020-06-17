package xlist

import "container/list"

func Sort(data *list.List, less func(a, b *list.Element) bool) {
	length := data.Len()
	for k := 1; k < length; k *= 2 {
		for start := 0; start+k < length; start += k * 2 {
			mid := start + k
			end := mid + k
			if end > length {
				end = length
			}
		}
	}
}
