/*
 * @lc app=leetcode id=338 lang=golang
 *
 * [338] Counting Bits
 */

// @lc code=start
func countBits(n int) []int {
	ans := make([]int, n+1)
	var count int
	ans[0] = 0
	for i := 1; i <= n; i++ {
		// 判断奇偶
		if i%2 == 1 {
			// 奇数，则比上一个数多一位
			ans[i] = ans[i-1] + 1
			// 计算末尾连续的1
			tmp := i
			count = 0
			for tmp != 0 && tmp%2 == 1 {
				count++
				tmp >>= 1
			}
		} else {
			// 偶数，按照公式计算
			ans[i] = ans[i-1] + 1 - count
		}
	}
	return ans
}

// @lc code=end

