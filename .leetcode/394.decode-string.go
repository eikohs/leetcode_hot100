/*
 * @lc app=leetcode id=394 lang=golang
 *
 * [394] Decode String
 */

// @lc code=start
type storage struct {
	letters string
	times   int
}

type stack []storage

func (pstk *stack) push(str string, n int) {
	*pstk = append(*pstk, storage{
		letters: str,
		times:   n,
	})
}

func (pstk *stack) pop() string {
	stk := (*pstk)
	letters, times := stk[len(stk)-1].letters, stk[len(stk)-1].times
	*pstk = stk[:len(stk)-1]

	rst := ""
	for time := 0; time < times; time++ {
		rst = rst + letters
	}
	return rst
}

func (pstk *stack) add(str string) {
	stk := (*pstk)
	stk[len(stk)-1].letters = stk[len(stk)-1].letters + str
}

func (pstk *stack) empty() bool {
	return len(*pstk) == 0
}

func decodeString(s string) string {
	stk := &stack{}

	result := ""

	n := len(s)
	for i := 0; i < n; {
		if s[i] >= '0' && s[i] <= '9' {
			// 读到数字，进行放入栈的准备
			// 1. 读出数字
			times := 0
			for ; s[i] >= '0' && s[i] <= '9'; i++ {
				times = int(s[i]-'0') + times*10
			}
			if s[i] == '[' {
				i++
			}
			// 2. 读出字符串
			bytes := []byte{}
			for ; (s[i] < '0' || s[i] > '9') && s[i] != ']'; i++ {
				bytes = append(bytes, byte(s[i]))
			}
			// 3. 放入栈中
			stk.push(string(bytes), times)
		} else if s[i] == ']' {
			// 读到右中括号，弹栈
			// 1. 弹栈
			str := stk.pop()
			// 2. 判断是否到达栈底，是则放入结果字符串中，否则放入 cache 中
			if stk.empty() {
				result = result + str
			} else {
				stk.add(str)
			}
			i++
		} else if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z') {
			// 读到原始字符串，直接加到 cache 中或结果字符串中
			bytes := []byte{}
			for ; i < n && (s[i] < '0' || s[i] > '9') && s[i] != ']'; i++ {
				bytes = append(bytes, byte(s[i]))
			}
			if stk.empty() {
				result = result + string(bytes)
			} else {
				stk.add(string(bytes))
			}
		}
	}

	return result
}

// @lc code=end

