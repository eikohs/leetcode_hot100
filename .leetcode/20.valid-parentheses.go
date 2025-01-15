/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 */

// @lc code=start
func isValid(s string) bool {
	// 定义要处理的括号
	leftBracket := "({["
	rightBracket := ")}]"
	// 定义栈结构
	stack := []rune{}
	// 遍历字符串，遇到左括号入栈，遇到右括号弹栈匹配
	for _, char := range s {
		if strings.Contains(leftBracket, string(char)) {
			// 如果 char 是左括号，执行相应的操作
			right := rightBracket[strings.Index(leftBracket, string(char))] // 获取对应的左括号
			stack = append(stack, rune(right))                              // 左括号入栈
		}
		if strings.Contains(rightBracket, string(char)) {
			// 如果 char 是右括号，执行相应的操作
			if (len(stack) == 0) ||
				(char != stack[len(stack)-1]) {
				return false // 不匹配
			} else {
				stack = stack[:len(stack)-1] // 匹配成功，弹栈
			}
		}
	}
	// 最后判断栈是否为空，非空则匹配失败
	if len(stack) != 0 {
		return false
	}
	return true
}

// @lc code=end

