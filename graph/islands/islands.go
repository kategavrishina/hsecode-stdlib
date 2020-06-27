package islands

import "hsecode.com/stdlib/matrix/int"

func Count(grid *matrix.Matrix) int {
	numOfIslands := 0
	// For each cell - we perform dfs.
	for x := 0; x < grid.Rows; x++ {
		for y := 0; y < grid.Cols; y++ {
			// Skip performing dfs from water.
			if grid.Get(x, y) != 0 {
				islandDFS(grid, x, y)
				numOfIslands++
			}
		}
	}
	return numOfIslands
}

func islandDFS(grid *matrix.Matrix, x, y int) {
	// Boundary check.
	if x < 0 || y < 0 || x >= grid.Rows || y >= grid.Cols || grid.Get(x, y) == 0 {
		return
	}

	// Mark current cell as visited.
	grid.Set(x, y, 0)

	// A neighbor can be traversed to (top, bottom, right, left).
	islandDFS(grid, x+1, y)
	islandDFS(grid, x-1, y)
	islandDFS(grid, x, y+1)
	islandDFS(grid, x, y-1)
}
