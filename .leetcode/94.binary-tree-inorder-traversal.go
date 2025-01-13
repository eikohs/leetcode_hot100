/*
 * @lc app=leetcode id=94 lang=golang
 *
 * [94] Binary Tree Inorder Traversal
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

func inorderTrave(root *TreeNode, rst *[]int) {
	if root == nil {
		return
	}
	inorderTrave(root.Left, rst)  // 左子树
	*rst = append(*rst, root.Val) // 根节点
	inorderTrave(root.Right, rst) // 右子树
}

func inorderTraversal(root *TreeNode) []int {
	rst := make([]int, 0)    // 初始化结果数组
	inorderTrave(root, &rst) // 调用递归函数进行中序遍历
	return rst
}

// @lc code=end

