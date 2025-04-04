/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 */

// @lc code=start
/*func threeSum(nums []int) [][]int {
	n := len(nums)
	result := [][]int{}

	comb := make([]int, 3)
	hash := make(map[string]bool)

	// 求出所有组合
	for i := 0; i < n; i++ {
		comb[0] = nums[i]
		for j := i + 1; j < n; j++ {
			comb[1] = nums[j]
			for k := j + 1; k < n; k++ {
				comb[2] = nums[k]
				if comb[0]+comb[1]+comb[2] == 0 {
					// 排序后转为字符串并查重
					minC := min(comb[0], comb[1], comb[2])
					maxC := max(comb[0], comb[1], comb[2])
					midC := 0 - minC - maxC
					str := strconv.Itoa(minC) + strconv.Itoa(midC) + strconv.Itoa(maxC)
					// 查重
					_, exists := hash[str]
					if !exists {
						combination := make([]int, 3)
						copy(combination, comb)
						result = append(result, combination)
						hash[str] = true
					}
				}
			}
		}
	}

	return result
}
*/

func threeSum(nums []int) [][]int {
	n := len(nums)
	result := [][]int{}

	// 1. 排序数组
	sort.Ints(nums) // 标准库默认是递增排序
	for i := 0; i < n; {
		// 同时进行第 2、3 层的搜索
		left, right := i+1, n-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			leftMove, rightMove := false, false
			if sum == 0 {
				combination := []int{nums[i], nums[left], nums[right]}
				result = append(result, combination)
				leftMove, rightMove = true, true
			} else if sum < 0 {
				leftMove = true
			} else {
				rightMove = true
			}
			if leftMove {
				tmp := nums[left]
				for left < right && nums[left] == tmp {
					left++
				}
			}
			if rightMove {
				tmp := nums[right]
				for left < right && nums[right] == tmp {
					right--
				}
			}
		}
		// 搜索完毕后，移动到更小的数
		tmp := nums[i]
		for i < n && nums[i] == tmp {
			i++
		}
	}

	return result
}

// @lc code=end

