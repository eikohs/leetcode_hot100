/*
 * @lc app=leetcode id=70 lang=golang
 *
 * [70] Climbing Stairs
 */

// @lc code=start

func combination(n, k int) int {
	rst := 1
	for i := 0; i < k; i++ {
		rst = rst * (n - i) / (i + 1)
	}
	return rst
}

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	k := n / 2
	rst := 0
	for i := 0; i <= k; i++ {
		rst += combination(n-i, i)
	}
	return rst
}

// @lc code=end

