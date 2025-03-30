/*
 * @lc app=leetcode id=221 lang=golang
 *
 * [221] Maximal Square
 */

// @lc code=start
func checkByte(val byte) int {
	if val == '1' {
		return 1
	} else {
		return 0
	}
}

func maximalSquare(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	if m <= 0 || n <= 0 {
		return 0
	}
	max := 0
	dp := make([][]int, m)
	for key, _ := range dp {
		dp[key] = make([]int, n)
	}
	for i := range m {
		for j := range n {
			if i == 0 || j == 0 || checkByte(matrix[i][j]) == 0 {
				dp[i][j] = checkByte(matrix[i][j])
				if max == 0 && dp[i][j] == 1 {
					max = 1
				}
				continue
			}
			dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
			if dp[i][j] > max {
				max = dp[i][j]
			}
		}
	}
	return max * max
}

// @lc code=end

