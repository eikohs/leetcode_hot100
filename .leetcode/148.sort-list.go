/*
 * @lc app=leetcode id=148 lang=golang
 *
 * [148] Sort List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	merge := func(list1 *ListNode, list2 *ListNode) *ListNode {
		var node = &ListNode{}
		var hd = node

		for list1 != nil && list2 != nil {
			if list1.Val < list2.Val {
				node.Next = list1
				list1 = list1.Next
			} else {
				node.Next = list2
				list2 = list2.Next
			}
			node = node.Next
		}
		if list1 == nil {
			node.Next = list2
		} else {
			node.Next = list1
		}

		return hd.Next
	}

	var mergeSort func(*ListNode) *ListNode
	mergeSort = func(node *ListNode) *ListNode {
		if node == nil || node.Next == nil {
			return node
		}

		var slow, fast = node, node
		var prev *ListNode
		for fast != nil && fast.Next != nil {
			prev = slow
			slow = slow.Next
			fast = fast.Next.Next
		}
		slow = node
		fast = prev.Next
		prev.Next = nil

		return merge(mergeSort(slow), mergeSort(fast))
	}

	return mergeSort(head)
}

// @lc code=end

