/*
 * @lc app=leetcode id=114 lang=golang
 *
 * [114] Flatten Binary Tree to Linked List
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
func dfs(node *TreeNode) (*TreeNode, *TreeNode) {
	if node == nil {
		return nil, nil
	}
	lHead, lEnd := dfs(node.Left)
	rHead, rEnd := dfs(node.Right)
	node.Left = nil

	if lHead != nil {
		node.Right = lHead
	} else {
		lEnd = node
	}
	if rHead != nil {
		lEnd.Right = rHead
	} else {
		rEnd = lEnd
	}

	return node, rEnd
}

func flatten(root *TreeNode) {
	dfs(root)
}

// @lc code=end

