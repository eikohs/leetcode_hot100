/*
 * @lc app=leetcode id=287 lang=golang
 *
 * [287] Find the Duplicate Number
 */

// @lc code=start
func findDuplicate(nums []int) int {
	// 快慢节点到相遇
	slow := nums[0]
	fast := nums[slow]
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}
	// 找 k
	slow = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}

// @lc code=end

