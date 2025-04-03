/*
 * @lc app=leetcode id=79 lang=golang
 *
 * [79] Word Search
 */

// @lc code=start
func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	flag := make([][]bool, m)
	for i := 0; i < m; i++ {
		flag[i] = make([]bool, n)
	}
	result := false

	var dfs func(int, int, int)
	dfs = func(x int, y int, idx int) {
		if result {
			return
		}
		if idx == len(word)-1 {
			result = true
			return
		}
		next := idx + 1
		flag[x][y] = true
		if x > 0 && !flag[x-1][y] && word[next] == board[x-1][y] {
			dfs(x-1, y, next)
		}
		if x+1 < m && !flag[x+1][y] && word[next] == board[x+1][y] {
			dfs(x+1, y, next)
		}
		if y > 0 && !flag[x][y-1] && word[next] == board[x][y-1] {
			dfs(x, y-1, next)
		}
		if y+1 < n && !flag[x][y+1] && word[next] == board[x][y+1] {
			dfs(x, y+1, next)
		}
		flag[x][y] = false
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[0] {
				dfs(i, j, 0)
				if result {
					return result
				}
			}
		}
	}

	return result
}

// @lc code=end

