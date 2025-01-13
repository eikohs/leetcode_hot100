/*
 * @lc app=leetcode id=98 lang=golang
 *
 * [98] Validate Binary Search Tree
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
func test(node *TreeNode, min, max int) bool {
	if node == nil {
		return true
	}
	if node.Val <= min || node.Val >= max {
		return false
	}
	return test(node.Left, min, node.Val) && // 左子树更新右边界（最大）为当前节点值
		test(node.Right, node.Val, max) // 右子树更新左边界（最小）为当前节点值
}

func isValidBST(root *TreeNode) bool {
	return test(root, -1<<31-1, 1<<31)
}

// @lc code=end

