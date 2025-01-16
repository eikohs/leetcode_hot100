/*
 * @lc app=leetcode id=21 lang=golang
 *
 * [21] Merge Two Sorted Lists
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var head *ListNode
	// 处理特殊情况
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	// 选择头节点
	if list1.Val < list2.Val {
		head = list1
		list1 = list1.Next
	} else {
		head = list2
		list2 = list2.Next
	}
	// 归并
	node := head
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			node.Next = list1
			node = list1
			list1 = list1.Next
		} else {
			node.Next = list2
			node = list2
			list2 = list2.Next
		}
	}
	// 处理剩余的节点
	if list1 != nil {
		node.Next = list1
	} else if list2 != nil {
		node.Next = list2
	}
	return head
}

// @lc code=end

