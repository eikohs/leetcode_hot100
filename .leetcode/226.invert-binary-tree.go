/*
 * @lc app=leetcode id=226 lang=golang
 *
 * [226] Invert Binary Tree
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
func invertTree(root *TreeNode) *TreeNode {
	var tmp *TreeNode
	if root != nil {
		tmp = root.Left
		root.Left = root.Right
		root.Right = tmp
		invertTree(root.Left)
		invertTree(root.Right)
	}
	return root
}

// @lc code=end

