/*
 * @lc app=leetcode id=234 lang=golang
 *
 * [234] Palindrome Linked List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	// 处理特殊值
	if head == nil || head.Next == nil {
		return true
	}
	// 快慢节点法找到中间节点
	var prev *ListNode
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	prev.Next = nil // 断开链表
	// 反转后半部分链表
	prev = nil
	for slow != nil {
		next := slow.Next
		slow.Next = prev
		prev = slow
		slow = next
	}
	// 比较前半部分和反转后的后半部分链表
	for head != nil {
		if head.Val != prev.Val {
			return false
		}
		head = head.Next
		prev = prev.Next
	}
	return true
}

// @lc code=end

