/*
 * @lc app=leetcode id=55 lang=golang
 *
 * [55] Jump Game
 */

// @lc code=start
/* first take
func canJump(nums []int) bool {
	n := len(nums)
	flag := make([]int, n)
	flag[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		for j := 1; j <= nums[i]; j++ {
			if i+j < n && flag[i+j] == 1 {
				flag[i] = 1
				break
			}
		}
	}
	return flag[0] == 1
}
*/

func canJump(nums []int) bool {
	target := len(nums) - 1
	for i := target - 1; i >= 0; i-- {
		nums[i] = nums[i] - target + i
		if nums[i] >= 0 {
			target = i
		}
	}
	return nums[0] >= 0
}

// @lc code=end

