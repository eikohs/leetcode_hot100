/*
 * @lc app=leetcode id=494 lang=golang
 *
 * [494] Target Sum
 */

// @lc code=start
func findTargetSumWays(nums []int, target int) int {
	n := len(nums)
	sum := make([]int, n)
	sum[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		sum[i] = sum[i+1] + nums[i]
	}

	result := 0

	var dfs func(int, int)
	dfs = func(idx int, remain int) {
		if idx == n {
			if remain == 0 {
				result++
			}
			return
		}
		if remain+sum[idx] < 0 || remain-sum[idx] > 0 {
			return
		}
		dfs(idx+1, remain-nums[idx])
		dfs(idx+1, remain+nums[idx])
	}

	dfs(0, target)
	return result
}

// @lc code=end

