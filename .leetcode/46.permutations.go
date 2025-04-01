/*
 * @lc app=leetcode id=46 lang=golang
 *
 * [46] Permutations
 */

// @lc code=start
func permute(nums []int) [][]int {
	n := len(nums)
	flag := make([]bool, n)
	result := [][]int{}
	permutation := []int{}

	var dfs func(int)
	dfs = func(idx int) {
		if idx == n {
			perm := make([]int, n)
			copy(perm, permutation)
			result = append(result, perm)
			return
		}
		for key, num := range nums {
			if !flag[key] {
				permutation = append(permutation, num)
				flag[key] = true
				dfs(idx + 1)
				permutation = permutation[:len(permutation)-1]
				flag[key] = false
			}
		}
	}

	dfs(0)
	return result
}

// @lc code=end

