/*
 * @lc app=leetcode id=62 lang=golang
 *
 * [62] Unique Paths
 */

// @lc code=start
/*
func uniquePaths(m int, n int) int {
	if m > n {
		m, n = n, m
	}

	result := 0
	var dfs func(int, int)
	dfs = func(idx int, remain int) {
		if idx == m-1 {
			result++
			return
		}
		for i := 0; i <= remain; i++ {
			dfs(idx+1, remain-i)
		}
	}

	dfs(0, n-1)
	return result
}
*/

func uniquePaths(m int, n int) int {
	if m < n {
		m, n = n, m
	}

	dp := make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				dp[j] = 1
			} else {
				dp[j] = dp[j] + dp[j-1]
			}
		}
	}

	return dp[n-1]
}

// @lc code=end

