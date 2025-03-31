/*
 * @lc app=leetcode id=39 lang=golang
 *
 * [39] Combination Sum
 */

// @lc code=start

func combinationSum(candidates []int, target int) [][]int {
	result := [][]int{}

	var backtrack func(int, int, []int)
	backtrack = func(idx int, target int, combination []int) {
		if target == 0 {
			// Need to make a copy of the combination to avoid overwriting
			comb := make([]int, len(combination))
			copy(comb, combination)
			result = append(result, comb)
			return
		}
		if target < 0 || idx >= len(candidates) {
			return
		}
		// Include current candidate
		combination = append(combination, candidates[idx])
		backtrack(idx, target-candidates[idx], combination)
		combination = combination[:len(combination)-1] // Backtrack

		// Skip to next candidate
		backtrack(idx+1, target, combination)
	}

	backtrack(0, target, []int{})
	return result
}

// @lc code=end

