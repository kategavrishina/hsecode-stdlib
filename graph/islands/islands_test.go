package islands_test

import (
	"hsecode.com/stdlib/graph/islands"
	matrix "hsecode.com/stdlib/matrix/int"
	"testing"
)

// Three islands
var rows = [][]int{
	{1, 0, 1},
	{0, 0, 1},
	{1, 1, 0},
}

func TestExample(t *testing.T) {
	grid := matrix.New(3, 3)

	// Fill in the matrix
	for i, row := range rows {
		for j, v := range row {
			grid.Set(i, j, v)
		}
	}

	t.Log(islands.Count(grid))
	// Output: 3
}
