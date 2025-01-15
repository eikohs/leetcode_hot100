/*
 * @lc app=leetcode id=240 lang=golang
 *
 * [240] Search a 2D Matrix II
 */

// @lc code=start

func searchMatrix(matrix [][]int, target int) bool {
	// 定义起始位置
	i, j := 0, len(matrix[0])-1
	// 进行 z 字搜索
	for j >= 0 && i < len(matrix) {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] > target {
			j--
		} else if matrix[i][j] < target {
			i++
		}
	}
	return false
}

// @lc code=end

