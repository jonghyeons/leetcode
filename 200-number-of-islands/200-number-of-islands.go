func numIslands(grid [][]byte) int {
    res := 0
	for x, v := range grid {
		for y := range v {
			if dfs(x, y, grid) {
				res += 1
			}
		}
	}
	return res
}

func dfs(x,y int, grid[][]byte)bool{
    if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) {
		return false
	}

	if grid[x][y] == '1' {
		grid[x][y] = '0'

		dfs(x-1, y, grid)
		dfs(x+1, y, grid)
		dfs(x, y-1, grid)
		dfs(x, y+1, grid)

		return true
	}
	return false
}