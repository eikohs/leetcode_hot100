/*
 * @lc app=leetcode id=200 lang=golang
 *
 * [200] Number of Islands
 */

// @lc code=start
func numIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])

	var dfs func(int, int)
	dfs = func(x int, y int) {
		grid[x][y] = '0'
		if x > 0 && grid[x-1][y] == '1' {
			dfs(x-1, y)
		}
		if y > 0 && grid[x][y-1] == '1' {
			dfs(x, y-1)
		}
		if x+1 < m && grid[x+1][y] == '1' {
			dfs(x+1, y)
		}
		if y+1 < n && grid[x][y+1] == '1' {
			dfs(x, y+1)
		}
	}

	result := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				dfs(i, j)
				result++
			}
		}
	}

	return result
}

// @lc code=end

