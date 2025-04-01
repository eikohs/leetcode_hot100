/*
 * @lc app=leetcode id=215 lang=golang
 *
 * [215] Kth Largest Element in an Array
 */

// @lc code=start
/*
func findKthLargest(nums []int, k int) int {
	bench := nums[0]
	// 左右指针进行二分
	left, right := 0, len(nums)-1
	for left < right {
		for left < right && nums[right] <= bench {
			right--
		}
		nums[left] = nums[right]
		nums[right] = bench

		for left < right && nums[left] >= bench {
			left++
		}
		nums[right] = nums[left]
		nums[left] = bench
	}
	mark := right + 1
	// 根据情况返回结果或者继续二分
	if mark == k {
		return bench
	} else if mark > k {
		return findKthLargest(nums[:left], k)
	} else {
		return findKthLargest(nums[mark:], k-mark)
	}
}
*/

func findKthLargest(nums []int, k int) int {
	k = k - 1

	var quickSort func(int, int) int
	quickSort = func(left int, right int) int {
		l, r := left-1, right+1
		if left == right {
			return nums[k]
		}
		bench := nums[left]
		for l < r {
			for l++; nums[l] > bench; l++ {
			}
			for r--; nums[r] < bench; r-- {
			}
			if l < r {
				nums[l], nums[r] = nums[r], nums[l]
			}
		}
		if k <= r {
			return quickSort(left, r)
		} else {
			return quickSort(r+1, right)
		}
	}

	return quickSort(0, len(nums)-1)
}

// @lc code=end

