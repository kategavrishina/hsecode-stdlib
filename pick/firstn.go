package pick

type Ordered interface {
	// Len returns the number of elements in the collection.
	Len() int
	// Less reports whether the element with index i should sort before the element with index j.
	Less(i, j int) bool
}

func FirstN(data Ordered, n int) []int {
	return []int{}
}
