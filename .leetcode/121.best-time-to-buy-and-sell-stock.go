/*
 * @lc app=leetcode id=121 lang=golang
 *
 * [121] Best Time to Buy and Sell Stock
 */

// @lc code=start
func maxProfit(prices []int) int {
	n := len(prices)
	min := make([]int, n)
	rst := 0
	min[0] = prices[0]

	// 完善 min 数组，同时尝试低买高卖
	for idx := 1; idx < n; idx++ {
		price := prices[idx]
		profit := price - min[idx-1]
		if profit > rst {
			rst = profit
		}
		if price < min[idx-1] {
			min[idx] = price
		} else {
			min[idx] = min[idx-1]
		}
	}
	return rst
}

// @lc code=end

