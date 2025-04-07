/*
 * @lc app=leetcode id=105 lang=golang
 *
 * [105] Construct Binary Tree from Preorder and Inorder Traversal
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

func buildTree(preorder []int, inorder []int) *TreeNode {
	n := len(preorder)
	node := 0

	// remember pos
	posHash := make(map[int]int)
	for key, val := range inorder {
		posHash[val] = key
	}

	var makeNode func(int, int) *TreeNode
	makeNode = func(left int, right int) *TreeNode {
		if node == n || left > right {
			return nil
		}
		// find pos
		pos := posHash[preorder[node]]
		// make node
		ptrTreeNode := &TreeNode{
			Val:   preorder[node],
			Left:  nil,
			Right: nil,
		}
		node++ // finish a node, jump to next one
		// get left child
		ptrTreeNode.Left = makeNode(left, pos-1)
		// get right child
		ptrTreeNode.Right = makeNode(pos+1, right)

		return ptrTreeNode
	}

	return makeNode(0, n-1)
}

// @lc code=end

