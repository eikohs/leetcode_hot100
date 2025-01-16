/*
 * @lc app=leetcode id=64 lang=golang
 *
 * [64] Minimum Path Sum
 */

// @lc code=start
func minPathSum(grid [][]int) int {
	// 定义 dp 数组
	m, n := len(grid), len(grid[0])
	dp := make([]int, n)
	// 动态规划求解
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				dp[0] = grid[0][0]
			} else if i == 0 {
				dp[j] = dp[j-1] + grid[0][j]
			} else if j == 0 {
				dp[0] += grid[i][j]
			} else {
				dp[j] = min(dp[j], dp[j-1]) + grid[i][j]
			}
		}
	}
	return dp[n-1]
}

// @lc code=end

