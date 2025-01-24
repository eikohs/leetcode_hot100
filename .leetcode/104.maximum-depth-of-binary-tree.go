/*
 * @lc app=leetcode id=104 lang=golang
 *
 * [104] Maximum Depth of Binary Tree
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

func getMaxDepth(node *TreeNode, depth int) int {
	if node == nil {
		return depth
	}
	left := getMaxDepth(node.Left, depth+1)
	right := getMaxDepth(node.Right, depth+1)
	return max(left, right)
}

func maxDepth(root *TreeNode) int {
	return getMaxDepth(root, 0)
}

// @lc code=end

