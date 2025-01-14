/*
 * @lc app=leetcode id=617 lang=golang
 *
 * [617] Merge Two Binary Trees
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

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	// 处理有 nil 值的情况
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	// 否则合并节点
	root1.Val += root2.Val
	// 继续合并左子树与右子树
	root1.Left = mergeTrees(root1.Left, root2.Left)
	root1.Right = mergeTrees(root1.Right, root2.Right)
	return root1
}

// @lc code=end

