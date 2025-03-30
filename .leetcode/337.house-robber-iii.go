/*
 * @lc app=leetcode id=337 lang=golang
 *
 * [337] House Robber III
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func startRub(node *TreeNode) (rubNode int, rubNext int) {
	if node == nil {
		return 0, 0
	}
	LrubNode, LrubNext := startRub(node.Left)
	RrubNode, RrubNext := startRub(node.Right)
	return node.Val + LrubNext + RrubNext, max(LrubNode, LrubNext) + max(RrubNode, RrubNext)
}

func rob(root *TreeNode) int {
	rubNode, rubNext := startRub(root)
	return max(rubNode, rubNext)
}

// @lc code=end

