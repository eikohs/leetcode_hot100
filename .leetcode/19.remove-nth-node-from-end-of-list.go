/*
 * @lc app=leetcode id=19 lang=golang
 *
 * [19] Remove Nth Node From End of List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 首先获取链表长度
	length := 0
	cur := head
	for cur != nil {
		length++
		cur = cur.Next
	}
	// 定位要删除的节点
	target := length - n
	if target == 0 { // 处理特殊值
		return head.Next
	}
	// 遍历链表找到目标节点并删除之
	var pre *ListNode
	cur = head
	for target > 0 {
		target--
		pre = cur
		cur = cur.Next
	}
	pre.Next = cur.Next
	return head
}

// @lc code=end

