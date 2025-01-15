/*
 * @lc app=leetcode id=560 lang=golang
 *
 * [560] Subarray Sum Equals K
 */

// @lc code=start
func oldSubarraySum(nums []int, k int) int {
	n := len(nums)
	// 定义 dp
	dp := make([]int, n)
	// 动态规划并解决问题
	count := 0
	dp[0] = nums[0]
	for i := 0; i < n; i++ {
		if i > 0 {
			dp[i] -= nums[i-1]
		}
		if dp[i] == k {
			count++
		}
		for j := i + 1; j < n; j++ {
			dp[j] = dp[j-1] + nums[j]
			if dp[j] == k {
				count++ // 找到合适的子数组
			}
		}
	}
	return count
}

func subarraySum(nums []int, k int) int {
	count := 0
	sum := 0
	// 通过哈希表记录 sum[j] - k 的值
	hash := make(map[int]int)
	hash[k] = 1 // 初始化为 1 是因为 sum[j] - k = 0 的情况，即子数组的和等于 k
	for _, num := range nums {
		sum += num
		// 查找哈希表
		count += hash[sum]
		// 记录 sum[i] + k 的值
		hash[sum+k]++
	}
	return count
}

// @lc code=end

