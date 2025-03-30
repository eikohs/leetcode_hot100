/*
 * @lc app=leetcode id=155 lang=golang
 *
 * [155] Min Stack
 */

// @lc code=start
type MinStack struct {
	stack  []int
	prefix []int
	length int
}

func Constructor() MinStack {
	minStack := MinStack{}
	return minStack
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	min := val
	if this.length > 0 && this.prefix[this.length-1] < val {
		min = this.prefix[this.length-1]
	}
	this.prefix = append(this.prefix, min)
	this.length++
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:this.length-1]
	this.prefix = this.prefix[:this.length-1]
	this.length--
}

func (this *MinStack) Top() int {
	return this.stack[this.length-1]
}

func (this *MinStack) GetMin() int {
	return this.prefix[this.length-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end

