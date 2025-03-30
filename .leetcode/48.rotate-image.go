/*
 * @lc app=leetcode id=48 lang=golang
 *
 * [48] Rotate Image
 */

// @lc code=start
func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < (n/2)+(n%2); i++ {
		for j := 0; j <= (n/2)-1; j++ {
			tmp := matrix[n-j-1][i]
			matrix[n-j-1][i] = matrix[n-i-1][n-j-1]
			matrix[n-i-1][n-j-1] = matrix[j][n-i-1]
			matrix[j][n-i-1] = matrix[i][j]
			matrix[i][j] = tmp
		}
	}
}

// @lc code=end

