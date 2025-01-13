/*
 * @lc app=leetcode id=461 lang=golang
 *
 * [461] Hamming Distance
 */

// @lc code=start
func oldHammingDistance(x int, y int) int {
	hammingDistance := 0
	for x != 0 || y != 0 {
		if x&1 != y&1 { // 比较最低位
			hammingDistance++
		}
		x >>= 1 // mod 2
		y >>= 1 // mod 2
	}
	for x != 0 {
		x >>= 1
		hammingDistance++
	}
	for y != 0 {
		y >>= 1
		hammingDistance++
	}
	return hammingDistance
}

func hammingDistance(x int, y int) int {
	// 异或后计算1的个数
	xor := x ^ y
	hammingDistance := 0
	for xor != 0 {
		hammingDistance += xor & 1 // 取最低位
		xor >>= 1
	}
	return hammingDistance
	// 使用库函数的解法
	// return bits.OnesCount(uint(x ^ y))
}

// @lc code=end

