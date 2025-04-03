/*
 * @lc app=leetcode id=72 lang=golang
 *
 * [72] Edit Distance
 */

// @lc code=start
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)

	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
		for j := 0; j <= n; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if i > 0 && j > 0 {
				if word1[i-1] == word2[j-1] {
					dp[i][j] = min(dp[i-1][j-1], dp[i-1][j]+1, dp[i][j-1]+1)
				} else {
					dp[i][j] = min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1]) + 1
				}
			} else if i == 0 {
				dp[i][j] = dp[i][j-1] + 1
			} else {
				dp[i][j] = dp[i-1][j] + 1
			}
		}
	}

	return dp[m][n]
}

// @lc code=end

