/*
 * @lc app=leetcode id=34 lang=golang
 *
 * [34] Find First and Last Position of Element in Sorted Array
 */

// @lc code=start
func searchRange(nums []int, target int) []int {
	var binarySearch func(int, int, bool) int
	binarySearch = func(left int, right int, leftBound bool) int {
		rst := -1
		for left <= right {
			mid := (left + right) / 2
			if nums[mid] == target {
				rst = mid
				if leftBound {
					right = mid - 1
				} else {
					left = mid + 1
				}
			} else if nums[mid] > target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
		return rst
	}

	n := len(nums)
	left, right := binarySearch(0, n-1, true), binarySearch(0, n-1, false)
	return []int{left, right}
}

// @lc code=end

