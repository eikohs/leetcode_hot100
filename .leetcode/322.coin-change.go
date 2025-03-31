/*
 * @lc app=leetcode id=322 lang=golang
 *
 * [322] Coin Change
 */

// @lc code=start
/*func coinChange(coins []int, amount int) int {
	var coinNums = -1

	var backtrack func(int, int, int)
	backtrack = func(idx int, target int, num int) {
		if target == 0 {
			if coinNums == -1 || num < coinNums {
				coinNums = num
			}
		}
		if target < 0 || idx >= len(coins) {
			return
		}
		// Use this coin
		num++
		backtrack(idx, target-coins[idx], num)
		num--

		// Not use
		backtrack(idx+1, target, num)
	}

	backtrack(0, amount, 0)
	return coinNums
}*/

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0

	for price := 1; price <= amount; price++ {
		dp[price] = amount + 1
		for _, coin := range coins {
			if price-coin >= 0 {
				dp[price] = min(dp[price], dp[price-coin]+1)
			}
		}
	}

	if dp[amount] > amount {
		return -1
	} else {
		return dp[amount]
	}
}

// @lc code=end

