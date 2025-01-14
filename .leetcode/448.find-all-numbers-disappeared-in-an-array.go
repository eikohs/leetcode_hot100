/*
 * @lc app=leetcode id=448 lang=golang
 *
 * [448] Find All Numbers Disappeared in an Array
 */

// @lc code=start
func oldFindDisappearedNumbers(nums []int) []int {
	n := len(nums)
	counts := make([]int, n)
	for _, num := range nums {
		counts[num-1]++ // 记录出现次数
	}
	index := 0
	for key, count := range counts {
		if count == 0 {
			// 筛选出未出现的数字
			nums[index] = key + 1
			index++
		}
	}
	// 截断数组并返回
	return nums[:index]
}

func findDisappearedNumbers(nums []int) []int {
	n := len(nums)
	rst := make([]int, 0)
	for _, num := range nums {
		nums[(num-1)%n] += n // 标记出现的数字
	}
	// 统计未出现的数字
	for key, num := range nums {
		if num <= n {
			rst = append(rst, key+1)
		}
	}
	return rst
}

// @lc code=end

