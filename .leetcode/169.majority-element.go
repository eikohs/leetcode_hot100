/*
 * @lc app=leetcode id=169 lang=golang
 *
 * [169] Majority Element
 */

// @lc code=start
func majorityElement(nums []int) int {
	count := 0
	var moore int
	for _, num := range nums {
		if count == 0 {
			moore = num
			count = 1
		} else if num == moore {
			count++
		} else {
			count--
		}
	}
	return moore
}

// @lc code=end

