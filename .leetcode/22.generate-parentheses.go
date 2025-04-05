/*
 * @lc app=leetcode id=22 lang=golang
 *
 * [22] Generate Parentheses
 */

// @lc code=start
func generateParenthesis(n int) []string {
	result := []string{}
	bytes := make([]rune, 0, 2*n)

	var dfs func(int, int)
	dfs = func(left int, right int) {
		// 处理错误的情况
		if left < 0 || left > right {
			return
		}
		// 处理结果
		if left == 0 && right == 0 {
			result = append(result, string(bytes))
			return
		}
		bytes = append(bytes, '(')
		dfs(left-1, right)
		bytes[len(bytes)-1] = ')'
		dfs(left, right-1)
		bytes = bytes[:len(bytes)-1]
	}

	dfs(n, n)
	return result
}

// @lc code=end

