/*
 * @lc app=leetcode id=538 lang=golang
 *
 * [538] Convert BST to Greater Tree
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
func convertBST(root *TreeNode) *TreeNode {
	var makeGT func(*TreeNode, int) int
	makeGT = func(node *TreeNode, accum int) int {
		if node == nil {
			return 0
		}
		sum := node.Val
		rightSum := makeGT(node.Right, accum)
		node.Val = node.Val + accum + rightSum
		leftSum := makeGT(node.Left, node.Val)
		return sum + rightSum + leftSum
	}

	makeGT(root, 0)

	return root
}

// @lc code=end

