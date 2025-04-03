/*
 * @lc app=leetcode id=17 lang=golang
 *
 * [17] Letter Combinations of a Phone Number
 */

// @lc code=start
func letterCombinations(digits string) []string {
	digitToLetter := []string{
		"abc",
		"def",
		"ghi",
		"jkl",
		"mno",
		"pqrs",
		"tuv",
		"wxyz",
	}

	result := []string{}
	combination := []byte{}
	n := len(digits)

	var dfs func(int)
	dfs = func(idx int) {
		if idx == n {
			result = append(result, string(combination))
			return
		}
		digit := int(digits[idx]-'0') - 2
		for _, char := range digitToLetter[digit] {
			combination = append(combination, byte(char))
			dfs(idx + 1)
			combination = combination[:len(combination)-1]
		}
	}

	if n > 0 {
		dfs(0)
	}

	return result
}

// @lc code=end

