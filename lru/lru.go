package lru

import "container/list"

type Cache struct {
	// contains filtered or unexported fields
	elements *list.List
	hash     map[int]*list.Element
	cap      int
}

type Pair struct {
	key, value int
}

func New(capacity int) *Cache {
	c := new(Cache)
	c.elements = list.New()
	c.cap = capacity
	c.hash = make(map[int]*list.Element)

	return c
}

func (cache *Cache) Get(key int) (int, bool) {
	elem, ok := cache.hash[key]
	if ok {
		cache.elements.MoveToFront(elem)
		val := elem.Value.(Pair).value
		return val, true
	}
	return 0, false
}

func (cache *Cache) Put(key int, value int) {
	elem, ok := cache.hash[key]
	if !ok {
		if cache.elements.Len() == cache.cap {
			el := cache.elements.Back()
			delete(cache.hash, el.Value.(Pair).key)
			cache.elements.Remove(el)
		}
		cache.elements.PushFront(Pair{
			key,
			value,
		})
		cache.hash[key] = cache.elements.Front()
	} else {
		if elem.Value.(Pair).value != value {
			elem.Value = Pair{
				key,
				value,
			}
		}
		cache.elements.MoveToFront(elem)
		cache.hash[key] = cache.elements.Front()
	}
}
