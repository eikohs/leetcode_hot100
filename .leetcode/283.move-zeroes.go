/*
 * @lc app=leetcode id=283 lang=golang
 *
 * [283] Move Zeroes
 */

// @lc code=start
func moveZeroes(nums []int) {
	// 标记 0 的位置
	var zeroFlag = 0
	// 遍历寻找每一个非 0 的元素，移动到前面
	for _, num := range nums {
		if num != 0 {
			nums[zeroFlag] = num
			zeroFlag++
		}
	}
	// 将剩余的位置填充为 0
	for i := zeroFlag; i < len(nums); i++ {
		nums[i] = 0
	}
}

// @lc code=end

