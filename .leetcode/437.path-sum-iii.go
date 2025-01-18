/*
 * @lc app=leetcode id=437 lang=golang
 *
 * [437] Path Sum III
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

func countPathSum(node *TreeNode, targetSum int, currSum int, hash map[int]int) int {
	// 处理特殊情况
	if node == nil {
		return 0
	}
	// 更新前缀和并匹配哈希表
	currSum += node.Val
	res := hash[currSum-targetSum]
	// 将前缀和存储到哈希表中并遍历子结点
	hash[currSum]++
	res += countPathSum(node.Left, targetSum, currSum, hash) + countPathSum(node.Right, targetSum, currSum, hash)
	// 回溯前缀和
	hash[currSum]--
	return res
}

func pathSum(root *TreeNode, targetSum int) int {
	// 定义哈希表
	hash := make(map[int]int)
	// 初始化哈希表
	hash[0] = 1 // 前缀和为0的路径数为1，表示从根节点开始的路径
	// 开始遍历的同时获取结果
	return countPathSum(root, targetSum, 0, hash)
}

// @lc code=end

