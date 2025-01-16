/*
 * @lc app=leetcode id=136 lang=golang
 *
 * [136] Single Number
 */

// @lc code=start
func singleNumber(nums []int) int {
	// 异或求出结果
	result := 0
	for _, num := range nums {
		result ^= num
	}
	return result
}

// @lc code=end

