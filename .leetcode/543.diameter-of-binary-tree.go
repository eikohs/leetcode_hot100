/*
 * @lc app=leetcode id=543 lang=golang
 *
 * [543] Diameter of Binary Tree
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

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func calculateDiameter(node *TreeNode, max *int) (depth int) {
	if node == nil {
		depth = 0
		return
	}
	leftDepth := calculateDiameter(node.Left, max)
	rightDepth := calculateDiameter(node.Right, max)
	*max = maxInt(*max, leftDepth+rightDepth)
	depth = maxInt(leftDepth, rightDepth) + 1
	return
}

func diameterOfBinaryTree(root *TreeNode) int {
	var result int
	calculateDiameter(root, &result)
	return result
}

// @lc code=end

