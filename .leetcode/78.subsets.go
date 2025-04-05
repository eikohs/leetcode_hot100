/*
 * @lc app=leetcode id=78 lang=golang
 *
 * [78] Subsets
 */

// @lc code=start
func subsets(nums []int) [][]int {
	n := len(nums)
	subset := make([]int, 0, n)
	result := [][]int{}

	var dfs func(int)
	dfs = func(idx int) {
		if idx == n {
			sst := make([]int, len(subset))
			copy(sst, subset)
			result = append(result, sst)
			return
		}
		// add this elem
		subset = append(subset, nums[idx])
		dfs(idx + 1)
		subset = subset[:len(subset)-1]
		// not add
		dfs(idx + 1)
	}

	dfs(0)
	return result
}

// @lc code=end

