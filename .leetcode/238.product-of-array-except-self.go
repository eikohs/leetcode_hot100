/*
 * @lc app=leetcode id=238 lang=golang
 *
 * [238] Product of Array Except Self
 */

// @lc code=start
func productExceptSelf(nums []int) []int {
	n := len(nums)
	dp := make([]int, n)
	dp[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		dp[i] = dp[i+1] * nums[i+1]
	}

	prev := 1
	for i := 0; i < n; i++ {
		next := prev * nums[i]
		nums[i] = prev * dp[i]
		prev = next
	}

	return nums
}

// @lc code=end

