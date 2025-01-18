/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	// 定义结果数组
	rst := make([]int, 2)
	// 定义哈希
	hash := make(map[int]int)
	// 遍历数组
	for key, num := range nums {
		if hash[num] != 0 {
			rst[0], rst[1] = hash[num]-1, key
			break
		} else {
			hash[target-num] = key + 1
		}
	}
	return rst
}

// @lc code=end

