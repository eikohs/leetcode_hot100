/*
 * @lc app=leetcode id=56 lang=golang
 *
 * [56] Merge Intervals
 */

// @lc code=start
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	ans := make([][]int, 0)
	for _, interval := range intervals {
		if len(ans) == 0 || ans[len(ans)-1][1] < interval[0] {
			ans = append(ans, interval)
		} else {
			ans[len(ans)-1][1] = max(ans[len(ans)-1][1], interval[1])
		}
	}
	return ans
}

// @lc code=end

