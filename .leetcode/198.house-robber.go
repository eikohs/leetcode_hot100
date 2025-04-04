/*
 * @lc app=leetcode id=198 lang=golang
 *
 * [198] House Robber
 */

// @lc code=start
func rob(nums []int) int {
	n := len(nums)

	robbed, unrob := 0, 0

	for i := 0; i < n; i++ {
		tmp := unrob
		unrob = max(robbed, unrob)
		robbed = nums[i] + tmp
	}

	return max(robbed, unrob)
}

// @lc code=end

